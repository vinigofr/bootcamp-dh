package exercise

import "fmt"

var temp float64
var moisture uint8
var hPa uint16

func Weather() {
	temp = 28.0
	moisture = 75
	hPa = 1013

	fmt.Println()
	fmt.Println("EXERCÍCIO 2")
	fmt.Printf("Temperatura %v, umidade %v, pressão atmosférica %v hpa", temp, moisture, hPa)
	fmt.Println()
}
