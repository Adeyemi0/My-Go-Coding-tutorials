package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var list []int64
	fmt.Printf("Enter a list of numbers seperated by comma (e.g 1,2,3,4,5): ")
	user := bufio.NewReader(os.Stdin)
	input, _ := user.ReadString('\n')

	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")

	for _, p := range parts {
		n, err := strconv.ParseInt(p, 10, 64)
		if err == nil {
			list = append(list, n)
		} else {
			fmt.Printf("There was an error with %v could not convert\n", p)
		}

	}
	fmt.Printf("The array is ")
	fmt.Print(list)
	fmt.Printf("\n")
	fmt.Printf("The last item of the array is %v\n", list[len(list)-1])

	for i := len(list) - 1; i >= 0; i-- {
		fmt.Print(list[i], " ")
	}
	fmt.Printf("]\n")
	count := 0
	five := 0
	for _, v := range list {
		if v == 5 {
			five += 1
		}

		if v == 5 {
			count += 1
		}

	}
	if five > 0 {
		fmt.Printf("There is 5 in the array\n")
	} else {
		fmt.Printf("We do not have 5 in the array\n")
	}
	fmt.Printf("The number of times 5 appear in the array is %v\n", count)

	var e []int64
	for index, value := range list {

		if index != 0 && index != len(list)-1 {
			e = append(e, value)
		}
	}
	fmt.Printf("The array without the first and the last item is")
	fmt.Print(e)

}
