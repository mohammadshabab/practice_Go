package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5
	Usage    = `Welcome to the Lucky Number Game!  🍀
	 
	The program will pick %d random numbers.
	Your misson is to guess one of those numbers.
	
	The greater your number is,the harder it gets.

	Wanna play?
	`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// if len(os.Args) != 2 {
	// 	fmt.Printf("You need to provide a lucky number ")
	// 	return
	// }
	num := os.Args[1:]
	if len(os.Args) != 2 {
		fmt.Printf(Usage, maxTurns)
		return
	}
	guess, err := strconv.Atoi(num[0])
	if err != nil {
		fmt.Printf("%s is not a number ", num)
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		if n == guess {
			fmt.Println("🎉 You WIN ")
			return
		}
	}
	fmt.Println("☠ YOU LOST... Try again?")
}
