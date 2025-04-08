  
package main

import "fmt"

func verif(n int)string {
	
	if n%2 == 0 {
		return "O número é par"

	}else {
		return "O número é impar"
		}
}
func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(verif(n))
}