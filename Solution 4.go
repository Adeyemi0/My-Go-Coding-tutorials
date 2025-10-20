package main

import (
	"fmt"
	"math/rand"
)

// rand.Intn(max-min+1) + min (whole number)
// rand.Float64()*(max-min) + min (float number)

func main() {

	var num float32 = rand.Float32()*9 + 1
	fmt.Printf("The random number is %.2f \n", num)

}
