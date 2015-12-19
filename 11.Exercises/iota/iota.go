package main
import "fmt"

func main () {
	const (
		a = iota
		b = iota
	)

	fmt.Println(a, b)
}
