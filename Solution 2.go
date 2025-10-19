package main

import (
	"fmt"
	"math"
	"math/rand"
)

// rand.Intn(max-min+1) + min

func main() {

	var x int = rand.Intn(50) + 1
	var y int = rand.Intn(4) + 2
	fmt.Printf("The value of x is %v and y is %v \n", x, y)
	fmt.Printf("%v raise to power %v is %v", x, y, math.Pow(float64(x), float64(y)))

}
