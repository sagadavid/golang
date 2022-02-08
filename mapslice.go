package main

import (
	"fmt"
	"sort"
)

func main() {
	////////////////MAP, unordered collection og keys and values
	states := make(map[string]string)
	fmt.Println(states)
	states["va"] = "virginia"
	states["ca"] = "california"
	states["ny"] = "new yorh"
	fmt.Println(states)

	////////since maps are unordered.. loop them to order
	for h, v := range states {
		//fmt.Printf(h, v)
		fmt.Printf("%v: %v\n", h, v)
	}
	///////slice it to sort

	slice := make([]string, len(states))
	i := 0
	for h := range states {
		slice[i] = h
		i++
	}
	//fmt.Println(slice)
	sort.Strings(slice)
	fmt.Println(slice)

	///////////sort alphabetical
	for i := range slice {
		fmt.Println(states[slice[i]])
	}

	//////////SLICES, resizable arrays, all elements same type, just dont mention amount of elements
	colors := []string{"red", "grønn", "blå"}
	colors = append(colors, "gul")
	fmt.Println(colors)

	///////////make takes 3 ARGUMENTS, type, initial size, capacity...
	///////Unlike new, make's return type is the same as the type of its argument, not a pointer to it.
	numbers := make([]int, 1, 3)
	numbers[0] = 23
	numbers[0] = 34
	numbers[0] = 245
	numbers[0] = 212
	numbers[0] = 3245
	fmt.Println(numbers) //?

}
