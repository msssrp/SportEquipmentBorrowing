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
		userRoutes.GET("", middleware.AccessTokenMiddleware(), middleware.JWTVerify(), userHandler.HandlerGetUsers)
		userRoutes.GET("/byId", middleware.AccessTokenMiddleware(), middleware.JWTMiddleware(), userHandler.HandlerGetUserByIDFromToken)
		userRoutes.GET("/byID/:id", userHandler.HandlersGetUserByID)
		userRoutes.GET("/auth/session", middleware.AuthenticateSession(), userHandler.HandlerVerifySession)
		userRoutes.GET("/auth/userRoles/:id", userHandler.HandlerGetUserRolesByID)
		userRoutes.POST("", middleware.RateLimiterMiddleware(), userHandler.HandlerCreateUser)
		userRoutes.POST("/auth/signIn", userHandler.HandlerSignIn)
		userRoutes.POST("/auth/refreshToken", middleware.AccessTokenMiddleware(), middleware.JWTGetClaims(), userHandler.HanlderNewAccessToken)
		userRoutes.PUT("/:id", middleware.AccessTokenMiddleware(), middleware.JWTVerify(), userHandler.HandlerUpdateUser)
		userRoutes.DELETE("/:id", middleware.AccessTokenMiddleware(), middleware.JWTVerify(), userHandler.HandlerDeleteUser)
	}

	equipmentRoutes := router.Group("/equipment")
	{
		equipmentRoutes.GET("", equipmentHandler.HandlerGetEquipments)
		equipmentRoutes.GET("/:id", equipmentHandler.HandlerGetEquipmentByID)
		equipmentRoutes.GET("/search", equipmentHandler.HandlerGetEquipmentBySearch)
		equipmentRoutes.POST("", middleware.AccessTokenMiddleware(), middleware.JWTVerify(), equipmentHandler.HandlerCreateEquipment)
		equipmentRoutes.PUT("/:id", middleware.AccessTokenMiddleware(), middleware.JWTVerify(), equipmentHandler.HandlerUpdateEquipment)
		equipmentRoutes.DELETE("/:id", middleware.AccessTokenMiddleware(), middleware.JWTVerify(), equipmentHandler.HandlerDeleteEquipment)
	}

	borrowingRoutes := router.Group("/borrowing")
	{
		borrowingRoutes.GET("/:id", borrowingHandler.HandlerGetBorrowingByID)
		borrowingRoutes.GET("/getByUser/:id", borrowingHandler.HandlerGetBorrowingsByUserID)
		borrowingRoutes.GET("/getByEquipment/:id", borrowingHandler.HandlerGetBorrowingByEquipmentID)
		borrowingRoutes.GET("", borrowingHandler.HandlerGetAllBorrowings)
		borrowingRoutes.POST("", middleware.AccessTokenMiddleware(), middleware.JWTVerify(), borrowingHandler.HandlerCreateBorrowing)
		borrowingRoutes.POST("/approveBorrowing", middleware.AccessTokenMiddleware(), middleware.JWTMiddleware(), borrowingHandler.HandlerApproveBorrowing)
		borrowingRoutes.PUT("/:id", middleware.AccessTokenMiddleware(), middleware.JWTVerify(), borrowingHandler.HandlerUpdateBorrowing)
		borrowingRoutes.DELETE("/:id/:equipmentID", middleware.AccessTokenMiddleware(), middleware.JWTMiddleware(), borrowingHandler.HandlerDeleteBorrowing)
	}
}
