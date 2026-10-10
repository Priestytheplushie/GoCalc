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

const (
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Cyan   = "\033[36m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
)

func main() {
	fmt.Println("Welcome to GoCalc")
	fmt.Println("-----------------------------------------")
	fmt.Println("Binary ops (e.g., 5 + 3, 2 ^ 3, 10 % 3): +, -, *, /, ^, %")
	fmt.Println("Unary ops  (e.g., sqrt 16, abs -5, sin 90): sqrt, abs, sin, cos, tan")
	fmt.Println("Memory:     Type 'ans' to use your last result")
	fmt.Println("Type 'exit' to quit")

	scanner := bufio.NewScanner(os.Stdin)
	var lastResult float64 = 0

	binaryRegex := regexp.MustCompile(`^\s*([+-]?\d*\.?\d+|ans)\s*([\+\-\*\/\^\%])\s*([+-]?\d*\.?\d+|ans)\s*$`)
	unaryRegex := regexp.MustCompile(`^\s*([a-zA-Z]+)\s+([+-]?\d*\.?\d+|ans)\s*$`)

	for {
		fmt.Print(Yellow + "\n<GoCalc> " + Reset)
		
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "exit" || strings.ToLower(input) == "quit" {
			fmt.Println(Cyan + "Exiting... Goodbye!" + Reset)
			break
		}

		if input == "" {
			continue
		}

		parseVal := func(s string) (float64, error) {
			if strings.ToLower(s) == "ans" {
				return lastResult, nil
			}
			return strconv.ParseFloat(s, 64)
		}

		if unaryMatches := unaryRegex.FindStringSubmatch(input); unaryMatches != nil {
			fn := strings.ToLower(unaryMatches[1])
			val, err := parseVal(unaryMatches[2])
			if err != nil {
				fmt.Println(Red + "Error: Invalid number provided." + Reset)
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
				fmt.Println(Red + "Error: " + uErr.Error() + Reset)
			} else {
				lastResult = res
				fmt.Printf(Green+"Result: %g\n"+Reset, res)
			}
			continue
		}

		binaryMatches := binaryRegex.FindStringSubmatch(input)
		if binaryMatches == nil {
			fmt.Println(Red + "Error: Invalid format. Try '5 + 3', 'ans * 2', or 'sqrt 16'" + Reset)
			continue
		}

		num1, err1 := parseVal(binaryMatches[1])
		operator := binaryMatches[2]
		num2, err2 := parseVal(binaryMatches[3])

		if err1 != nil || err2 != nil {
			fmt.Println(Red + "Error: Invalid numbers provided." + Reset)
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
			fmt.Println(Red + "Error: " + calcErr.Error() + Reset)
		} else {
			lastResult = result
			fmt.Printf(Green+"Result: %g\n"+Reset, result)
		}
	}
}