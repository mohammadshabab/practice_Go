package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	sort.Strings(args)

	//fmt.Println(args)

	var data []byte

	for i, s := range args {
		data = strconv.AppendInt(data, int64(i+1), 10)
		data = append(data, '.', ' ')
		data = append(data, s...)
		data = append(data, '\n')
	}

	err := ioutil.WriteFile("sortedNames.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
