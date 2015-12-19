package main
import "fmt"

func main () {
	sum := 0
	for i := 3; i < 1000; i += 3 {
	   sum += i
	}

	for i := 5; i < 1000; i += 5 {
	   sum += i
	}
	fmt.Print(sum)
}
