package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurn = 5
	Usage   = `Welcome to the Lucky Number Game!  🍀
	 
	The program will pick %d random numbers.
	Your misson is to guess one of those numbers.
	
	The greater your number is,the harder it gets.

	Wanna play?
	`
)

func main() {

	rand.Seed(time.Now().UnixNano())
	if len(os.Args) != 2 {
		fmt.Printf(Usage, maxTurn)
		return
	}
	num := os.Args[1]
	guess, err := strconv.Atoi(num)
	if err != nil {
		fmt.Printf("%s is not a number", num)
		return
	}
	if guess < 0 {
		fmt.Println("Please provide a positive number.")
		return
	}

	for turn := 0; turn < maxTurn; turn++ {
		n := rand.Intn(guess) + 1

		if n == guess {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("🎉  YOU WIN!")
			case 1:
				fmt.Println("🎉  YOU'RE AWESOME!")
			case 2:
				fmt.Println("🎉  PERFECT!")
			}
			return
		}

	}

	msg := "%s Try again?\n"

	switch rand.Intn(2) {
	case 0:
		fmt.Printf(msg, "☠️  YOU LOST...")
	case 1:
		fmt.Printf(msg, "☠️  JUST A BAD LUCK...")
	}
}
