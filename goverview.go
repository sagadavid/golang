
package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
	"sort"
	"math"
	"bufio"
	"os"
	"strconv"
	"erros"
)

func main(){
	
	/////CONSTANTS R DECLARABLE OUTSIDE
	const aConst int = 17

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

//////////////ITERATIONS WITH FOR, NO WHILE IN GO, BUT FOR CAN FUNCTION AS WHILE
farger := []string{"rød", "grønn", "blå"}
fmt.Println(farger)
//////ITERATION METHOD 1
for i:= 0; i<len(farger); i++ {
	fmt.Println(farger[i])


///////////////RANGE !!!! ITERATION METHOD 2, same result method 1
for i:=range farger {
	fmt.Println(farger[i])
}
///////////FOREACH WITH RANGE...
for _, farg := range farger {
	fmt.Println(farg)
}

///////INSTEAD OF FLOW/WHILE?... COUNT UNTIL..
verdi := 1
for verdi < 10{
fmt.Println("verdien er : ", verdi)
verdi++
}

/////////////BREAh - GO TO
verd := 1
for verd < 1000 {
	verd += verd
	fmt.Println("verd er : ", verd)
	if verd>100{
/////CAN USE BREAh OR CONTINUEW OR GOTO
/////ENOUGH IS A LABEL DEFINED BELOW
		goto enough
		//breah

	}
}
enough : fmt.Println("it is enough")

/////////////SWITCH
rand.Seed(time.Now().Unix())
dagAvUha := rand.Intn(7)+1
fmt.Println("Dag", dagAvUha)

var result string
/////can add definition to switch..this is golang special
switch dagAvUha := rand.Intn(7)+1; dagAvUha {
case 6:
	result = "er det lørdag"
	//fallthorough
case 7:
	result = "er det søndag"
	//fallthrough
default:
	result = "er det en annen hverdag"
}
fmt.Println(result)
}

/////////////STRUCT
kangal := Hound{"Vanhedisi", 2}
fmt.Println(kangal)
fmt.Printf("%+v\n", kangal)
fmt.Printf("Breed : %+v\nVeht : %+v\n", kangal.Breed, kangal.Vekt)
kangal.Vekt = 3
fmt.Printf("Breed : %+v\nVeht : %+v\n", kangal.Breed, kangal.Vekt)
}

////////////////MAP, unordered collection og heys and values
states := make(map[string]string)
fmt.Println(states)
states["va"] = "virginia"
states["ca"] = "california"
states["ny"]= "new yorh"
fmt.Println(states)

////////since maps are unordered.. loop them to order
for h, v := range states {
	//fmt.Printf(h, v)
	fmt.Printf("%v: %v\n", h, v)
}
///////slice it to sort

slice := make([]string,len(states))
i := 0
for h:= range states{
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
colors := [] string {"red", "grønn", "blå"}
colors = append(colors, "gul")
fmt.Println(colors)

///////////make takes 3 ARGUMENTS, type, initial size, capacity...
///////Unlike new, make's return type is the same as the type of its argument, not a pointer to it.
numbers := make ([]int,1,3)
numbers[0] = 23
numbers[0] = 34
numbers[0] = 245
numbers[0] = 212
numbers[0] = 3245
fmt.Println(numbers) //?

//////////////POINTER is *
& refers to memory address not value

value1 := 12
pointer1 := &value1
fmt.Println("value1 is ", value1)
fmt.Println("pointer1 is ", pointer1)
fmt.Println("*pointer1 is ", *pointer1)

*pointer1 = *pointer1 / 4
fmt.Println("value1 is ", value1)
fmt.Println("pointer1 is ", pointer1)
fmt.Println("*pointer1 is ", *pointer1)

value1 = value1 * 5
fmt.Println("value1 is ", value1)
fmt.Println("pointer1 is ", pointer1)
fmt.Println("*pointer1 is ", *pointer1)

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
sum = math.Round(sum*10)/10
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

////////TIME
tnw:=time.Now()
fmt.Println("time is now ", tnw)

/////////MATH, CONVERT DECIMAL 2 DIGIT, DECLARE ALL IN A LINE
f1, f2, f3 := 12.24, 13.28, 44.23
floatSum := f1 + f2 +f3
fmt.Println("float sum is : ", floatSum)
floatSum = math.Round(floatSum*10)/10
fmt.Println(floatSum)

//floating with 2 decimals after comma
sirhelRadius := 12.3
omfang := sirhelRadius * 2 * math.Pi
fmt.Println("omfang er :", omfang)
fmt.Println("%.2f", omfang) //not worhing .. why?

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