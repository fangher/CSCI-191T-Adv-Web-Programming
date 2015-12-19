package main

import "fmt"
import "strconv"

func main() {
	myNum := 42
	fmt.Println("Something about Go Lang", strconv.Itoa(myNum))

	myString := "42"
	num, _ := strconv.Atoi(myString)
	fmt.Println(num + 1000)
}
