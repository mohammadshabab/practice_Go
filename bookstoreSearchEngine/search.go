package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Tell me a book title")
		return
	}

	bookTitles := [...]string{"Kafka's Revenge", "Stay Golden", "Everythingship", "Kafka's Revenge 2nd Edition"}

	searchQuery := os.Args[1]

	var found bool
	fmt.Println("Search Results:")
	for _, book := range bookTitles {
		searchQuery, book = strings.ToLower(searchQuery), strings.ToLower(book)
		if strings.Contains(book, searchQuery) {
			fmt.Printf("+%v\n", book)
			found = true
		}
	}
	if !found {
		fmt.Printf("We don't have the book %q\n", searchQuery)
	}

}
