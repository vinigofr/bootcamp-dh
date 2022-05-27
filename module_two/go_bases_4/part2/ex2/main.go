package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type client struct {
	arquivo, nome, sobrenome, rg, tel, endereco string
}

func recoverFunc() {
	err := recover()
	fmt.Println(err)
}

func verifyIfExists(filepath string, currentId string) {
	defer recoverFunc()

	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		panic(err)
	}

	parsedData := strings.Split(string(data), "\n")

	for _, line := range parsedData {
		if strings.Contains(line, currentId) {
			panic("Já existe um usuário com este ID no banco")
		}
	}
}

func registerClient(c client, filepath string) (err error) {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic("erro ao abrir o arquivo")
	}

	defer file.Close()

	formatedUser := fmt.Sprintf("id: %s, nome: %s, sobrenome: %s, rg: %s, tel: %s, endereco: %s\n",
		c.arquivo,
		c.nome,
		c.sobrenome,
		c.rg,
		c.tel,
		c.endereco)

	verifyIfExists(filepath, c.arquivo)

	_, err = file.WriteString(formatedUser)

	if err != nil {
		panic("Ocorreu um erro ao gravar dados no arquivo")
	}

	return nil
}

func generateID() string {
	var possibilities = "!@#$%^&*()1234567890ABCDEFGHIJKLMNOPQRSTUVXZWYabcdefghijklmnopqrstuvxzkwy"

	var token []string
	var idSize = 16

	// Referência: https://stackoverflow.com/questions/45753397/why-this-repeats-the-same-random-number
	for i := 0; i < idSize; i++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(1)
		currentCharPosition := rand.Intn(len(possibilities))
		token = append(token, string(possibilities[currentCharPosition]))
		fmt.Println(token)
	}

	return strings.Join(token, "")
}

func main() {
	var cliente1 client = client{
		arquivo:   "12312312312312",
		rg:        "123",
		nome:      "vini",
		tel:       "99",
		endereco:  "av",
		sobrenome: "gouveia",
	}

	registerClient(cliente1, "customers.txt")
	return
}
