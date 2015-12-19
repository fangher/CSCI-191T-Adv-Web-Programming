package main
import "fmt"

func main () {
	var user string

	fmt.Print ("Enter name - ")
	fmt.Scan (&user)
	fmt.Println ("Name is " + user)
}
