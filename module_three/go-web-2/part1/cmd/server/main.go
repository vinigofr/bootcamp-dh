package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vinigofr/bootcamp-dh/module_three/go-web-2/cmd/server/controllers"
	"github.com/vinigofr/bootcamp-dh/module_three/go-web-2/product"
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

	repo := product.NewRepository()
	service := product.NewService(repo)
	productController := controllers.NewProduct(service)

	productsGroup := r.Group("/products")
	productsGroup.Use(ValidateToken())
	{
		productsGroup.POST("/", productController.Store())
		productsGroup.GET("/", productController.GetAll())
	}

	r.Run(":4000")
}
