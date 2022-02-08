/////////////////////PREVIOUS EXERXICES////////

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
fmt.Println("hei verden på go !")
	var x int
	x = 5
	var y int = 3
	z:=6
	var sum = x+y+z
	fmt.Println(sum)

	if sum > 2 {fmt.Println("sum is bigger than 2")}

///////////////ARRAYS, mention how much element and type
var array [4]int
array[0] = 34
array[1] = 23
array[2] = 45
array[3] = 56
fmt.Println(array)

	bArr := [5]int{3,6,7,88,92}
	fmt.Println(bArr)

	cArr:=[]int{4,5}
cArr = append(cArr, 34,94,23,1)
fmt.Println(cArr)

mappe := mahe(map[string]int)
mappe["board"]= 3
mappe["colon"]=6
mappe["sisu"]=5
fmt.Println(mappe)
fmt.Println(mappe["sisu"])
delete(mappe,"board")
fmt.Println(mappe)

for i :=0; i <6; i++{fmt.Println(i)}
i :=0
for i<3{fmt.Println(i)
i++}
fmt.Println("bArr")
for index, value :=range bArr{
	fmt.Println("index", index, "sitt value er", value)
}
fmt.Println("mappe")
for hey, value :=range mappe{
	fmt.Println("hey is ", hey, "value is", value)
}

rez:=addal(3,8)
fmt.Println("addit is ")
fmt.Println(rez)

result, err := hvdrt(-1)
if err!= nil {
	fmt.Println(err)
} else {
	fmt.Print("hvdrtRot is ")
fmt.Println(result)
}

nihy := enPerson{name:"olav", age : 23}
fmt.Println(nihy)
fmt.Println("age is")
fmt.Println(nihy.age)

r := 8
fmt.Println("memory reference of int r is")
fmt.Println(&r)
fmt.Println("dereference by using *")
}}

func hvdrt(x float64) (float64, error){if x < 0 {
	return 0, errors.New("not defined for negative numbers")
}
return math.Sqrt(x), nil
}

func addal(m int, n int)int{
	return m+n
}

type enPerson struct{
	name string
	age int}}
