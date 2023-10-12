package borrowing

import "go.mongodb.org/mongo-driver/bson/primitive"

type BorrowingRepository interface {
	GetByID(id primitive.ObjectID) (*Borrowing, error)
	Create(borrowing *Borrowing) error
	Update(borrowing *Borrowing) error
	DeleteByID(id primitive.ObjectID) error
	GetAll() ([]*Borrowing, error)
	GetByUserID(userID primitive.ObjectID) ([]*Borrowing, error)
	GetByEquipmentID(equipmentID primitive.ObjectID) ([]*Borrowing, error)
}
