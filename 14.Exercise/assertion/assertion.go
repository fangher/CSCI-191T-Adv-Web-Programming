package main
import "fmt"

func main () {
	var myInterface interface {} = 4
	fmt.Println(4 + myInterface.(int))
}
