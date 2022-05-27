package main

import (
	"errors"
	"fmt"
)

var salario int

type customError error

var salarioMinimoErr customError
var impostoNecessarioErr customError

func main() {
	salarioMinimoErr = errors.New("erro: o mínimo tributável é 15000")
	impostoNecessarioErr = errors.New("Necessário pagamento de imposto")

	salario = 14999
	if salario <= 15000 {
		err := fmt.Errorf("%s e o salário informado foi %d", salarioMinimoErr.Error(), salario)
		fmt.Println(err)
	} else {
		fmt.Println(impostoNecessarioErr.Error())
	}
}
