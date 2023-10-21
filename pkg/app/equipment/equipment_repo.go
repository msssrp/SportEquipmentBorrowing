package equipment

import "go.mongodb.org/mongo-driver/bson/primitive"

type EquipmentRepository interface {
	GetAll() ([]*Equipment, error)
	GetByID(id primitive.ObjectID) (*Equipment, error)
	Create(equipment *Equipment) error
	Update(equipment *Equipment) error
	UpdateQuantityToPending(id primitive.ObjectID) error
	DeleteByID(id primitive.ObjectID) error
}
