package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// input go run str.go Hello Shabab "İNANÇ"

func main() {

	name := "İnanç"
	fmt.Println(utf8.RuneCountInString(name))

	msg := os.Args[1]
	l := len(msg)
	s := strings.Repeat("!", l) + msg + strings.Repeat("!", l)
	s = strings.ToUpper(s)
	fmt.Println(s)

	name2 := os.Args[2]

	msg2 := `hi ` + name2 + `!
how are you?`

	fmt.Println(msg2)

	msg3 := os.Args[3]

	nb := msg3 + strings.Repeat("!", utf8.RuneCountInString(msg3))

	fmt.Println(nb)

	msg4 := `
	
	The weather looks good.
I should go and play.
	`
	ss := strings.TrimSpace(msg4)
	fmt.Println(ss)

	names := "inanç           "
	names = strings.TrimRight(names, " ")
	lenth := utf8.RuneCountInString(names)
	fmt.Println(lenth)
}
