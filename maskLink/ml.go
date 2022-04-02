package main

import (
	"fmt"
	"os"
)

const (
	link  = "http://"
	nLink = len(link)
	mask  = '*'
)

var in bool

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Please provide a string having url to mask")
		return
	}
	args := os.Args[1:]

	byteSlice := make([]byte, 0, len(args))

	text := args[0]
	//fmt.Println("text ", text)

	for i := 0; i < len(text); i++ {

		if len(text[i:]) >= nLink && text[i:i+nLink] == link {
			in = true
			byteSlice = append(byteSlice, link...)
			i += nLink
		}
		c := text[i]

		switch c {
		case ' ', '\n', '\t':
			in = false
		}

		if in {
			c = mask
		}

		byteSlice = append(byteSlice, c)

	}

	fmt.Println(string(byteSlice))
}
