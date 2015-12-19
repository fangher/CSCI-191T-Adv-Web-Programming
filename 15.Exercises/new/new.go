package main
import "fmt"

func main () {

	var a *int = new (int)
	fmt.Println(a)
	fmt.Println(*a)

	var b *string = new (string)
	fmt.Println(b)
	fmt.Println(*b)

	var c *bool = new (bool)
	fmt.Println(c)
	fmt.Println(*c)

	slice := make([]int, 3, 3)
	fmt.Println(slice)

	myMap := make(map[int]string)
	fmt.Println(myMap)

}
