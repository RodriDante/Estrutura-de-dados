package main
import "fmt"
func main() {
  var b float64 = 23.5
  var p2 *float64 = &b
  fmt.Println(*p2)
  }
