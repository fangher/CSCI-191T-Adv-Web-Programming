package main
import "fmt"

func main () {
	myInt := 2
	fmt.Println(int64(myInt))

	myFloat := 3.14
	fmt.Println(int(myFloat))

	myBytes := []byte{'h','e','l','l','o'}
	fmt.Println(string(myBytes))

	myString := "hello world"
	fmt.Println([]byte(myString))
	
	fmt.Println(string('a'))
	fmt.Println(string([]byte{'h','e','l','l','o'}))
	fmt.Println([]byte("Hello"))
	fmt.Println(float64(12))
	fmt.Println(int(12.1230123))
}
