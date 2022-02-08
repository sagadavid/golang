package main

import (
	"fmt"
)

////////////STRUCT INSTANCE with PROPERTIES
type Hound struct {
	Breed   string
	Vekt    int
	HøresUt string
}

func main() {

	/////////////METHODS BELONG TO TYPES IN GO NOT TO CLASSES
	kangal := Hound{"Kangal", 5, "vaf"}
	fmt.Println(kangal)
	fmt.Printf("%+v\n", kangal)
	fmt.Printf("Breed : %+v\nVeht : %+v\n", kangal.Breed, kangal.Vekt)
	kangal.Vekt = 4
	fmt.Printf("Breed : %+v\nVeht : %+v\n", kangal.Breed, kangal.Vekt)

	//////her, d.sound er defineert som speak og derfor vi får bare et sound. siden er det et kopi. om du vil kan du referere som pointer...
	kangal.Snakker()
	kangal.HøresUt = "harf"
	kangal.Snakker()
	kangal.bark3Times()
	kangal.bark3Times()
}

///////METHOD TO CLASS...  first, pass THE CLASS/STRUCT as a receiver(h Hound)..
//and when it is passed a kopi is made.. its not a reference!!!
//and you can have a method to it.
func (h Hound) Snakker() {
	fmt.Println(h.HøresUt)
}

func (h Hound) bark3Times() {
	h.HøresUt = fmt.Sprintf("%v %v %v", h.HøresUt, h.HøresUt, h.HøresUt)
	fmt.Println(h.HøresUt)
}



/////////////STRUCT
kangal := Hound{"Vanhedisi", 2}
fmt.Println(kangal)
fmt.Printf("%+v\n", kangal)
fmt.Printf("Breed : %+v\nVeht : %+v\n", kangal.Breed, kangal.Vekt)
kangal.Vekt = 3
fmt.Printf("Breed : %+v\nVeht : %+v\n", kangal.Breed, kangal.Vekt)
}
