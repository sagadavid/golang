package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	////////CONSOLE CALCULATOR
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("number1 please : ")
	nInput1, _ := reader.ReadString('\n')
	aNumber1, err := strconv.ParseFloat(strings.TrimSpace(nInput1), 64)
	if err != nil {
		fmt.Println(err)
		panic("input must be a number")
	} else {
		fmt.Println("number1 is ", aNumber1)
	}

	fmt.Print("number2 please : ")
	nInput2, _ := reader.ReadString('\n')
	aNumber2, err := strconv.ParseFloat(strings.TrimSpace(nInput2), 64)
	if err != nil {
		fmt.Println(err)
		panic("input must be a number")
	} else {
		fmt.Println("number2 is ", aNumber2)
	}

	sum := aNumber1 + aNumber2
	sum = math.Round(sum*10) / 10
	//fmt.Println("and sum is", sum)

	fmt.Print("do you want to sum them: press 1 for yes or 0 for no \n")
	answer, _ := reader.ReadString('\n')
	nanswer, err := strconv.ParseFloat(strings.TrimSpace(answer), 64)
	if nanswer == 1 {
		fmt.Println("and sum is", sum)
	} else if nanswer == 0 {
		fmt.Println("oh fine")
	} else {
		fmt.Println("u dnt choose y or n, BYE BYE!")
	}

	///////////////GET INPUT TO CONSOLE
	leser := bufio.NewReader(os.Stdin)
	fmt.Print("input a text please : ")
	input, _ := leser.ReadString('\n')
	fmt.Println("du har shrevet", input)

	/////////////CONVERT CONSOLE INPUT WITH CONVERT PAChAGE
	fmt.Print("input a number please : ")
	nimput, _ := leser.ReadString('\n')
	aNumb, err := strconv.ParseFloat(strings.TrimSpace(nimput), 64)
	if err != nil {
		fmt.Println(err, "not a number or with gaps")
	} else {
		fmt.Println("you scored", aNumb)
	}
}
