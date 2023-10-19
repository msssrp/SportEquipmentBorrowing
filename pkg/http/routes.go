package http

import (
	"github.com/gin-gonic/gin"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/http/handler"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/middleware"
)

func SetRoutes(router *gin.Engine, app *app.App) {

	userHandler := handler.NewUserHandler(app)
	equipmentHandler := handler.NewEquipmentHandler(app)
	borrowingHandler := handler.NewBorrowingHandler(app)

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", middleware.JWTMiddleware(), userHandler.HandlerGetUsers)
		userRoutes.GET("/byId", middleware.JWTMiddleware(), userHandler.HandlerGetUserByID)
		userRoutes.POST("", middleware.RateLimiterMiddleware(), userHandler.HandlerCreateUser)
		userRoutes.POST("/auth/signIn", userHandler.HandlerSignIn)
		userRoutes.PUT("/:id", middleware.JWTMiddleware(), userHandler.HandlerUpdateUser)
		userRoutes.DELETE("/:id", middleware.JWTMiddleware(), userHandler.HandlerDeleteUser)
	}

	equipmentRoutes := router.Group("/equipment")
	{
		equipmentRoutes.GET("", equipmentHandler.HandlerGetEquipments)
		equipmentRoutes.GET("/:id", equipmentHandler.HandlerGetEquipmentByID)
		equipmentRoutes.POST("", middleware.JWTMiddleware(), equipmentHandler.HandlerCreateEquipment)
		equipmentRoutes.PUT("/:id", middleware.JWTMiddleware(), equipmentHandler.HandlerUpdateEquipment)
		equipmentRoutes.DELETE("/:id", middleware.JWTMiddleware(), equipmentHandler.HandlerDeleteEquipment)
	}

	borrowingRoutes := router.Group("/borrowing")
	{
		borrowingRoutes.GET("/:id", borrowingHandler.HandlerGetBorrowingByID)
		borrowingRoutes.GET("/getByUser/:id", borrowingHandler.HandlerGetBorrowingsByUserID)
		borrowingRoutes.GET("/getByEquipment/:id", borrowingHandler.HandlerGetBorrowingByEquipmentID)
		borrowingRoutes.GET("", borrowingHandler.HandlerGetAllBorrowings)
		borrowingRoutes.POST("", middleware.JWTMiddleware(), borrowingHandler.HandlerCreateBorrowing)
		borrowingRoutes.PUT("/:id", middleware.JWTMiddleware(), borrowingHandler.HandlerUpdateBorrowing)
		borrowingRoutes.DELETE("/:id", middleware.JWTMiddleware(), borrowingHandler.HandlerDeleteBorrowing)
	}
}
