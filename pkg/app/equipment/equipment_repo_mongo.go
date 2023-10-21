package equipment

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type equipmentRepositoryMongo struct {
	collection *mongo.Collection
}

func NewEquipmentRepositoryMongo(client *mongo.Client, dbName string, collectionName string) EquipmentRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &equipmentRepositoryMongo{
		collection: collection,
	}
}

//Get
func (r *equipmentRepositoryMongo) GetAll() ([]*Equipment, error) {
	var equipments []*Equipment

	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var equipment Equipment
		if err := cursor.Decode(&equipment); err != nil {
			return nil, err
		}
		equipments = append(equipments, &equipment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return equipments, nil
}

func (r *equipmentRepositoryMongo) GetByID(id primitive.ObjectID) (*Equipment, error) {
	filter := bson.M{"_id": id}

	var equipment Equipment
	err := r.collection.FindOne(context.Background(), filter).Decode(&equipment)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("equipment not found")
		}
		return nil, err
	}

	return &equipment, nil
}

//Post
func (r *equipmentRepositoryMongo) Create(equipment *Equipment) error {
	_, err := r.collection.InsertOne(context.Background(), equipment)
	return err
}

//Put
func (r *equipmentRepositoryMongo) Update(equipment *Equipment) error {
	filter := bson.M{"_id": equipment.Id}
	update := bson.M{
		"$set": equipment,
	}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (r *equipmentRepositoryMongo) UpdateQuantityToPending(equipmentID primitive.ObjectID) error {
	filter := bson.M{"_id": equipmentID}
	update := bson.M{
		"$set": bson.M{"quantity_available": "In use"},
	}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

//Delete
func (r equipmentRepositoryMongo) DeleteByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}
