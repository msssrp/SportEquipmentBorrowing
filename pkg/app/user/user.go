package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id           primitive.ObjectID `json:"user_id" bson:"_id,omitempty"`
	Username     string             `json:"username" bson:"username"`
	Password     string             `json:"-" bson:"password"`
	Email        string             `json:"email" bson:"email"`
	F_name       string             `json:"f_name" bson:"f_name"`
	L_name       string             `json:"l_name" bson:"l_name"`
	Phone_number string             `json:"phone_number" bson:"phone_number"`
	Address      string             `json:"address" bson:"address"`
	Roles        []string           `json:"roles" bson:"roles"`
}
