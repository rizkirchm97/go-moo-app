package repositories

import "github.com/rizkirchm97/go-moo-app/entities"

type CategoryRepository interface {
	FindById(id string) *entities.Category
}
