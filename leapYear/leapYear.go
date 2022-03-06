package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q, is not a valid year ", os.Args[1])
		return
	}

	if year%4 == 0 {
		if year%100 == 0 && year%400 == 0 {
			fmt.Printf("%d is a leap year ", year)
		} else if year%100 != 0 && year%4 == 0 {
			fmt.Printf("%d is a leap year ", year)
		} else {
			fmt.Printf("%d is not a leap year ", year)
		}
	} else {
		fmt.Printf("%d is not a leap year ", year)
	}

}
