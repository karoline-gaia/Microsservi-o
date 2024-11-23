package repositories

import "github.com/karoline-gaia/go-categories-mvc/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
}
