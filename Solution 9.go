package main

import (
	"fmt"
	"math/rand/v2"
)

// rand.Intn(max-min+1) + min (whole number)
// rand.Float64()*(max-min) + min (float number)

func main() {

	for i := 0; i < 10; i++ {
		var x = rand.IntN(10) + 1
		var y = rand.IntN(10) + 1
		var user int
		fmt.Printf("Question %v:  %v x %v = ", i, x, y)
		fmt.Scanln(&user)
		if user != x*y {
			fmt.Printf("Wrong. The answer is %v \n", x*y)
		} else if user == x*y {
			fmt.Printf("Right!\n")
		}

	}

}
