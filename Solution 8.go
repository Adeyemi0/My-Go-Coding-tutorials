package main

import (
	"fmt"
	"math/rand/v2"
)

// rand.Intn(max-min+1) + min (whole number)
// rand.Float64()*(max-min) + min (float number)

func main() {

	fmt.Printf("Guess the number i am thinking, it is between 1 and 10: ")
	var num = rand.IntN(10) + 1
	var user int
	fmt.Scanln(&user)
	if user == num {
		fmt.Printf("You got it right the number is indeed %v", num)
	} else {
		fmt.Printf("You got it wrong the number is %v", num)

	}

}
