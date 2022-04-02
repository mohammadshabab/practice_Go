package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const (
	header = "Location,Size,Beds,Baths,Price"
	data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

	separator = ","
)

func main() {
	png, header_1 := []byte{'P', 'N', 'G'}, []byte{}

	header_1 = append(header_1, png...)

	if bytes.Equal(png, header_1) {
		fmt.Println("They are equal")
	}

	toppings := []string{"black olives", "green peppers"}

	var pizza []string
	pizza = append(pizza, toppings...)
	pizza = append(pizza, "onions", "extra cheese")

	fmt.Printf("toppings: %s\n", pizza)

	var (
		locs                       []string
		sizes, beds, baths, prices []int
	)

	rows := strings.Split(data, "\n")

	for _, row := range rows {
		cols := strings.Split(row, separator)

		locs = append(locs, cols[0])

		n, _ := strconv.Atoi(cols[1])
		sizes = append(sizes, n)

		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)

		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)

		n, _ = strconv.Atoi(cols[4])
		prices = append(prices, n)
	}

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range rows {
		fmt.Printf("%-15s", locs[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}

	words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
	words = append(words[:1], "is", "everywhere")
	fmt.Println(words)
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))
	words = append(words, words[len(words)+1:cap(words)]...)
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))
	fmt.Println(words)

	lyric := []string{"show", "me", "my", "silver", "lining"}
	part := lyric[:2:2]
	fmt.Println(part)

	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	const niho = ""
	for index, runeValue := range niho {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

}
