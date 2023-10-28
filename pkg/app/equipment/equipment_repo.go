package equipment

import "go.mongodb.org/mongo-driver/bson/primitive"

type EquipmentRepository interface {
	GetAll() ([]*Equipment, error)
	GetByID(id primitive.ObjectID) (*Equipment, error)
	GetBySearch(searchQuery string) ([]*Equipment, error)
	Create(equipment *Equipment) error
	Update(equipment *Equipment) error
	UpdateQuantity(id primitive.ObjectID, command string) error
	DeleteByID(id primitive.ObjectID) error
}
