package main

import (
	"fmt"
	"math/rand"
)

// rand.Intn(max-min+1) + min (whole number)
// rand.Float64()*(max-min) + min (float number)

func main() {

	for i := 1; i < 51; i++ {
		fmt.Println(i, rand.Intn(0+i)+1)
	}

}
