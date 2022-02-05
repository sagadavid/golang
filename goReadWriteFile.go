package main

///////CREATING AND READING TEXT FILES
import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	innhold := "hello .. lag en tekst og les den.. eksampel"
	//CREATING A FILE
	fil, err := os.Create("./txt2sameFolder.txt")
	checkError(err)
	langt, err := io.WriteString(fil, innhold)
	checkError(err)
	fmt.Printf("lagt en fil i %v karakter\n", langt)
	//defer means; wait until all done and then execute
	defer fil.Close()
	//her trenger det adressen
	defer lesFil("./txt2sameFolder.txt")
}

func lesFil(filsNavn string) {
	//få data av fil via lesing
	data, err := ioutil.ReadFile(filsNavn)
	checkError(err)
	//print den som lest, bytt til string
	fmt.Println("leser av filen: ", string(data))

}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
