package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var currentPositionExpression int = -1
var expressionLimit int = 0
var expressions []string

func parseExpression() Token {
	if len(expressions) == 0 {
		expressions = strings.Split(expression, " ")
		expressionLimit = len(expressions)
	}

	currentPositionExpression++
	value := expressions[currentPositionExpression]

	if match, _ := regexp.MatchString("[\\+\\-\\*\\/]", value); match == true {
		return Token{value, operator}
	} else if match, _ := regexp.MatchString("[0-9]", value); match == true {
		return Token{value, number}
	}

	fmt.Println("Erro lexico na espressao:", expression, "\nString com erro:", value)
	os.Exit(1)

	return Token{}
}
