package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Give me a month name ")
		return
	}

	if strings.EqualFold(os.Args[1], "january") || strings.EqualFold(os.Args[1], "march") || strings.EqualFold(os.Args[1], "may") || strings.EqualFold(os.Args[1], "july") || strings.EqualFold(os.Args[1], "august") || strings.EqualFold(os.Args[1], "october") || strings.EqualFold(os.Args[1], "december") {
		fmt.Printf("%q has 31 days.", os.Args[1])
	} else if strings.EqualFold(os.Args[1], "february") {
		fmt.Printf("%q has 28 days.", os.Args[1])
	} else if strings.EqualFold(os.Args[1], "april") || strings.EqualFold(os.Args[1], "june") || strings.EqualFold(os.Args[1], "september") || strings.EqualFold(os.Args[1], "november") {
		fmt.Printf("%q has 30 days.", os.Args[1])
	} else {
		fmt.Printf("%q is not an year ", os.Args[1])
	}
}
