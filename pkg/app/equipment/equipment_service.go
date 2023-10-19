package equipment

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EquipmentService interface {
	GetAllEquipments() ([]*Equipment, error)
	GetEquipmentByID(id primitive.ObjectID) (*Equipment, error)
	CreateEquipment(equipment *Equipment) error
	UpdateEquipment(equipment *Equipment) error
	DeleteEquipment(id primitive.ObjectID) error
}

type equipmentService struct {
	equipmentRepo EquipmentRepository
}

func NewequipmentService(equipmentRepo EquipmentRepository) EquipmentService {
	return &equipmentService{
		equipmentRepo: equipmentRepo,
	}
}

//Get
func (s *equipmentService) GetAllEquipments() ([]*Equipment, error) {
	return s.equipmentRepo.GetAll()
}

func (s *equipmentService) GetEquipmentByID(id primitive.ObjectID) (*Equipment, error) {
	if id == primitive.NilObjectID {
		return nil, errors.New("invalid id please provide id")
	}
	return s.equipmentRepo.GetByID(id)
}

//Post
func (s *equipmentService) CreateEquipment(equipment *Equipment) error {
	if equipment == nil {
		return errors.New("equipment is null please pass the equipment infomations")
	}
	if equipment.Quantity_available == "" || equipment.Name == "" || equipment.Category == "" || equipment.Condition == "" || equipment.Image_url == "" || equipment.Description == "" {
		return errors.New("all the fields are required plrease provide all the fields")
	}
	return s.equipmentRepo.Create(equipment)
}

//Put
func (s *equipmentService) UpdateEquipment(equipment *Equipment) error {
	if equipment == nil {
		return errors.New("equipment is null, please provide equipment information")
	}

	if equipment.Description == "" && equipment.Image_url == "" && equipment.Condition == "" && equipment.Category == "" && equipment.Name == "" && equipment.Quantity_available == "" {
		return errors.New("at least one field is required, please provide at least one non-empty field")
	}

	existingequipment, err := s.equipmentRepo.GetByID(equipment.Id)
	if err != nil {
		return err
	}

	if equipment.Quantity_available == "" {
		equipment.Quantity_available = existingequipment.Quantity_available
	}
	if equipment.Name == "" {
		equipment.Name = existingequipment.Name
	}
	if equipment.Category == "" {
		equipment.Category = existingequipment.Category
	}
	if equipment.Condition == "" {
		equipment.Condition = existingequipment.Condition
	}
	if equipment.Image_url == "" {
		equipment.Image_url = existingequipment.Image_url
	}
	if equipment.Description == "" {
		equipment.Description = existingequipment.Description
	}

	return s.equipmentRepo.Update(equipment)
}

//Delete
func (s *equipmentService) DeleteEquipment(id primitive.ObjectID) error {
	if id == primitive.NilObjectID {
		return errors.New("invalid id please provide id")
	}
	return s.equipmentRepo.DeleteByID(id)
}
