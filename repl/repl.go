package repl

import (
	"bufio"
	"fmt"
	"github.com/abhinav-0401/hyde/lexer"
	"github.com/abhinav-0401/hyde/token"
	"os"
)

func Start() {
	var scanner = bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(">> ")

		var scanned = scanner.Scan()
		if !scanned {
			return
		}
		var line = scanner.Text()
		var l = lexer.New(line)

		for tok := l.LexToken(); tok.Kind != token.Eof; tok = l.LexToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
