package services

import (
	"github.com/rizkirchm97/go-moo-app/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	categoryRepository = &repositories.CategoryRepositoryMock{Mock: mock.Mock{}}
	categoryService    = CategoryService{Repository: categoryRepository}
)

func TestCategoryService_Get(t *testing.T) {

	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}
