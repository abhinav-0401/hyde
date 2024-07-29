package main

import (
	"fmt"
	"github.com/abhinav-0401/hyde/repl"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Hyde programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start()
}
