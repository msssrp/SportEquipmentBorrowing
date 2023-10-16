package http

import (
	"github.com/gin-gonic/gin"
	"github.com/msssrp/SportEquipmentBorrowing/function"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/http/handler"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/middleware"
)

func SetRoutes(router *gin.Engine, app *app.App) {

	userHandler := handler.NewUserHandler(app)
	equipmentHandler := handler.NewEquipmentHandler(app)
	borrowingHandler := handler.NewBorrowingHandler(app)

	secret, err := function.GetDotEnv("SECRET")
	if err != nil {
		panic(err)
	}

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", middleware.JWTMiddleware(secret), userHandler.HandlerGetUsers)
		userRoutes.GET("/:id", middleware.JWTMiddleware(secret), userHandler.HandlerGetUserByID)
		userRoutes.POST("", userHandler.HandlerCreateUser)
		userRoutes.POST("/auth/signIn", userHandler.HandlerSignIn)
		userRoutes.PUT("/:id", middleware.JWTMiddleware(secret), userHandler.HandlerUpdateUser)
		userRoutes.DELETE("/:id", middleware.JWTMiddleware(secret), userHandler.HandlerDeleteUser)
	}

	equipmentRoutes := router.Group("/equipment")
	{
		equipmentRoutes.GET("", equipmentHandler.HandlerGetEquipments)
		equipmentRoutes.GET("/:id", equipmentHandler.HandlerGetEquipmentByID)
		equipmentRoutes.POST("", middleware.JWTMiddleware(secret), equipmentHandler.HandlerCreateEquipment)
		equipmentRoutes.PUT("/:id", middleware.JWTMiddleware(secret), equipmentHandler.HandlerUpdateEquipment)
		equipmentRoutes.DELETE("/:id", middleware.JWTMiddleware(secret), equipmentHandler.HandlerDeleteEquipment)
	}

	borrowingRoutes := router.Group("/borrowing")
	{
		borrowingRoutes.GET("/:id", borrowingHandler.HandlerGetBorrowingByID)
		borrowingRoutes.GET("/getByUser/:id", borrowingHandler.HandlerGetBorrowingsByUserID)
		borrowingRoutes.GET("/getByEquipment/:id", borrowingHandler.HandlerGetBorrowingByEquipmentID)
		borrowingRoutes.GET("", borrowingHandler.HandlerGetAllBorrowings)
		borrowingRoutes.POST("", middleware.JWTMiddleware(secret), borrowingHandler.HandlerCreateBorrowing)
		borrowingRoutes.PUT("/:id", middleware.JWTMiddleware(secret), borrowingHandler.HandlerUpdateBorrowing)
		borrowingRoutes.DELETE("/:id", middleware.JWTMiddleware(secret), borrowingHandler.HandlerDeleteBorrowing)
	}
}
