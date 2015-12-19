package main
import "fmt"

func main () {
	var a int = 10
	var b *int = &a

	fmt.Println("a's value - ", a)
	fmt.Println("a's address - ", &a)
	fmt.Println("b's value - ", b)
	fmt.Println("b points ", *b)
}
