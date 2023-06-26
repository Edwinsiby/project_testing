package handlers

import (
	"net/http"
	"test/domain"
	"test/repository"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var input domain.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.GetByName(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = repository.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "creating user failed"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "user created succesfuly"})
	}
}

func Login(c *gin.Context) {
	var input domain.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := repository.Login(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "user loged in succesfuly"})
	}

}
