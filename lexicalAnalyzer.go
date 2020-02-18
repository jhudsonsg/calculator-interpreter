package main

import (
	"fmt"
	"os"
	"regexp"
)

var currentPositionExpression int = 0

func getCaracter() string {
	var value string

	for {
		value = string(expression[currentPositionExpression])
		if value != " " {
			break
		}
		currentPositionExpression++
	}

	return value
}

func parseExpression() {
	if len(expression) == currentPositionExpression {
		currentToken = token{}
		return
	}

	value := getCaracter()

	if match := isOperator(value); match == true {
		currentToken = token{value, operator}
	} else if match := isNumber(value); match == true {
		complementNumber := value

		for {
			if len(expression)-1 == currentPositionExpression {
				break
			}

			currentPositionExpression++
			value := getCaracter()

			if match := isNumber(value); match == true {
				complementNumber = complementNumber + value
			} else {
				currentPositionExpression--
				break
			}
		}

		currentToken = token{complementNumber, number}
	} else if match := isSeparator(value); match == true {
		currentToken = token{value, separator}
	} else {
		fmt.Println("Erro lexico na espressao:", expression, "\nString com erro:", value)
		os.Exit(1)
	}

	currentPositionExpression++
}

func isOperator(str string) bool {
	match, _ := regexp.MatchString("[\\+\\-\\*\\/]", str)
	return match
}

func isNumber(str string) bool {
	match, _ := regexp.MatchString("[0-9]", str)
	return match
}

func isSeparator(str string) bool {
	match, _ := regexp.MatchString("[\\(\\)]", str)
	return match
}
