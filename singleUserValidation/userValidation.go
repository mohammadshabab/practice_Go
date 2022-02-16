package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q "
	errPwd   = "Invalid password for %q "
	accessOk = "Access granted for %q "
	username = "Shabab"
	password = "abc123"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}

	if os.Args[1] != username {
		fmt.Printf(errUser, username)
	} else if os.Args[2] != password {
		fmt.Printf(errPwd, username)
	} else {
		fmt.Printf(accessOk, username)
	}

}
