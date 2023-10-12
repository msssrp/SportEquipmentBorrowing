package borrowing

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Borrowing struct {
	Id           primitive.ObjectID `json:"borrowing_id" bson:"_id,omitempty"`
	User_id      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Equipment_id primitive.ObjectID `json:"equipment_id" bson:"equipment_id"`
	Borrow_date  time.Time          `json:"borrow_date" bson:"borrow_date"`
	Return_date  time.Time          `json:"return_date" bson:"return_date"`
	Status       string             `json:"status" bson:"status"`
}
