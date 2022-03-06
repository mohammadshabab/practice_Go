package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) != 2 {
		fmt.Println("[your name]")
		return
	}

	name := os.Args[1]

	moods := [...]string{
		"sad 😥",
		"good 👍 ",
		"bad 👎 ",
		"happy 😀 ",
		"awesome 😎 ",
		"terrible 😩 ",
		"fearful 😨 ",
		"anguished 😧 ",
		"screaming in fear 😱 ",
	}

	mood := rand.Intn(len(moods))
	fmt.Printf("%s feels %s\n", name, moods[mood])

	////or
	// switch rand.Intn(10) {
	// case 0:
	// 	fmt.Printf("%s feels sad 😥 ", name)
	// case 1:
	// 	fmt.Printf("%s feels good 👍 ", name)
	// case 2:
	// 	fmt.Printf("%s feels bad 👎 ", name)
	// case 3:
	// 	fmt.Printf("%s feels happy 😀 ", name)
	// case 4:
	// 	fmt.Printf("%s feels awesome 😎 ", name)
	// case 5:
	// 	fmt.Printf("%s feels terrible 😩 ", name)
	// case 6:
	// 	fmt.Printf("%s feels fearful 😨 ", name)
	// case 7:
	// 	fmt.Printf("%s feels anguished 😧 ", name)
	// case 8:
	// 	fmt.Printf("%s feels screaming in fear 😱 ", name)
	// }
}
