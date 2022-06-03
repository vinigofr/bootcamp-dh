package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vinigofr/bootcamp-dh/module_three/go-web-2/part1/product"
	"net/http"
)

type ProductController struct {
	service product.Service
}

func NewProduct(p product.Service) *ProductController {
	return &ProductController{service: p}
}

func (p *ProductController) Store() gin.HandlerFunc {
	var req request
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Verifique os campos e tente novamente",
				"code":  http.StatusBadRequest,
			})
			return
		}

		if req.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "O campo 'name' não pode estar vazio",
				"code":    http.StatusBadRequest,
			})
			return
		}

		if req.Price <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "O campo 'price' precisa ser maior que zero",
				"code":    400,
			})
			return
		}

		// Se não der erro, armazena no código que ficará abaixo
		registeredProduct := p.service.Store(req.Name, req.Price)

		c.JSON(http.StatusOK, gin.H{
			"message": "produto registrado com sucesso",
			"data":    registeredProduct,
		})
		return
	}
}

func (p *ProductController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		allProducts := p.service.GetAll()

		if len(allProducts) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Infelizmente ainda não temos registros de dados me nosso banco",
				"code":    http.StatusNotFound,
			})
			return
		}

		c.JSON(http.StatusOK, allProducts)
	}
}

type request struct {
	Name string `json:"name"`

	Price int `json:"price"`
}
