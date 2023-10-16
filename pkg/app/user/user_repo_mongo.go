package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryMongo struct {
	collection *mongo.Collection
}

func NewUserRepositoryMongo(client *mongo.Client, dbName string, collectionName string) UserRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &userRepositoryMongo{
		collection: collection,
	}
}

//Get
func (r *userRepositoryMongo) GetAll() ([]*User, error) {
	var users []*User

	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepositoryMongo) GetByID(id primitive.ObjectID) (*User, error) {
	filter := bson.M{"_id": id}

	var user User
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

//Post
func (r *userRepositoryMongo) Create(user *User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

//Put
func (r *userRepositoryMongo) Update(user *User) error {
	filter := bson.M{"_id": user.Id}
	update := bson.M{
		"$set": user,
	}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

//Delete
func (r *userRepositoryMongo) DeleteByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}
