package handlers

import (
	"net/http"
	"test/domain"
	"test/repository"

	"github.com/gin-gonic/gin"
)

var productRepo repository.ProductRepository

func SetProductRepository(repo repository.ProductRepository) {
	productRepo = repo
}

func AddProduct(c *gin.Context) {
	var input domain.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := productRepo.GetProductByName(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = productRepo.CreateProduct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "creating product failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "product created successfully"})
}

func ListProducts(c *gin.Context) {
	products, err := productRepo.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Products": products})
}
