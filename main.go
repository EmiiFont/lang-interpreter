package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/EmiiFont/lang-interpreter/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hellow %s! This is the Moneky programming language\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
