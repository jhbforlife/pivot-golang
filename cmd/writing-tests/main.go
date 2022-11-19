package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jhbforlife/pivot-golang/calculator"
)

func main() {
	fmt.Println("Calculator + | - | * | /")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter calculation (ex: 2 + 2): ")
		str, err := reader.ReadString('\n')
		if checkPrintError(err) {
			return
		}
		handleInput(str)
	}
}

func handleInput(str string) {
	var sepStr = []string{}
	str = strings.Replace(str, "\n", "", -1)
	operator, err := checkContains(str)
	if checkPrintError(err) {
		return
	}
	sepStr = strings.Split(str, operator)
	if len(sepStr) > 3 {
		fmt.Println(errors.New("calculator: too many arguments. please try again"))
	}
	for i := range sepStr {
		sepStr[i] = strings.Replace(sepStr[i], " ", "", -1)
		if i > 1 {
			fmt.Println(errors.New("calculator: only one calculation allowed at a time. please try again"))
			return
		}
	}
	a, err := strconv.Atoi(sepStr[0])
	if checkPrintError(err) {
		return
	}
	b, err := strconv.Atoi(sepStr[1])
	if checkPrintError(err) {
		return
	}
	var result int
	switch operator {
	case "+":
		result = calculator.Add(a, b)
	case "-":
		result = calculator.Subtract(a, b)
	case "*":
		result = calculator.Multiply(a, b)
	case "/":
		result, err = calculator.Divide(a, b)
		if checkPrintError(err) {
			return
		}
	default:
		fmt.Println(errors.New("calculator: error with calculation. please try again"))
	}
	fmt.Printf("Result: %d\n", result)
}

func checkContains(str string) (string, error) {
	operators := []string{"+", "-", "*", "/"}
	containsOperator := []string{}
	for _, v := range operators {
		if strings.Contains(str, v) {
			containsOperator = append(containsOperator, v)
		}
	}
	switch len(containsOperator) {
	case 0:
		return "", errors.New("calculator: no operator found. please try again")
	case 1:
		return containsOperator[0], nil
	default:
		return "", errors.New("calculator: more than one operator found. please try again")
	}
}

func checkPrintError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
