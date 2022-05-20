package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := LoanReceiver(23, 2, 100)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}

func LoanReceiver(age, workYears int, salary float64) (string, error) {

	if age < 22 {
		return "", errorMessage("Idade precisa ser maior que 22 anos!")
	}

	if workYears < 1 {
		return "", errorMessage("Você precisa ter experiência superior a 1 ano!")
	}

	if salary > 100000 {
		return fmt.Sprintln("Olá, você obteve o benefício de empréstimo sem juros! :D"), nil
	}

	return fmt.Sprintln("Olá, você obteve o benefício de empréstimo\n"+
		"com uma taxa juros de 5%! O valor final é de", salary*0.05+salary), nil
}

func errorMessage(message string) error {
	return errors.New(message)
}
