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
	Usage   = `
	Welcome to the Lucky Number Game!  🍀
	 
	The program will pick %d random numbers.
	Your misson is to guess one of those numbers.
	
	The greater your number is,the harder it gets.
	You will get special bonus if won in first attempt.

	Wanna play?
	`
)

func main() {

	rand.Seed(time.Now().UnixNano())
	if len(os.Args) != 2 {
		fmt.Printf(Usage, maxTurn)
		return
	}

	num := os.Args[1:]
	guess, err := strconv.Atoi(num[0])
	if err != nil {
		fmt.Printf("%s is not a number.", num[0])
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 1; turn <= maxTurn; turn++ {
		n := rand.Intn(guess) + 1
		if n != guess {
			continue
		}
		if turn == 1 {
			fmt.Println("🥇 FIRST TIME WINNER!!!")
		} else {
			fmt.Println("🎉  YOU WON!")
		}
		return
	}
	fmt.Println("☠️  YOU LOST... Try again?")

}
