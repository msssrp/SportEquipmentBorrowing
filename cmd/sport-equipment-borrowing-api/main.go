package main

import (
	"github.com/gin-gonic/gin"
	"github.com/msssrp/SportEquipmentBorrowing/function"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app/borrowing"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app/equipment"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app/user"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/database"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/http"
)

func main() {

	mongoDB, err := database.ConnectMongo()
	if err != nil {
		panic(err)
	}

	secretKey, err := function.GetDotEnv("SECRET")
	if err != nil {
		panic(err)
	}

	//init all repositories
	userRepo := user.NewUserRepositoryMongo(mongoDB.Client(), "SportEquipmentBorrowing", "users", []byte(secretKey))
	equipmentRepo := equipment.NewEquipmentRepositoryMongo(mongoDB.Client(), "SportEquipmentBorrowing", "equipments")
	borrowingRepo := borrowing.NewBorrowingRepositoryMongo(mongoDB.Client(), "SportEquipmentBorrowing", "borrowing")
	//init app
	a := app.NewApp(user.NewUserService(userRepo), equipment.NewequipmentService(equipmentRepo), borrowing.NewBorrowingService(borrowingRepo))

	router := gin.Default()

	router.Use(gin.Recovery(), gin.Logger())

	router.Use(TrustProxyHeaders())

	http.SetRoutes(router, a)

	router.Run(":8080")

}

func TrustProxyHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("X-Real-IP", ctx.GetHeader("X-Real-IP"))
		ctx.Set("X-Forwarded-For", ctx.GetHeader("X-Forwarded-For"))
		ctx.Set("X-Forwarded-Proto", ctx.GetHeader("X-Forwarded-Proto"))
		ctx.Next()
	}
}
