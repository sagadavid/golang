package main

import "fmt"

//func zero(x int) {
//	x = 0
//}

//func main() {
//	x := 5
//	zero(x)
//	fmt.Println(x)
//	//x skulle nulles/0, men for verdien limmes inni func
//	//når det er brukt som en parameter!
//	//func zero () kan ikke nulle den enda.derfor;
//	//x er 5 i terminal output
//}
//LØSNING ER POINTER BRUK
func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}
func zero(x *int) {
	*x = 0
}
