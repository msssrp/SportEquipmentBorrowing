package borrowing

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type borrowingRepositoryMongo struct {
	collection *mongo.Collection
}

func NewBorrowingRepositoryMongo(client *mongo.Client, dbName, collectionName string) BorrowingRepository {
	collection := client.Database(dbName).Collection(collectionName)

	return &borrowingRepositoryMongo{
		collection: collection,
	}
}

func (r *borrowingRepositoryMongo) GetByID(id primitive.ObjectID) (*Borrowing, error) {
	filter := bson.M{"_id": id}

	var borrowing Borrowing
	err := r.collection.FindOne(context.Background(), filter).Decode(&borrowing)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("borrowing not found")
		}
		return nil, err
	}

	return &borrowing, nil
}

func (r *borrowingRepositoryMongo) Create(borrowing *Borrowing) error {
	_, err := r.collection.InsertOne(context.Background(), borrowing)
	return err
}

func (r *borrowingRepositoryMongo) Update(borrowing *Borrowing) error {
	filter := bson.M{"_id": borrowing.Id}
	update := bson.M{
		"$set": borrowing,
	}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (r *borrowingRepositoryMongo) DeleteByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}

func (r *borrowingRepositoryMongo) GetAll() ([]*Borrowing, error) {
	var borrowings []*Borrowing

	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var borrowing Borrowing
		if err := cursor.Decode(&borrowing); err != nil {
			return nil, err
		}
		borrowings = append(borrowings, &borrowing)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return borrowings, nil
}

func (r *borrowingRepositoryMongo) GetByUserID(userID primitive.ObjectID) ([]*Borrowing, error) {
	filter := bson.M{"user_id": userID}

	var borrowings []*Borrowing
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var borrowing Borrowing
		if err := cursor.Decode(&borrowing); err != nil {
			return nil, err
		}
		borrowings = append(borrowings, &borrowing)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return borrowings, nil
}

func (r *borrowingRepositoryMongo) GetByEquipmentID(equipmentID primitive.ObjectID) ([]*Borrowing, error) {
	filter := bson.M{"equipment_id": equipmentID}

	var borrowings []*Borrowing
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var borrowing Borrowing
		if err := cursor.Decode(&borrowing); err != nil {
			return nil, err
		}
		borrowings = append(borrowings, &borrowing)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return borrowings, nil
}
