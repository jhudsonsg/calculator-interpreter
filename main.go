package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var expression string
var currentToken token
var result int

func evalExp() {
	parseExpression()
	evalExp2(&result)
}

func evalExp2(result *int) {
	var temp int
	var op string

	evalExp3(result)
	for {
		op = currentToken.value
		if op == "+" || op == "-" {
			parseExpression()
			evalExp3(&temp)

			switch op {
			case "-":
				*result = *result - temp
			case "+":
				*result = *result + temp
			}

		} else {
			break
		}
	}
}

func evalExp3(result *int) {
	var op string = ""

	if currentToken.typeValue == operator && (currentToken.value == "+" || currentToken.value == "-") {
		op = currentToken.value
		parseExpression()
	}

	evalExp4(result)

	if op == "-" {
		*result = -(*result)
	}
}

func evalExp4(result *int) {
	if currentToken.value == "(" {
		parseExpression()
		evalExp2(result)

		if currentToken.value != ")" {
			fmt.Println("Erro de sintaxe")
			os.Exit(1)
		}

		parseExpression()
	} else {
		parseValue(result)
	}
}

func parseValue(result *int) {
	if currentToken.typeValue == number {
		*result, _ = strconv.Atoi(currentToken.value)
		parseExpression()
	} else {
		fmt.Println("Erro de sintaxe")
		os.Exit(1)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		expression = scanner.Text()
		evalExp()
		fmt.Println(result)
	}
}
