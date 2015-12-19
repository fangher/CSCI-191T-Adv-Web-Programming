package main
import "fmt"

type person struct {
	name string
	age int
}

func main () {
	jane := person{name: "Jane", age:25}
	john := person{name:"John"}

	fmt.Println(jane.name, jane.age)
	fmt.Println(john.name)

	john.age = 52
	fmt.Println(john.age)
}
