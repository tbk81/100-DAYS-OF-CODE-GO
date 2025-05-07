package main

import "fmt"

var li = []string{"+", "-", "/", "*"}

func operandDisplay() {
	for _, v := range li {
		fmt.Println(v)
	}
}

func add(f, s float64) float64 {
	return f + s
}

func substract(f, s float64) float64 {
	return f - s
}

func divide(f, s float64) float64 {
	return f / s
}

func multiply(f, s float64) float64 {
	return f * s
}

func main() {
	var firstNum, secondNum, solution float64
	var operand, usrChoice string

	fmt.Println("Welcome to the 4 function calculator")
	fmt.Print("Type your first number: ")
	fmt.Scanln(&firstNum)
	operandDisplay()
	fmt.Print("Pick your operation: ")
	fmt.Scanln(&operand)
	fmt.Print("Type your second number: ")
	fmt.Scanln(&secondNum)

	switch {
	case operand == "+":
		solution = add(firstNum, secondNum)
	case operand == "-":
		solution = substract(firstNum, secondNum)
	case operand == "/":
		solution = divide(firstNum, secondNum)
	case operand == "*":
		solution = multiply(firstNum, secondNum)
	}

	fmt.Println(firstNum, operand, secondNum, "=", solution)
	fmt.Println("Type 'y' to continue calculating or type 'n' to start a new calculation: ")
	fmt.Scanln(&usrChoice)

}
