package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	var lines int

	for in.Scan() {
		lines++
	}
	fmt.Printf("There are %d lines ", lines)

}
