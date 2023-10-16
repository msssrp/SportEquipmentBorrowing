package user

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetAllUsers() ([]*User, error)
	GetUserByID(id primitive.ObjectID) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id primitive.ObjectID) error
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

//Get
func (s *userService) GetAllUsers() ([]*User, error) {
	return s.userRepo.GetAll()
}

func (s *userService) GetUserByID(id primitive.ObjectID) (*User, error) {
	if id == primitive.NilObjectID {
		return nil, errors.New("id is null please provide id")
	}
	return s.userRepo.GetByID(id)
}

//Post
func (s *userService) CreateUser(user *User) error {
	if user == nil {
		return errors.New("user is null please pass the user infomations")
	}
	if user.Password == "" || user.Username == "" || user.Phone_number == "" || user.Address == "" || user.L_name == "" || user.F_name == "" || user.Email == "" {
		return errors.New("all the fields are required plrease provide all the fields")
	}
	return s.userRepo.Create(user)
}

//Put
func (s *userService) UpdateUser(user *User) error {
	if user == nil {
		return errors.New("user is null, please provide user information")
	}

	if user.Password == "" && user.Username == "" && user.Phone_number == "" && user.Address == "" && user.L_name == "" && user.F_name == "" && user.Email == "" {
		return errors.New("at least one field is required, please provide at least one non-empty field")
	}

	existingUser, err := s.userRepo.GetByID(user.Id)
	if err != nil {
		return err
	}

	if user.Username != "" {
		existingUser.Username = user.Username
	}
	if user.Password != "" {
		existingUser.Password = user.Password
	}
	if user.Phone_number != "" {
		existingUser.Phone_number = user.Phone_number
	}
	if user.Address != "" {
		existingUser.Address = user.Address
	}
	if user.L_name != "" {
		existingUser.L_name = user.L_name
	}
	if user.F_name != "" {
		existingUser.F_name = user.F_name
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}

	return s.userRepo.Update(user)
}

//Delete
func (s *userService) DeleteUser(id primitive.ObjectID) error {
	if id == primitive.NilObjectID {
		return errors.New("invalid id please provide id")
	}
	return s.userRepo.DeleteByID(id)
}
