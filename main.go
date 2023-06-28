package main

import (
	"test/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/usersignup", handlers.Signup)
	router.GET("/userlogin", handlers.Login)
	router.POST("/addproduct", handlers.AddProduct)
	router.GET("/listproducts", handlers.ListProducts)

	router.Run(":8080")

}
