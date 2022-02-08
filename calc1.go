package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

////SEQ1-A READER TO READ CONSOLE TEXT
var reader = bufio.NewReader(os.Stdin)

func main() {
	////SEQ2-2
	value1 := getInput("Value 1")
	value2 := getInput("Value 2")

	var result float64
	/////SEQ4
	switch operation := getOperation(); operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	case "*":
		result = multiplyValues(value1, value2)
	case "/":
		result = divideValues(value1, value2)
	default:
		panic("Invalid operation")
	}
	fmt.Printf("The result is %v\n\n", result)

}

///SEQ2-1-TAKEIN INPUT
func getInput(prompt string) float64 {
	//SCREEN PROMPT IS REQUESTED AS INPUTSOURCE
	fmt.Printf("%v: ", prompt)
	//NOW PROMPT/INPUT IS READ
	input, _ := reader.ReadString('\n')
	//INPUT IS TRIMMED FOR SPACES AND PARTSED AFTERWARDS
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	//PARSING IS CHECKED FOR ERROR
	if err != nil {
		//CREATED ERROR MESSAGE IS PROPOSED
		message := fmt.Sprintf("%v must be a number", prompt)
		panic(message)
	}
	//PARSED RESULT IS RETURNED
	return value
}

/////SEQ3-PROMPT IS RENDERED AS OPERAND
func getOperation() string {
	fmt.Print("Select an operation (+ - * /): ")
	//PROMPT IS TAKED AS OPERAND
	op, _ := reader.ReadString('\n')
	//PROMPT INPUIT IS TRIMMED
	return strings.TrimSpace(op)
}

/////SEQ5
func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
