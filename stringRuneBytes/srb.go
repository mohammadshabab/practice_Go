package main

import (
	"fmt"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Convert the strings
//
//   1. Loop over the words slice
//
//   2. In the loop:
//      1. Convert each string value to a byte slice
//      2. Print the byte slice
//      3. Append the byte slice to the `bwords`
//
//   3. Print the words using the `bwords`
//
// EXPECTED OUTPUT
//  [103 111 112 104 101 114]
//  [112 114 111 103 114 97 109 109 101 114]
//  [103 111 32 108 97 110 103 117 97 103 101]
//  [103 111 32 115 116 97 110 100 97 114 100 32 108 105 98 114 97 114 121]
//  gopher
//  programmer
//  go language
//  go standard library
// ---------------------------------------------------------

func main() {
	// Please uncomment the code below

	words := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}

	var bwords [][]byte

	for _, i := range words {
		bw := []byte(i)
		fmt.Println(bw)

		bwords = append(bwords, bw)
	}

	for _, w := range bwords {
		fmt.Println(string(w))
	}

	// EXERCISE: Print the runes
	//
	//  1. Loop over the "console" word and print its runes one by one,
	//     in decimals, hexadecimals and binary.
	//
	//  2. Manually put the runes of the "console" word to a byte slice, one by one.
	//
	//     As the elements of the byte slice use only the rune literals.
	//
	//     Print the byte slice.
	//
	//  3. Repeat the step 2 but this time, as the elements of the byte slice,
	//     use only decimal numbers.
	//
	//  4. Repeat the step 2 but this time, as the elements of the byte slice,
	//     use only hexadecimal numbers.
	const word = "console"

	for _, i := range word {
		fmt.Printf("%d\t %x\t %b", i, i, i)
		fmt.Println()

	}

	words = []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, s := range words {
		fmt.Printf("%q\n", s)

		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("\thas %d bytes and %d runes\n",
			len(s), utf8.RuneCountInString(s))

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\tbytes   : % x\n", s)

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Print("\trunes   :")
		for _, r := range s {
			fmt.Printf("% x", r)
		}
		fmt.Println()

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		fmt.Print("\trunes   :")
		for _, r := range s {
			fmt.Printf("%q", r)
		}
		fmt.Println()

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("\tfirst   : %q (%d bytes)\n", r, size)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		r, size = utf8.DecodeLastRuneInString(s)
		fmt.Printf("\tlast    : %q (%d bytes)\n", r, size)

		// Slice and print the first two runes of the strings
		_, first := utf8.DecodeRuneInString(s)
		_, second := utf8.DecodeRuneInString(s[first:])
		fmt.Printf("\tfirst 2 : %q\n", s[:first+second])

		// Slice and print the last two runes of the strings
		_, last1 := utf8.DecodeLastRuneInString(s)
		_, last2 := utf8.DecodeLastRuneInString(s[:len(s)-last1])
		fmt.Printf("\tlast 2  : %q\n", s[len(s)-last2-last1:])

		// Convert the string to []rune
		// Print the first and last two runes
		rs := []rune(s)
		fmt.Printf("\tfirst 2 : %q\n", string(rs[:2]))
		fmt.Printf("\tlast 2  : %q\n", string(rs[len(rs)-2:]))

	}
}
