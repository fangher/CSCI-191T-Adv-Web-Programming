package main
import (
	"fmt"
	"math"
)

func main () {
	var x float64
	fmt.Println ("Enter float number")
	fmt.Scanln (&x)
	fmt.Println(math.Ceil(x))
}
