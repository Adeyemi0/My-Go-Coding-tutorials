package main

import (
	"fmt"
	"math/rand/v2"
)

// rand.Intn(max-min+1) + min (whole number)
// rand.Float64()*(max-min) + min (float number)

func main() {
	var x []int
	for i := 1; i < 21; i++ {
		var num int = rand.IntN(100) + 1
		x = append(x, num)
	}
	fmt.Println(x)
}
