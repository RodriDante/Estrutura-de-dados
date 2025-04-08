package main

import (
	"fmt"
)

type Produto struct {
	Nome      string
	Preco     float64
	Quantidade int
}

func main() {
	var numProdutos int

	fmt.Print("Quantos produtos deseja cadastrar? ")
	fmt.Scan(&numProdutos)

	produtos := make([]Produto, numProdutos)


	for i := 0; i < numProdutos; i++ {
		fmt.Printf("\nProduto %d:\n", i+1)

		fmt.Print("Nome: ")
		fmt.Scan(&produtos[i].Nome)

		fmt.Print("Preço: ")
		fmt.Scan(&produtos[i].Preco)

		fmt.Print("Quantidade: ")
		fmt.Scan(&produtos[i].Quantidade)
	}

	
	var indice int
	fmt.Print("\nDigite o número do produto que deseja acessar (1 a ", numProdutos, "): ")
	fmt.Scan(&indice)

	
	if indice >= 1 && indice <= numProdutos {
		p := produtos[indice-1]
		fmt.Println("\nProduto acessado:")
		fmt.Printf("Nome: %s\n", p.Nome)
		fmt.Printf("Preço: R$ %.2f\n", p.Preco)
		fmt.Printf("Quantidade: %d unidades\n", p.Quantidade)
	} else {
		fmt.Println("Índice inválido!")
	}
}