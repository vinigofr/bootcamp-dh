package main

/*
	1. Uma funcionalidade para guardar um arquivo de texto com as seguintes informações de produtos:
	[ { id, preco, quantidade } ]
	1.1 Estes estarão separados por vírgula, pois será um CSV

	2. Os arquivos devem ser salvos no formato CSV como mencionado acima
*/

type product struct {
	id         int
	preco      float64
	quantidade int
	nome       string
}

var produtos []product

func main() {
	produtos = append(produtos, product{id: 1, preco: 6.64, quantidade: 10, nome: "Lanche"})
	produtos = append(produtos, product{id: 2, preco: 7.99, quantidade: 5, nome: "Suco"})
	produtos = append(produtos, product{id: 3, preco: 5.19, quantidade: 2, nome: ""})
	produtos = append(produtos, product{id: 3, preco: 5.19, quantidade: 2, nome: "Coisa"})

	CreateCSV(produtos, "xablau")
}