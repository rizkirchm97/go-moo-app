package services

import (
	"errors"
	"github.com/rizkirchm97/go-moo-app/entities"
	"github.com/rizkirchm97/go-moo-app/repositories"
)

type CategoryService struct {
	Repository repositories.CategoryRepository
}

func (service *CategoryService) Get(id string) (*entities.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category not found")
	} else {
		return category, nil
	}
}
