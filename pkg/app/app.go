package app

import (
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app/borrowing"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app/equipment"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app/user"
)

type App struct {
	UserService      user.UserService
	EquipmentService equipment.EquipmentService
	BorrowingService borrowing.BorrowingService
}

func NewApp(userService user.UserService, equipmentService equipment.EquipmentService, borrowingService borrowing.BorrowingService) *App {
	return &App{
		UserService:      userService,
		EquipmentService: equipmentService,
		BorrowingService: borrowingService,
	}
}
