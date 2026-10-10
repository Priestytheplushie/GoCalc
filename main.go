package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
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

func factorial(n int) (float64, error) {
	if n < 0 || n > 170 {
		return 0, fmt.Errorf("number must be between 0 and 170")
	}
	res := 1.0
	for i := 2; i <= n; i++ {
		res *= float64(i)
	}
	return res, nil
}

func main() {
	fmt.Println("Welcome to GoCalc")
	fmt.Println("-----------------------------------------")
	fmt.Println("Binary ops:  +, -, *, /, ^, %  (e.g., 5 + 3)")
	fmt.Println("Unary ops:   sqrt, abs, sin, cos, tan (e.g., sqrt 16)")
	fmt.Println("Percentages: pct [num] [percent], tip [bill] [tip%]")
	fmt.Println("Statistics:  stats, mean, sum, median, min, max [nums...]")
	fmt.Println("Probability: fact [n]        (e.g., fact 5)")
	fmt.Println("             ncr [n] [r]     (Combinations: n choose r)")
	fmt.Println("             npr [n] [r]     (Permutations)")
	fmt.Println("             prob [fav] [tot] (Success odds percentage)")
	fmt.Println("Memory:      Type 'ans' to use your last result")
	fmt.Println("Type 'exit' to quit")
	fmt.Println("-----------------------------------------")

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

		parseList := func(fields []string) ([]float64, error) {
			var nums []float64
			for _, f := range fields {
				cleaned := strings.TrimRight(f, ",")
				if cleaned == "" {
					continue
				}
				val, err := parseVal(cleaned)
				if err != nil {
					return nil, fmt.Errorf("invalid number: %s", cleaned)
				}
				nums = append(nums, val)
			}
			return nums, nil
		}

		fields := strings.Fields(input)
		if len(fields) == 0 {
			continue
		}
		command := strings.ToLower(fields[0])

		if command == "fact" || command == "ncr" || command == "npr" || command == "prob" {
			if command == "fact" {
				if len(fields) != 2 {
					fmt.Println(Red + "Error: Format is 'fact [n]' (e.g., fact 5)" + Reset)
					continue
				}
				nVal, err := parseVal(fields[1])
				if err != nil {
					fmt.Println(Red + "Error: Invalid number." + Reset)
					continue
				}
				res, pErr := factorial(int(nVal))
				if pErr != nil {
					fmt.Println(Red + "Error: " + pErr.Error() + Reset)
				} else {
					lastResult = res
					fmt.Printf(Green+"Factorial: %g\n"+Reset, res)
				}
				continue
			}

			if command == "ncr" || command == "npr" {
				if len(fields) != 3 {
					fmt.Printf(Red+"Error: Format is '%s [n] [r]'\n"+Reset, command)
					continue
				}
				n, err1 := parseVal(fields[1])
				r, err2 := parseVal(fields[2])
				if err1 != nil || err2 != nil {
					fmt.Println(Red + "Error: Invalid numbers provided." + Reset)
					continue
				}
				ni, ri := int(n), int(r)
				if ri > ni || ri < 0 || ni < 0 {
					fmt.Println(Red + "Error: Invalid parameters (ensure 0 <= r <= n)." + Reset)
					continue
				}

				fn1, err := factorial(ni)
				if err != nil {
					fmt.Println(Red + "Error: " + err.Error() + Reset)
					continue
				}
				fn2, _ := factorial(ri)
				fn3, _ := factorial(ni - ri)

				var res float64
				if command == "ncr" {
					res = fn1 / (fn2 * fn3)
				} else {
					res = fn1 / fn3
				}

				lastResult = res
				fmt.Printf(Green+"Result (%s): %g\n"+Reset, strings.ToUpper(command), res)
				continue
			}

			if command == "prob" {
				if len(fields) != 3 {
					fmt.Println(Red + "Error: Format is 'prob [favorable] [total]' (e.g., prob 3 10)" + Reset)
					continue
				}
				fav, err1 := parseVal(fields[1])
				tot, err2 := parseVal(fields[2])
				if err1 != nil || err2 != nil || tot == 0 {
					fmt.Println(Red + "Error: Invalid numbers or total cannot be zero." + Reset)
					continue
				}
				pct := (fav / tot) * 100.0
				lastResult = pct
				fmt.Printf(Green+"Probability: %g%% (Odds: %g out of %g)\n"+Reset, pct, fav, tot)
				continue
			}
		}

		if command == "stats" || command == "mean" || command == "sum" || command == "median" || command == "min" || command == "max" {
			if len(fields) < 2 {
				fmt.Println(Red + "Error: Please provide numbers (e.g., stats 10, 20, 30)" + Reset)
				continue
			}

			nums, err := parseList(fields[1:])
			if err != nil {
				fmt.Println(Red + "Error: " + err.Error() + Reset)
				continue
			}
			if len(nums) == 0 {
				fmt.Println(Red + "Error: No valid numbers found." + Reset)
				continue
			}

			sum := 0.0
			minVal := nums[0]
			maxVal := nums[0]
			for _, n := range nums {
				sum += n
				if n < minVal {
					minVal = n
				}
				if n > maxVal {
					maxVal = n
				}
			}
			mean := sum / float64(len(nums))

			sorted := make([]float64, len(nums))
			copy(sorted, nums)
			sort.Float64s(sorted)
			var median float64
			mid := len(sorted) / 2
			if len(sorted)%2 == 0 {
				median = (sorted[mid-1] + sorted[mid]) / 2.0
			} else {
				median = sorted[mid]
			}

			lastResult = mean

			switch command {
			case "sum":
				lastResult = sum
				fmt.Printf(Green+"Sum: %g\n"+Reset, sum)
			case "mean":
				fmt.Printf(Green+"Mean (Average): %g\n"+Reset, mean)
			case "median":
				lastResult = median
				fmt.Printf(Green+"Median: %g\n"+Reset, median)
			case "min":
				lastResult = minVal
				fmt.Printf(Green+"Min: %g\n"+Reset, minVal)
			case "max":
				lastResult = maxVal
				fmt.Printf(Green+"Max: %g\n"+Reset, maxVal)
			case "stats":
				fmt.Println(Cyan + "--- Statistical Summary ---" + Reset)
				fmt.Printf("Count:  %d\n", len(nums))
				fmt.Printf("Sum:    %g\n", sum)
				fmt.Printf("Mean:   %g\n", mean)
				fmt.Printf("Median: %g\n", median)
				fmt.Printf("Min:    %g\n", minVal)
				fmt.Printf("Max:    %g\n", maxVal)
			}
			continue
		}

		if command == "pct" || command == "tip" {
			if len(fields) != 3 {
				fmt.Println(Red + "Error: Format is 'pct [num] [percent]' or 'tip [bill] [tip%]'" + Reset)
				continue
			}
			val1, err1 := parseVal(fields[1])
			val2, err2 := parseVal(fields[2])
			if err1 != nil || err2 != nil {
				fmt.Println(Red + "Error: Invalid numbers provided." + Reset)
				continue
			}

			if command == "pct" {
				result := (val1 * val2) / 100.0
				lastResult = result
				fmt.Printf(Green+"Result: %g (%g%% of %g)\n"+Reset, result, val2, val1)
			} else {
				tipAmount := (val1 * val2) / 100.0
				total := val1 + tipAmount
				lastResult = total
				fmt.Printf(Green+"Tip: $%g | Total: $%g (Tip: %g%%)\n"+Reset, tipAmount, total, val2)
			}
			continue
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
			fmt.Println(Red + "Error: Unknown command or format. Try 'ncr 5 2' or '5 + 3'" + Reset)
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