package http

import (
	"github.com/gin-gonic/gin"
	"github.com/msssrp/SportEquipmentBorrowing/pkg/app"
)

func NewServer(app *app.App) *gin.Engine {
	router := gin.Default()

	SetRoutes(router, app)

	return router
}
