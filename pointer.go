package main

import "fmt"

func main() {

	//////////////POINTER is *
	//* refers to a location in memory there value stored
	//& refers to memory address not value
	value := 12
	verdi := &value
	fmt.Println("value is ", value)
	fmt.Println("&value is ", &value)
	fmt.Println("verdi is ", verdi)
	fmt.Println("*verdi is ", *verdi)
	fmt.Println("-----------")
	*verdi = *verdi / 4
	fmt.Println("value is ", value)
	fmt.Println("&value is ", &value)
	fmt.Println("verdi is ", verdi)
	fmt.Println("*verdi is ", *verdi)
	fmt.Println("-----------")
	value = value * 5
	fmt.Println("value is ", value)
	fmt.Println("&value is ", &value)
	fmt.Println("verdi is ", verdi)
	fmt.Println("*verdi is ", *verdi)

}
