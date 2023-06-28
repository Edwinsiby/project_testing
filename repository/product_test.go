package repository_test

import (
	"test/domain"
	"test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProductRepository is a mock implementation of the ProductRepository interface
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetProductByName(product domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) CreateProduct(product domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) GetAllProduct() ([]domain.Product, error) {
	args := m.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (m *MockProductRepository) DeleteProduct(product domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func TestAddProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)

	testProduct := domain.Product{
		ProductName: "Shirt",
		Price:       1250,
		Catergory:   "men",
	}

	// Set up expectations
	mockRepo.On("GetProductByName", testProduct).Return(nil)
	mockRepo.On("CreateProduct", testProduct).Return(nil)
	mockRepo.On("DeleteProduct", testProduct).Return(nil)

	SetProductRepository(mockRepo)

	err := repository.GetProductByName(testProduct)
	assert.NoError(t, err, "Expected no error")

	err = repository.CreateProduct(testProduct)
	assert.NoError(t, err, "Expected no error")

	err = repository.DeleteProduct(testProduct)
	assert.NoError(t, err, "Expected no error")

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestGetAllProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)

	expectedProducts := []domain.Product{
		{ProductName: "T-shirt", Price: 1250.00, Catergory: "men"},
	}

	// Set up expectations
	mockRepo.On("GetAllProduct").Return(expectedProducts, nil)

	repository.SetProductRepository(mockRepo)

	list, err := repository.GetAllProduct()
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, expectedProducts, list, "Returned products do not match expected products")

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}
