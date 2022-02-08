package main

import "fmt"

func main() {

	multiSum, counted := multipleCounts(5,86,12,65,3,101)
	fmt.Println("multiSum is : ", multiSum)
	fmt.Println(counted, "numbers ar counted")
}

/////// RETURN 2 VALUES FROM A FUNCTION
func multipleCounts (values ...int) (int, int) {
	total := 0
	for _, value := range values{
		total += value
	}
	return total, len(values)
}

}
