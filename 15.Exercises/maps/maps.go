package main
import "fmt"

func main () {
	myMap := map[string]string {
		"Me":"Fang",
		"He":"Him",
		"She":"Her",
	}

	myMap["Something"] = "Stuff"
	myMap["Something"] = "More Stuff"
	delete(myMap, "Something")

	for key, val := range myMap {
		fmt.Println (key, " - ", val)
	}
	fmt.Println(len(myMap))

	if val, ok := myMap["He"]; ok {
		fmt.Println(val)
	}
}
