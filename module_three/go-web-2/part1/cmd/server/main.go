package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vinigofr/bootcamp-dh/module_three/go-web-2/part1/cmd/server/controllers"
	product2 "github.com/vinigofr/bootcamp-dh/module_three/go-web-2/part1/product"
	"net/http"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		fmt.Println(token)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "É necessário informar um token para realizar operações",
				"code":  http.StatusBadRequest,
			})
			return
		}

		if token != "12345" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Token informado é incorreto",
				"code":  http.StatusUnauthorized,
			})
			fmt.Println("antes do return")
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()

	repo := product2.NewRepository()
	service := product2.NewService(repo)
	productController := controllers.NewProduct(service)

	productsGroup := r.Group("/products")
	productsGroup.Use(ValidateToken())
	{
		productsGroup.POST("/", productController.Store())
		productsGroup.GET("/", productController.GetAll())
	}

	r.Run(":4000")
}
