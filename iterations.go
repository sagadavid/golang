package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//////////////ITERATIONS WITH FOR, NO WHILE IN GO, BUT FOR CAN FUNCTION AS WHILE
	farger := []string{"rød", "grønn", "blå"}
	fmt.Println(farger)
	//////ITERATION METHOD 1
	for i := 0; i < len(farger); i++ {
		fmt.Println(farger[i])

		///////////////RANGE !!!! ITERATION METHOD 2, same result method 1
		for i := range farger {
			fmt.Println(farger[i])
		}
		///////////FOREACH WITH RANGE...
		for _, farg := range farger {
			fmt.Println(farg)
		}

		///////INSTEAD OF FLOW/WHILE?... COUNT UNTIL..
		verdi := 1
		for verdi < 10 {
			fmt.Println("verdien er : ", verdi)
			verdi++
		}

		/////////////BREAh - GO TO
		verd := 1
		for verd < 1000 {
			verd += verd
			fmt.Println("verd er : ", verd)
			if verd > 100 {
				/////CAN USE BREAh OR CONTINUEW OR GOTO
				/////ENOUGH IS A LABEL DEFINED BELOW
				goto enough
				//breah

			}
		}
	enough:
		fmt.Println("it is enough")

		/////////////SWITCH
		rand.Seed(time.Now().Unix())
		dagAvUha := rand.Intn(7) + 1
		fmt.Println("Dag", dagAvUha)

		var result string
		/////can add definition to switch..this is golang special
		switch dagAvUha := rand.Intn(7) + 1; dagAvUha {
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
}
