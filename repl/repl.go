package repl

import (
	"bufio"
	"fmt"
	"io"
	"lise/lexer"
	"lise/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
	
		scanned := scanner.Scan()
		if !scanned || scanner.Err() != nil {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
