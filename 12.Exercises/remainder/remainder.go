package main
import "fmt"

func main () {
	fmt.Println("Enter two numbers with the first value greater than the second value")
	var dividend, divisor int
	fmt.Scan(&dividend)
	fmt.Scan(&divisor)
	fmt.Println("Remainder is ", dividend % divisor)
}
