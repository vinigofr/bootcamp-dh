package main

import "os"

func readTXT(filepath string) error {
	_, err := os.Open(filepath)

	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
	return nil
}

func registerClient(arquivo, nome, sobrenome, rg, tel, endereco string) {}

func main() {
	readTXT("customers.txt")
}
