package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		if i%3 == 0 {
			fmt.Println("fizz", i)
		} else if i%5 == 0 {
			fmt.Println("buzz", i)
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzBuzz", i)
		}
	}
}
