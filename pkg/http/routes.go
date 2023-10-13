package http

import (
	"github.com/gin-gonic/gin"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/http/handler"
)

func SetRoutes(router *gin.Engine, app *app.App) {

	userHandler := handler.NewUserHandler(app)
	equipmentHandler := handler.NewEquipmentHandler(app)
	borrowingHandler := handler.NewBorrowingHandler(app)

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", userHandler.HandlerGetUsers)
		userRoutes.GET("/:id", userHandler.HandlerGetUserByID)
		userRoutes.POST("", userHandler.HandlerCreateUser)
		userRoutes.PUT("/:id", userHandler.HandlerUpdateUser)
		userRoutes.DELETE("/:id", userHandler.HandlerDeleteUser)
	}

	equipmentRoutes := router.Group("/equipment")
	{
		equipmentRoutes.GET("", equipmentHandler.HandlerGetEquipments)
		equipmentRoutes.GET("/:id", equipmentHandler.HandlerGetEquipmentByID)
		equipmentRoutes.POST("", equipmentHandler.HandlerCreateEquipment)
		equipmentRoutes.PUT("/:id", equipmentHandler.HandlerUpdateEquipment)
		equipmentRoutes.DELETE("/:id", equipmentHandler.HandlerDeleteEquipment)
	}

	borrowingRoutes := router.Group("/borrowing")
	{
		borrowingRoutes.GET("/:id", borrowingHandler.HandlerGetBorrowingByID)
		borrowingRoutes.GET("/getByUser/:id", borrowingHandler.HandlerGetBorrowingsByUserID)
		borrowingRoutes.GET("", borrowingHandler.HandlerGetAllBorrowings)
		borrowingRoutes.POST("", borrowingHandler.HandlerCreateBorrowing)
		borrowingRoutes.PUT("/:id", borrowingHandler.HandlerUpdateBorrowing)
		borrowingRoutes.DELETE("/:id", borrowingHandler.HandlerDeleteBorrowing)
	}
}
