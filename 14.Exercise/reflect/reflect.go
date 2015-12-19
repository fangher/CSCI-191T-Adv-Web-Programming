package main
import "fmt"
import "reflect"

func main () {
	myInt := 5
	fmt.Println(reflect.TypeOf(myInt))
	fmt.Printf("%T", myInt)
}
