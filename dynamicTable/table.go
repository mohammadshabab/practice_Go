package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a size of a table like axa")
		return
	}
	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%s is not a number.", os.Args[1])
	}
	fmt.Printf("%5s", "X")
	for i := 0; i < max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i < max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j < max; j++ {
			fmt.Printf("%5d", i*j)
		}

		fmt.Println()
	}
}
