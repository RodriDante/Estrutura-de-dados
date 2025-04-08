package main

import "fmt"

func SomaAte(n int)int {
	soma := 0
	for i :=1; i<n; i++{
		soma = soma + i
	}
	return soma
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(SomaAte(n))
}

package main

import "fmt"

func ClassificarNota(nota int){
 switch{
	 case nota > 9:
	 fmt.Println("Excelente")
	 case nota >= 7 && nota < 9:
	 fmt.Println("Bom")
	 case nota >= 5 && nota < 7:
	 fmt.Println("Regular")
	 case nota >= 3 && nota <5:
	 fmt.Println("Ruim")
	 default:
	 fmt.Println("Muito ruim")
 }
	
}