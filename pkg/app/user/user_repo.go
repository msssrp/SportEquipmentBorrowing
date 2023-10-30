package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRepository interface {
	GetAll() ([]*User, error)
	GetByID(id primitive.ObjectID) (*User, error)
	Create(user *User) error
	Update(user *User) error
	DeleteByID(id primitive.ObjectID) error
	SignIn(username string, password string) (string, string, error)
	GenerateNewAccessToken(userID string) (string, error)
	GetUserRoles(userID primitive.ObjectID) ([]string, error)
}
