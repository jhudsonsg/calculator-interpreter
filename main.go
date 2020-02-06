package main

import (
	"bufio"
	"fmt"
	"os"
)

var expression string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		expression = scanner.Text()
		fmt.Println(parseExpression())
		fmt.Println(parseExpression())
		fmt.Println(parseExpression())
		fmt.Println(expressionLimit)
	}
}
