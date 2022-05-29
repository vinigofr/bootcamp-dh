package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type transaction struct {
	Id              int     `json:"id"`
	TransactionCode int     `json:"transaction_code"`
	Currency        string  `json:"currency"`
	Value           float64 `json:"value"`
	Sender          string  `json:"sender"`
	Receiver        string  `json:"receiver"`
}

func GetAll(c *gin.Context) {

	data, err := ioutil.ReadFile("tranfdsfsd" +
		"sactions.json")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro":        err.Error(),
			"mensagem":    "Ocorreu um erro ao resgatar as transações",
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	var transactions []transaction

	json.Unmarshal(data, &transactions)

	c.JSON(http.StatusOK, transactions)
}

func helloWorld(c *gin.Context) {
	username := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"messagem": "Olá, " + username,
	})
}

func main() {
	req := gin.Default()

	// Referência para usar string params
	// https://stackoverflow.com/questions/39489175/how-to-define-a-go-gin-route-with-an-id-in-the-middle
	req.GET("/hi/:name", helloWorld)
	req.GET("/transactions", GetAll)
	req.Run()
}
