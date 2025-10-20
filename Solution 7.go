package main

import (
	"fmt"
)

// rand.Intn(max-min+1) + min (whole number)
// rand.Float64()*(max-min) + min (float number)

func main() {

	fmt.Printf("How many credits have you taken: ")
	var credit int
	fmt.Scanln(&credit)
	if credit <= 23 {
		fmt.Printf("You're a freshman")
	} else if 24 >= credit && credit < 53 {
		fmt.Printf("You're a Sophomore.")
	} else if 54 >= credit && credit < 84 {
		fmt.Println("You're a Junior")
	} else {
		fmt.Println("You're a Senior")
	}
}
