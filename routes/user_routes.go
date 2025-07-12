package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sonu31/expreimnet-go-lang-with-mongoDb/controller"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users", controller.CreateUser)
	router.GET("/users", controller.GetUsers)
}
