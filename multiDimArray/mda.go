package main

import (
	"fmt"
	"strings"
)

func main() {

	scientiesName := [...][3]string{
		{"First Name", "Last Name", "Nick Name"},
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	for i := range scientiesName {
		n := scientiesName[i]
		fmt.Printf("%-15s %-15s %-15s\n", n[0], n[1], n[2])
		if i == 0 {
			fmt.Println(strings.Repeat("=", 45))
		}
	}

}
