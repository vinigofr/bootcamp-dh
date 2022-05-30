package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type transaction struct {
	Id              int     `json:"id"`
	TransactionCode int     `json:"transaction_code"`
	Currency        string  `json:"currency"`
	Value           float64 `json:"value"`
	Sender          string  `json:"sender"`
	Receiver        string  `json:"receiver"`
}

func QueryStringFilter(c *gin.Context) {
	data, err := ioutil.ReadFile("transactions.json")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro":        err.Error(),
			"mensagem":    "Ocorreu um erro ao resgatar as transações",
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	queryParams := c.Request.URL.Query()

	if len(queryParams) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "É necessário passar ao manos um parâmetro de query string via URL",
		})
		return
	}

	var transactions []transaction
	var filteredTransactions []transaction

	json.Unmarshal(data, &transactions)

	var (
		idValue              = c.Query("id")
		transactionCodeValue = c.Query("transaction_code")
		currencyValue        = c.Query("currency")
		senderValue          = c.Query("sender")
		receiverValue        = c.Query("receiver")
	)

	var (
		convertedId, errId                   = strconv.Atoi(idValue)
		convertedTransactionCodeValue, errTc = strconv.Atoi(transactionCodeValue)
	)

	if errId != nil {
		convertedId = 0
	}

	if errTc != nil {
		convertedTransactionCodeValue = 0
	}

	for _, anyTransaction := range transactions {
		if anyTransaction.TransactionCode == convertedTransactionCodeValue ||
			anyTransaction.Id == convertedId ||
			anyTransaction.Receiver == receiverValue ||
			anyTransaction.Sender == senderValue ||
			anyTransaction.Currency == currencyValue {
			filteredTransactions = append(filteredTransactions, anyTransaction)
		}
		fmt.Println(idValue == string(anyTransaction.Id))
	}

	if len(filteredTransactions) == 0 {
		// Retornando um array vazio em Golang JSON: https://stackoverflow.com/questions/56200925/return-an-empty-array-instead-of-null-with-golang-for-json-return-with-gin
		c.JSON(http.StatusNotFound, make([]string, 0))
		return
	}

	c.JSON(200, filteredTransactions)
	return
}

func GetAll(c *gin.Context) {
	data, err := ioutil.ReadFile("transactions.json")

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
	req.GET("/search", QueryStringFilter)
	req.Run()
}
