package main

////READ A FILE FROM WEB
import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	/////JSON DECODING SEQUENCE 2
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	////GET DATA OVER HTTP
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type %T\n", resp)

	////FETCHING JSON as BYTES
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	////CONVERT BYTES TO STRING
	content := string(bytes)
	///THIS CONTENT BELOW IS REFINED BELOW... NO NEED IT ANYMORE
	// fmt.Print(content)

	/////JSON DECODING SEQUENCE 5
	///CALL THE FUNCTION
	tours := toursFromJson(content)
	///LOOP THROUG THE TOUR OBJECT
	for _, tour := range tours {
		fmt.Println(tour.Name)
	}
}

/////JSON DECODING SEQUENCE 3
func toursFromJson(content string) []Tour {
	/////CREATE A SLICE
	tours := make([]Tour, 0, 20)
	////DECODE JSON
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}
	/////JSON DECODING SEQUENCE 4
	//////TRANSFORM JSON TEKST TO SLICE FORMATTED TEXT
	var tour Tour
	/////MORE READS NEXT AVAILABLE OBJECT AS BOOLEAN
	for decoder.More() {
		////PASS IN MEMORY ADRESS OF TOUR OBJECT
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		//////ADD ACHIEVED OBJECTS TO TOUR
		tours = append(tours, tour)
	}
	return tours
}

/////JSON DECODING SEQUENCE 1
type Tour struct {
	Name, Price string
}
