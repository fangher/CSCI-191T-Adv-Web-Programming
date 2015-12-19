package main
import "fmt"

func main () {
	var x uint = 5
	var y uint = 9
	fmt.Println("x - ", x)
	fmt.Println("y - ", y)
	x = 1 << x
	y = y >> 1
	fmt.Println("x becomes ", x)
	fmt.Println("y becomes ", y)
}
