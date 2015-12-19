package main
import "fmt"

func myFunc (a ...int) {
   fmt.Println (a)
}

func main () {
   myFunc (1, 2, 3, 4, 5)
   var a = [] int {1, 2, 3, 4, 5}
   myFunc (a...)
}
