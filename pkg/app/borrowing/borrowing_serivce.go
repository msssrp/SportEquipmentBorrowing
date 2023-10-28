package borrowing

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BorrowingService interface {
	GetAllBorrowings() ([]*Borrowing, error)
	GetBorrowingByID(id primitive.ObjectID) (*Borrowing, error)
	GetBorrowingsByUserID(userID primitive.ObjectID) ([]*Borrowing, error)
	GetBorrowingByEquipmentID(equipmentID primitive.ObjectID) (*Borrowing, error)
	CreateBorrowing(borrowing *Borrowing) error
	UpdateBorrowing(borrowing *Borrowing) error
	DeleteBorrowingByID(id primitive.ObjectID) error
	ApproveBorrow(id primitive.ObjectID) error
}

type borrowingService struct {
	borrowingRepo BorrowingRepository
}

func NewBorrowingService(borrowingRepo BorrowingRepository) BorrowingService {
	return &borrowingService{
		borrowingRepo: borrowingRepo,
	}
}

// serivce func

//Get
func (s *borrowingService) GetAllBorrowings() ([]*Borrowing, error) {
	return s.borrowingRepo.GetAll()
}

func (s *borrowingService) GetBorrowingByID(id primitive.ObjectID) (*Borrowing, error) {
	objectID, err := s.GetID(id)
	if err != nil {
		return nil, err
	}

	return s.borrowingRepo.GetByID(objectID)
}

func (s *borrowingService) GetBorrowingsByUserID(userID primitive.ObjectID) ([]*Borrowing, error) {
	objectID, err := s.GetID(userID)
	if err != nil {
		return nil, err
	}

	return s.borrowingRepo.GetByUserID(objectID)
}

func (s *borrowingService) GetBorrowingByEquipmentID(equipmentID primitive.ObjectID) (*Borrowing, error) {
	objectID, err := s.GetID(equipmentID)
	if err != nil {
		return nil, err
	}
	return s.borrowingRepo.GetByEquipmentID(objectID)
}

func (s *borrowingService) GetID(id primitive.ObjectID) (primitive.ObjectID, error) {
	if id == primitive.NilObjectID {
		return primitive.NilObjectID, errors.New("invalid id please provide id")
	}

	return id, nil
}

//Post
func (s *borrowingService) CreateBorrowing(borrowing *Borrowing) error {
	if borrowing == nil {
		return errors.New("borrowing is nil")
	}

	if borrowing.User_id == primitive.NilObjectID || borrowing.Equipment_id == primitive.NilObjectID || borrowing.Borrow_date.IsZero() || borrowing.Return_date.IsZero() || borrowing.DayLeft == 0 {
		return errors.New("all the fields are required please provide all the fields")
	}
	return s.borrowingRepo.Create(borrowing)
}

func (s *borrowingService) ApproveBorrow(id primitive.ObjectID) error {
	objectID, err := s.GetID(id)
	if err != nil {
		return err
	}

	return s.borrowingRepo.ApproveEquipmentBorrow(objectID)
}

//Put
func (s *borrowingService) UpdateBorrowing(borrowing *Borrowing) error {
	if borrowing == nil {
		return errors.New("borrowing is nil")
	}

	if borrowing.Id == primitive.NilObjectID && borrowing.Equipment_id == primitive.NilObjectID && borrowing.User_id == primitive.NilObjectID && borrowing.Status == "" && borrowing.Borrow_date.IsZero() && borrowing.Return_date.IsZero() {
		return errors.New("atleast one field is required please provide some field")
	}

	existingBorrowing, err := s.borrowingRepo.GetByID(borrowing.Id)
	if err != nil {
		return err
	}

	if borrowing.Return_date.IsZero() {
		existingBorrowing.Return_date = borrowing.Return_date
	}
	if borrowing.Borrow_date.IsZero() {
		existingBorrowing.Return_date = borrowing.Borrow_date
	}
	if borrowing.Status != "" {
		existingBorrowing.Status = borrowing.Status
	}
	if borrowing.Equipment_id != primitive.NilObjectID {
		existingBorrowing.Equipment_id = borrowing.Equipment_id
	}
	if borrowing.User_id != primitive.NilObjectID {
		existingBorrowing.User_id = borrowing.User_id
	}

	return s.borrowingRepo.Update(existingBorrowing)
}

//Delete
func (s *borrowingService) DeleteBorrowingByID(id primitive.ObjectID) error {
	objectID, err := s.GetID(id)
	if err != nil {
		return err
	}

	return s.borrowingRepo.DeleteByID(objectID)
}
