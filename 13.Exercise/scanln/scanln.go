package main
import "fmt"

func main () {
	var num int
	fmt.Print ("Enter number - ")
	fmt.Scanln (&num)
	newNum := num + 10
	fmt.Println(num, "+ 10 =", newNum)
}
