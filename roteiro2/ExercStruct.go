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
	var produto Produto

	
	fmt.Print("Digite o nome do produto: ")
	fmt.Scan(&produto.Nome)

	fmt.Print("Digite o preço do produto: ")
	fmt.Scan(&produto.Preco)

	fmt.Print("Digite a quantidade em estoque: ")
	fmt.Scan(&produto.Quantidade)

	
	fmt.Println("\nProduto cadastrado:")
	fmt.Printf("Nome: %s\n", produto.Nome)
	fmt.Printf("Preço: R$ %.2f\n", produto.Preco)
	fmt.Printf("Quantidade em estoque: %d\n", produto.Quantidade)
}
