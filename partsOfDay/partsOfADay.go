package main

import (
	"fmt"
	"time"
)

func main() {

	switch hour := time.Now().Hour(); {
	case hour >= 18:
		fmt.Println("good evening!")
	case hour >= 12:
		fmt.Println("good Afternoon!")
	case hour >= 6:
		fmt.Println("Good morning!")
	default:
		fmt.Println("Good Night!")
	}

}
