package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karoline-gaia/go-categories-mvc/internal/repositories"
	use_cases "github.com/karoline-gaia/go-categories-mvc/internal/use-cases"
)

func ListCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {

	useCase := use_cases.NewListCategoryUseCase(repository)

	categories, err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"sucess": false,
				"error":  err.Error(),
			})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"sucess": true, "categories": categories})
}
