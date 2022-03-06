package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("You need to provide your age to watch movie ")
		return
	}
	age, err := strconv.Atoi(os.Args[1])
	if err != nil || age < 0 {
		fmt.Printf("Wrong age %q\n", os.Args[1])
	} else if age > 17 {
		fmt.Println("R-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else if age < 13 {
		fmt.Println("PG-Rated")
	}

}
