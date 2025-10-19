package main

import (
	"fmt"
	"math/rand"
)

// rand.Intn(max-min+1) + min

func main() {

	var num int = rand.Intn(10) + 1
	fmt.Printf("The random number is %v \n", num)
	for i := 0; i < num; i++ {
		println(i+1, "Adeyemi")
	}

}
