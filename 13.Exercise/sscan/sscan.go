package main
import "fmt"

func main () {
	someSentence := "3 4"
	var three, four int
	fmt.Sscan(someSentence, &three, &four)
	fmt.Println(three, four)

	someSentence = "again"
	var someWords []string
	num, err := fmt.Sscan(someSentence, someWords...)
	fmt.Println(num, err)
	fmt.Println(someWords)
}
