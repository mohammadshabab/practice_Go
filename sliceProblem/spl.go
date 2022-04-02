package main

import "fmt"

func main() {
	// DON'T TOUCH THE FOLLOWING CODE
	nums := []int{56, 89, 15, 25, 30, 50}

	var mine []int
	mine = append(mine, nums[:3]...)
	mine[0], mine[1], mine[2] = -50, -100, -150
	mine = mine[:3]

	fmt.Println("Mine         :", mine)
	fmt.Println("Original nums:", nums[:3])
}
