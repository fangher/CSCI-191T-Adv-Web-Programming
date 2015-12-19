package main
import "fmt"

func myFunc (x [] int) {
   fmt.Println (x)
}

func main () {
   var x = [] int {1, 2, 3, 4, 5}
   myFunc(x)
}
