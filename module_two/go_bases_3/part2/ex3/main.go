package main

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

func SomarProdutos(produtos []Produto) (pTotal float64) {
	for _, product := range produtos {
		pTotal += product.preco * float64(product.quantidade)
	}
	return
}

func SomarServicos(servicos []Servico) (pTotal float64) {
	for _, servico := range servicos {
		if servico.minTrabalhados < 30 {
			pTotal += servico.preco * 30
		}
		pTotal += servico.preco * float64(servico.minTrabalhados)
	}
	return
}

func SomarManutencao(manutencoes []Manutencao) (pTotal float64) {
	for _, manutencao := range manutencoes {
		pTotal += manutencao.preco
	}
	return
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

	SomarProdutos()
	SomarServicos()
	SomarManutencao()
}
