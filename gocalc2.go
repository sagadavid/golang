package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	verdi1 := girdi("verdi 1")
	verdi2 := girdi("verdi 2")

	var result float64

	switch calcul := operandus(); calcul {
	case "+":
		result = addit(verdi1, verdi2)
	case "-":
		result = subt(verdi1, verdi2)
	case "*":
		result = mult(verdi1, verdi2)
	case "/":
		result = divd(verdi1, verdi2)
	default:
		panic("invalid operand")
	}
	result = math.Round(result*100) / 100
	fmt.Printf("result is %v\n\n", result)

}

func girdi(scrInput string) float64 {
	fmt.Printf("%v input a value : ", scrInput)
	input, _ := reader.ReadString('\n')
	parsdInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		tilbakemelding := fmt.Sprintf("%v not a number", scrInput)
		fmt.Println(tilbakemelding)
	}
	return parsdInput
}

func operandus() string {
	fmt.Print("choose operation: + - * / : ")
	oper, _ := reader.ReadString('\n')
	return strings.TrimSpace(oper)
}

func addit(verdi1, verdi2 float64) float64 {
	return verdi1 + verdi2
}

func subt(verdi1, verdi2 float64) float64 {
	return verdi1 - verdi2
}

func mult(verdi1, verdi2 float64) float64 {
	return verdi1 * verdi2
}

func divd(verdi1, verdi2 float64) float64 {
	return verdi1 / verdi2
}
