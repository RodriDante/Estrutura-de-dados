package main

import (
	"fmt"
)

func main() {
	const linhas = 3
	const colunas = 3

	var alunos [linhas][colunas]string

	// Inserção dos dados
	for i := 0; i < linhas; i++ {
		fmt.Printf("Aluno %d:\n", i+1)

		fmt.Print("Digite o índice: ")
		fmt.Scan(&alunos[i][0])

		fmt.Print("Digite o RA: ")
		fmt.Scan(&alunos[i][1])

		fmt.Print("Digite o nome: ")
		fmt.Scan(&alunos[i][2])

		fmt.Println()
	}

	
	fmt.Println("Dados dos alunos:")
	for i := 0; i < linhas; i++ {
		fmt.Printf("Aluno %d - Índice: %s | RA: %s | Nome: %s\n", i+1, alunos[i][0], alunos[i][1], alunos[i][2])
	}
}

