package main

import (
	"github.com/gin-gonic/gin"
	"github.com/karoline-gaia/go-categories-mvc/cmd/api/controllers"
	"github.com/karoline-gaia/go-categories-mvc/internal/repositories"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("categories")

	InMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()
	categoryRoutes.POST("/", func (ctx *gin.Context) {
		controllers.CreateCategory(ctx, InMemoryCategoryRepository)
	})
}
