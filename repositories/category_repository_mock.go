package repositories

import (
	"github.com/rizkirchm97/go-moo-app/entities"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entities.Category {
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		category := args.Get(0).(entities.Category)
		return &category
	}
}
