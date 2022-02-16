package main

import "fmt"

func main() {
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	//Timezones (EST -5, MST -7, PST -8)
	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)

	fmt.Println(EST, MST, PST)

	const (
		Nov = 11 - iota // 11 - 0 = 11
		Oct             // 11 - 1 = 10
		Sep             // 11 - 2 = 9
	)

	fmt.Println(Sep, Oct, Nov)

	const (
		_   = iota // 0
		Jan        // 1
		Feb        // 2
		Mar        // 3
	)

	fmt.Println(Jan, Feb, Mar)

	const (
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)

}
