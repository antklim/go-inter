package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/antklim/go-inter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Println("Please enter a command:")
	repl.Start(os.Stdin, os.Stdout)
}
