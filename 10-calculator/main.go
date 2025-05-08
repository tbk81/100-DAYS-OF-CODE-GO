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

func calc(n1, n2 float64, o string) float64 {
	var solution float64
	switch o {
	case "+":
		solution = add(n1, n2)
	case "-":
		solution = substract(n1, n2)
	case "/":
		solution = divide(n1, n2)
	case "*":
		solution = multiply(n1, n2)
	}
	return solution
}

func main() {
	var firstNum, secondNum float64
	var operand, usrChoice string

	fmt.Println("Welcome to the 4 function calculator")
Loop:
	for {
		fmt.Print("Type your first number: ")
		fmt.Scanln(&firstNum)
		operandDisplay()
		fmt.Print("Pick your operation: ")
		fmt.Scanln(&operand)
		fmt.Print("Type your second number: ")
		fmt.Scanln(&secondNum)
		answer := calc(firstNum, secondNum, operand)

		fmt.Println(firstNum, operand, secondNum, "=", answer)
		fmt.Printf("Type 'y' to continue calculating with %v or type 'n' to start a new calculation: ", answer)
		fmt.Scanln(&usrChoice)
		switch usrChoice {
		case "y":
			goto Loop
		case "n":
			goto Loop
		case "exit":
			break Loop
		}
	}

}
