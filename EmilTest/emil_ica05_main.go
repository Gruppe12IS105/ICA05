package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	VERSION        = 1
	CITI_BIKE_JSON = "https://www.citibikenyc.com/stations/json"
	OPENWEATHER    = "http://samples.openweathermap.org/data/2.5/weather?zip=94040,us&appid=b1b15e88fa797225412429c1c50c122a1"
)

type Coord struct {
	Coordinates []string `json:"coords"`
	Weather     []string `json:"weather"`
	Humidity    int      `json:"humidity"`
}
type Blanding struct {
	BlandingListe []Coord `json:"blandingListe"`
}

type API struct {
}

/*
   Denne koden er bare kopiert fra https://github.com/josephmisiti/go-citibike
   Utskriften blir kun:
   ------
   &{[]}
*/
func main() {
	fmt.Println("------")
	API := new(API)
	stations, err := API.getStations()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stations)
}

func (r API) getStations() (*Blanding, error) {

	body, err := makeRequest(OPENWEATHER)
	if err != nil {
		return nil, err
	}
	s, err := parseStations(body)
	return s, err
}

func makeRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return []byte(body), err
}

func parseStations(body []byte) (*Blanding, error) {
	var s = new(Blanding)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

//////////////////////////////////////////////////////
/*
  apiLink, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tekst)
}
*/
