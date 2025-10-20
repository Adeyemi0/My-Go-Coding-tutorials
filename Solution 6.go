package main

import (
	"fmt"
)

// rand.Intn(max-min+1) + min (whole number)
// rand.Float64()*(max-min) + min (float number)

func main() {

	fmt.Printf("Enter a length in centimeters: \n")
	var length float32
	fmt.Scanln(&length)
	if length < 0 {
		fmt.Println("You have entered a negative value which is invalid")
	} else if length >= 0 {
		fmt.Printf("The length in inches is %2.f", length*2.54)
	}

}
