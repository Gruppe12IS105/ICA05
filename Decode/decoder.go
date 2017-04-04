package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func ExDecodeMine() {
	const jsonStream = `
{
     "coord":{"lon":-122.08,"lat":37.39},
     "main":{"temp":277.14,"pressure":1025,"humidity":86,"temp_min":275.15,
     "temp_max":279.15}
`
	type Coordinates struct {
		Lon float64
		Lat float64
	}
	type Measurements struct {
		Temp     float64
		Pressure float64
		Humidity float64
		TempMin  float64
		TempMax  float64
	}
	type Weather struct {
		Coord Coordinates
		Main  Measurements
	}
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var w Weather
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
		fmt.Printf("Coordinates are: longitude %.2f and latitude %.2f\n",
			w.Coord.Lon, w.Coord.Lat)
		fmt.Printf("Temperature: %.f\n", w.Main.Temp)
	}
}

func main() {
	var url = "http://samples.openweathermap.org/data/2.5/weather?zip=94040,us&appid=b1b15e88fa797225412429c1c50c122a1"
}
