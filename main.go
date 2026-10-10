package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to GoCalc")
	fmt.Println("-----------------------------------------")
	fmt.Println("Binary ops (e.g., 5 + 3, 2 ^ 3, 10 % 3): +, -, *, /, ^, %")
	fmt.Println("Unary ops  (e.g., sqrt 16, abs -5, sin 90): sqrt, abs, sin, cos, tan")
	fmt.Println("Type 'exit' to quit")

	scanner := bufio.NewScanner(os.Stdin)

	binaryRegex := regexp.MustCompile(`^\s*([+-]?\d*\.?\d+)\s*([\+\-\*\/\^\%])\s*([+-]?\d*\.?\d+)\s*$`)
	unaryRegex := regexp.MustCompile(`^\s*([a-zA-Z]+)\s+([+-]?\d*\.?\d+)\s*$`)

	for {
		fmt.Print("\n<GoCalc> ")
		
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "exit" || strings.ToLower(input) == "quit" {
			fmt.Println("Exiting...")
			break
		}

		if input == "" {
			continue
		}

		if unaryMatches := unaryRegex.FindStringSubmatch(input); unaryMatches != nil {
			fn := strings.ToLower(unaryMatches[1])
			val, err := strconv.ParseFloat(unaryMatches[2], 64)
			if err != nil {
				fmt.Println("Error: Invalid number provided.")
				continue
			}

			var res float64
			var uErr error

			switch fn {
			case "sqrt":
				if val < 0 {
					uErr = fmt.Errorf("cannot calculate square root of a negative number")
				} else {
					res = math.Sqrt(val)
				}
			case "abs":
				res = math.Abs(val)
			case "sin":
				res = math.Sin(val * math.Pi / 180)
			case "cos":
				res = math.Cos(val * math.Pi / 180)
			case "tan":
				res = math.Tan(val * math.Pi / 180)
			default:
				uErr = fmt.Errorf("unknown function '%s'", fn)
			}

			if uErr != nil {
				fmt.Printf("Error: %v\n", uErr)
			} else {
				fmt.Printf("Result: %g\n", res)
			}
			continue
		}

		binaryMatches := binaryRegex.FindStringSubmatch(input)
		if binaryMatches == nil {
			fmt.Println("Error: Invalid format. Try '5 + 3', '2 ^ 3', or 'sqrt 16'")
			continue
		}

		num1, err1 := strconv.ParseFloat(binaryMatches[1], 64)
		operator := binaryMatches[2]
		num2, err2 := strconv.ParseFloat(binaryMatches[3], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Error: Invalid numbers provided.")
			continue
		}

		var result float64
		var calcErr error

		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				calcErr = fmt.Errorf("division by zero is not allowed")
			} else {
				result = num1 / num2
			}
		case "^":
			result = math.Pow(num1, num2)
		case "%":
			result = math.Mod(num1, num2)
		default:
			calcErr = fmt.Errorf("unknown operator '%s'", operator)
		}

		if calcErr != nil {
			fmt.Printf("Error: %v\n", calcErr)
		} else {
			fmt.Printf("Result: %g\n", result)
		}
	}
}