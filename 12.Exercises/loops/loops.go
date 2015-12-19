package main
import "fmt"

func main () {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
		   fmt.Print (i, " ")
		}
	}

	fmt.Println()

	i := 0
	for i < 10 {
		if i % 2 == 0 {
		   fmt.Print (i, " ")
		}

		i++
	}

	fmt.Println()

	i = 0
	for {
		if i >= 10 {
			break
		}

		if i % 2 == 0 {
		   fmt.Print (i, " ")
		}

		i++
	}
}
