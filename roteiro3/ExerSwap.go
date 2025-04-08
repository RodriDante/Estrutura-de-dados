package main
import "fmt"
func main() {
var x int = 10
var y int = 20
  fmt.Println("Os valores de x e y:" , x, y)
  x, y = y, x
  fmt.Println("Os valores de x e y agora é:" , x, y)
}