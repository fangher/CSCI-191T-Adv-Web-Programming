package main
import "fmt"

func main () {
	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)
	
	secondSlice := []string{"something", "there"}
	fmt.Println(secondSlice)

	var thirdSlice []int = make([]int, 5, 10)
	thirdSlice[0] = 1
	thirdSlice[1] = 2
	thirdSlice[2] = 3
  	thirdSlice[3] = 4
  	thirdSlice[4] = 5
	fmt.Println(thirdSlice)
	fmt.Println(len(thirdSlice))
	fmt.Println(cap(thirdSlice))

	newSlice := []int{1, 2, 3}
	someSlice := []int{4, 5, 6}
	fmt.Println(append(newSlice, someSlice...))


	newSlice = append(newSlice[:3], newSlice[4:]...)
	fmt.Println(newSlice)


}
