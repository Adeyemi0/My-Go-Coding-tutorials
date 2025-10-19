

package main

import (
	"fmt"
	"math/rand"
)

// rand.Intn(max-min+1) + min

func main() {

	for i := 0; i < 50; i++ {
		fmt.Println(i+1, rand.Intn(4)+3)
	}
}
