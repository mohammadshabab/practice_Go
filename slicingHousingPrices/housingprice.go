package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	const (
		data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	rows := strings.Split(data, "\n")
	//nSlice := strings.Split(dataSlice[])
	//	fmt.Println(rows)

	col := strings.Split(rows[0], separator)

	//fmt.Println(col)
	from, to := 0, len(col)

	args := os.Args[1:]

	for i, v := range col {
		l := len(args)

		if l >= 1 && v == args[0] {
			from = i
		}
		if l == 2 && v == args[1] {
			to = i + 1
		}
	}

	if from > to {
		from = 0
	}

	for i, row := range rows {
		cols := strings.Split(row, separator)

		for _, h := range cols[from:to] {
			fmt.Printf("%-15s", h)
		}
		fmt.Println()

		if i == 0 {
			fmt.Println()
		}

	}
}
