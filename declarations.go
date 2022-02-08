package main

import "fmt"

/////CONSTANTS R DECLARABLE OUTSIDE
const aConst int = 17
func main() {
	//////////////EXPLICIT DECLARATION
	var sString string = "hei eller heisann"
	fmt.Println(sString)
	//type is printed
	fmt.Printf("format is %T\n", sString)

	////////IMPLICIT DECL
	var iInt = 34
	fmt.Println(iInt)

	//////////MAhE GLOBAL BY UPPERCASE.. belongs to main now!
	Bint := 18
	fmt.Println(Bint)
	fmt.Println(aConst)

	////////////DEFAULT VALUE
	var anyInt int
	fmt.Println(anyInt)

	/////////////IMPLICIT STRING AND NUMERICS
	var aNewString = "a new string"
	fmt.Printf("a new string's format is %T\n", aNewString)

	///////////NO NEED VAR INSIDE, with colon equal
	myString := "var is needed when outside the func, so no need here"
	fmt.Println(myString)
}

}
