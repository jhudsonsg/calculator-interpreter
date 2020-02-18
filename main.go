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

func eval() {
	parseExpression()
	parseAdditionSubtraction(&result)
}

func parseAdditionSubtraction(result *int) {
	var temp int
	var op string

	parseMultiplicationDivision(result)
	for {
		op = currentToken.value
		if op == "+" || op == "-" {
			parseExpression()
			parseMultiplicationDivision(&temp)

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

func parseMultiplicationDivision(result *int) {
	var temp int
	var op string

	parseContinuationAdditionSubtraction(result)
	for {
		op = currentToken.value
		if op == "*" || op == "/" {
			parseExpression()
			parseContinuationAdditionSubtraction(&temp)

			switch op {
			case "*":
				*result = *result * temp
			case "/":
				*result = *result / temp
			}
		} else {
			break
		}
	}
}

func parseContinuationAdditionSubtraction(result *int) {
	var op string = ""

	if currentToken.typeValue == operator && (currentToken.value == "+" || currentToken.value == "-") {
		op = currentToken.value
		parseExpression()
	}

	parseSeparator(result)

	if op == "-" {
		*result = -(*result)
	}
}

func parseSeparator(result *int) {
	if currentToken.value == "(" {
		parseExpression()
		parseAdditionSubtraction(result)

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
		return
	}

	fmt.Println("Erro de sintaxe")
	os.Exit(1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		expression = scanner.Text()
		eval()
		fmt.Println(result)
	}
}
