package main

import (
	"github.com/gin-gonic/gin"
	"github.com/karoline-gaia/go-categories-mvc/cmd/api/controllers"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("categories")
	categoryRoutes.POST("/", controllers.CreateCategory)
}
