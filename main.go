package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/abhinav-0401/hyde/lexer"
	"github.com/abhinav-0401/hyde/util"
)

func main() {
	var srcFile = flag.String("src", "", "source script file name")
	flag.Parse()

	if *srcFile == "" {
		runPrompt()
	} else {
		runFile(*srcFile)
	}
}

func runPrompt() {
	var reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		var line, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading input from REPL: ", err)
		}

		run(line)
		util.StateFlags.HadErr = false
	}
}

func runFile(srcFile string) {
	var src, err = os.ReadFile(srcFile)
	if err != nil {
		log.Fatal("Error reading input from the file: ", err)
	}

	run(string(src))

	if util.StateFlags.HadErr {
		os.Exit(65)
	}
}

func run(s string) {
	var lex = lexer.New(s)
	lex.LexTokens()

	for _, v := range lex.Tokens {
		fmt.Printf("%v %v %v\n", v.Kind, v.Lexeme, v.Line)
	}
}
