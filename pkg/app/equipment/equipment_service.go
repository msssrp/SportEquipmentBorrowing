package equipment

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EquipmentService interface {
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

func (s *equipmentService) GetEquipmentByID(id primitive.ObjectID) (*Equipment, error) {
	if id != primitive.NilObjectID {
		return nil, errors.New("invalid id type")
	}
	return s.equipmentRepo.GetByID(id)
}

func (s *equipmentService) CreateEquipment(equipment *Equipment) error {
	if equipment == nil {
		return errors.New("equipment is null please pass the equipment infomations")
	}
	if equipment.Quantity_available == "" || equipment.Name == "" || equipment.Category == "" || equipment.Condition == "" || equipment.Image_url == "" || equipment.Description == "" {
		return errors.New("all the fields are required plrease provide all the fields")
	}
	return s.equipmentRepo.Create(equipment)
}

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

	if equipment.Quantity_available != "" {
		existingequipment.Quantity_available = equipment.Quantity_available
	}
	if equipment.Name != "" {
		existingequipment.Name = equipment.Name
	}
	if equipment.Category != "" {
		existingequipment.Category = equipment.Category
	}
	if equipment.Condition != "" {
		existingequipment.Condition = equipment.Condition
	}
	if equipment.Image_url != "" {
		existingequipment.Image_url = equipment.Image_url
	}
	if equipment.Description != "" {
		existingequipment.Description = equipment.Description
	}

	return s.equipmentRepo.Update(equipment)
}

func (s *equipmentService) DeleteEquipment(id primitive.ObjectID) error {
	if id != primitive.NilObjectID {
		return errors.New("invalid id type")
	}
	return s.equipmentRepo.DeleteByID(id)
}
