package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	gad := [...]string{}
	fmt.Printf("type: %T, value: %d", gad, len(gad))
	if len(os.Args) != 3 {
		fmt.Println("Usage: [your name] [positive][negative]")
		return
	}

	name := os.Args[1]
	moodType := os.Args[2]
	rand.Seed(time.Now().UnixNano())

	mood := [...][5]string{
		{"good 👍 ",
			"happy 😀 ",
			"awesome 😎 ",
			"nerd 🤓",
			"smiling 🥰",
		},
		{
			"sad 😥",
			"terrible 😩 ",
			"fearful 😨 ",
			"anguished 😧 ",
			"screaming in fear 😱 ",
		},
	}

	n := rand.Intn(len(mood[0]))

	var mi int
	if moodType != "positive" {
		mi = 1
	}
	fmt.Printf("%s feels %s\n", name, mood[mi][n])

}
