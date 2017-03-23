package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	url := "http://samples.openweathermap.org/data/2.5/weather?zip=94040,us&appid=b1b15e88fa797225412429c1c50c122a1"
	//  url := os.Args[1]
	doGet(url)
	const jsonStream = `
{"Name": "Værmelding", "Text": "Været i dag."}
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}

func doGet(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
			fmt.Printf("%q", contents)
			fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
			fmt.Println(" ", response.StatusCode)
			hdr := response.Header
			for key, value := range hdr {
				fmt.Println(" ", key, ":", value)
			}
		}
	}
}
func get(hdr) {
hdr := response.Header
}
func exDecodeMine()  {
	var jsonStream =
"coord":{"lon":-122.08,"lat":37.39},
"main":{"temp":277.14,"pressure":1025,"humidity":86,"temp_min":275.15,
"temp_max":279.15}
}

type Coordinates struct {
Lon float64
Lat float64
}
type Measurements struct {
Temp float64
Pressure float64
Humidity float64
Temp_min float64
Temp_max float64
}
type Weather struct {
Coord Coordinates
Main Measurements
}
