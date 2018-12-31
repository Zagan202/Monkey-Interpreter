/*
REPL reads input, sends it to the interpreter for evaluation, prints the result/output of the
interpreter and starts again. Read, Eval, Print, Loop.
*/

package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

//PROMPT is the cursor that the user inputs code after
const PROMPT = ">> "

/*
Start reads from the input source until encountering a newline, take
the just read line and pass it to an instance of our lexer and finally print all the tokens the lexer
gives us until we encounter EOF.
*/
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
