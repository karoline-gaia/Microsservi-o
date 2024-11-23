package use_cases

import (
	"log"

	"github.com/karoline-gaia/go-categories-mvc/internal/entities"
)

type createCategoryUseCase struct {
	//db
}

func NewCreateCategoryUseCase() *createCategoryUseCase {
	return &createCategoryUseCase{}

}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	//TODO: persist.entity.to.db
	log.Println(category)
	return nil
}
