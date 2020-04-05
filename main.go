package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func extractOperators(equation string) []string {
	fmt.Println(equation)
	splitOperators := strings.Fields(equation)
	return splitOperators
}

func isSymbol(item string) bool {
	symbols := [4]string{"+", "-", "/", "*"}
	for _, v := range symbols {
		if v == item {
			return true
		}
	}
	return false
}

func performOperations(operators []string) float64 {
	var total float64
	var lastSymbol string
	for i, v := range operators {
		if isSymbol(v) {
			lastSymbol = v
		} else {
			val, _ := strconv.ParseFloat(v, 64)
			switch lastSymbol {
			case "+":
				total += val
			case "-":
				total -= val
			case "*":
				total *= val
			case "/":
				total /= val
			default:
				total = val
			}
		}

		fmt.Println("Run: ", i, v, total, lastSymbol)
	}
	return total
}

func main() {
	fmt.Println("Welcome To Calculator")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter You Equation: ")
	text, _ := reader.ReadString('\n')
	if text == "\n" {
		fmt.Println("Nothing Entered")
	}
	splitOperators := extractOperators(text)
	for i, v := range splitOperators {
		fmt.Printf("%d - %s\n", i, v)
	}
	total := performOperations(splitOperators)
	fmt.Println("Result: ", total)
}
