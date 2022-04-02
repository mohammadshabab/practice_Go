package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide directory paths")
		return
	}

	var dirs []byte
	for _, dir := range args {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println(err)
			return
		}
		dirs = append(dirs, dir...)
		dirs = append(dirs, '\n')

		for _, file := range files {
			if file.IsDir() {
				dirs = append(dirs, '\t')
				dirs = append(dirs, file.Name()...)
				dirs = append(dirs, '/', '\n')
			}
		}
		dirs = append(dirs, '\n')

	}

	fmt.Printf("%s", dirs)

}
