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
func main(){
	var nota int
	fmt.Scan(&nota)
	ClassificarNota(nota)
}

