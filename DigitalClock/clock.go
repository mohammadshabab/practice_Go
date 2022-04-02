package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	screen.Clear()

	for {
		screen.MoveTopLeft()

		now := time.Now()
		hour, minute, second := now.Hour(), now.Minute(), now.Second()
		//fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, minute, second)
		ssec := now.Nanosecond() / 1e8
		//fmt.Println("ssec  ", ssec)
	
		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[minute/10], digits[minute%10],
			colon,
			digits[second/10], digits[second%10],
			dot,
			digits[ssec],
		}

		fmt.Println()

		//alarmed := second%10 == 0

		for line := range clock[0] {
			// if alarmed {
			// 	clock = alarm
			// }
			for index, digit := range clock {
				next := clock[index][line]
				if (digit == colon || digit == dot) && second%2 == 0 {
					next = "   "
				}

				fmt.Print(next, " ")
			}
			fmt.Println()
		}
		fmt.Println()
		const splitSecond = time.Second / 10
		time.Sleep(splitSecond)
	}
}
