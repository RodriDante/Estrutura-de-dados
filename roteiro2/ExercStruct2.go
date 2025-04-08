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

	
	fmt.Println("\nProdutos cadastrados:")
	for i, p := range produtos {
		fmt.Printf("\nProduto %d\n", i+1)
		fmt.Printf("Nome: %s\n", p.Nome)
		fmt.Printf("Preço: R$ %.2f\n", p.Preco)
		fmt.Printf("Quantidade: %d unidades\n", p.Quantidade)
	}
}
