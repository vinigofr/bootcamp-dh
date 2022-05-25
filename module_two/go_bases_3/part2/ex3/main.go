package main

import (
	"fmt"
	"sync"
)

type Produto struct {
	nome       string
	preco      float64
	quantidade int
}

type Servico struct {
	nome           string
	preco          float64
	minTrabalhados int
}

type Manutencao struct {
	nome  string
	preco float64
}

var totalGlobal float64
var waitG sync.WaitGroup

func SomarProdutos(produtos []Produto) {
	for _, product := range produtos {
		fmt.Println("Somando produtos")
		totalGlobal += product.preco * float64(product.quantidade)
	}
	waitG.Done()
}

func SomarServicos(servicos []Servico) {
	for _, servico := range servicos {
		fmt.Println("Somando serviços")
		if servico.minTrabalhados < 30 {
			totalGlobal += servico.preco * 30
		}
		totalGlobal += servico.preco * float64(servico.minTrabalhados)
	}
	waitG.Done()
}

func SomarManutencao(manutencoes []Manutencao) {
	for _, manutencao := range manutencoes {
		fmt.Println("Somando manutenção")
		totalGlobal += manutencao.preco
	}
	waitG.Done()
}

func main() {
	var manutencoes = []Manutencao{
		Manutencao{nome: "Conserto impressora", preco: 50},
		Manutencao{nome: "Conserto pc", preco: 25},
		Manutencao{nome: "Conserto mouse", preco: 15},
		Manutencao{nome: "Montagem pc", preco: 150},
	}

	var servicos = []Servico{
		{nome: "Ex1", preco: 5, minTrabalhados: 50},
		{nome: "Ex2", preco: 5.25, minTrabalhados: 50},
		{nome: "Ex3", preco: 5.7, minTrabalhados: 150},
		{nome: "Ex4", preco: 5.96, minTrabalhados: 750},
	}

	var produtos = []Produto{
		{nome: "Pc", preco: 2500, quantidade: 5},
		{nome: "Mouse", preco: 250, quantidade: 5},
		{nome: "Teclado", preco: 350, quantidade: 5},
		{nome: "Monitor", preco: 1200, quantidade: 5},
		{nome: "Roteador", preco: 750, quantidade: 5},
	}

	waitG.Add(3)
	go SomarManutencao(manutencoes)
	go SomarServicos(servicos)
	go SomarProdutos(produtos)
	waitG.Wait()

	fmt.Println(totalGlobal)
}
