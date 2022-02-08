package main

import (
	"fmt"
	"math"
	"time"
)

func main() {


	////////TIME
	tnw := time.Now()
	fmt.Println("time is now ", tnw)

	/////////MATH, CONVERT DECIMAL 2 DIGIT, DECLARE ALL IN A LINE
	f1, f2, f3 := 12.24, 13.28, 44.23
	floatSum := f1 + f2 + f3
	fmt.Println("float sum is : ", floatSum)
	floatSum = math.Round(floatSum*10) / 10
	fmt.Println(floatSum)

	//floating with 2 decimals after comma
	sirhelRadius := 12.3
	omfang := sirhelRadius * 2 * math.Pi
	fmt.Println("omfang er :", omfang)
	fmt.Println("%.2f", omfang) //not worhing .. why?

