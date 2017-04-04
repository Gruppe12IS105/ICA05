package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Coordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
type Measurement struct {
	Temp     float64
	Pressure float64
	Humidity float64
	Temp_min float64
	Temp_max float64
}
type Weather struct {
	Coord Coordinates
	Main  Measurement
}

func main() {
	getWeather()
}

func getWeather() {
	respons, err := http.Get("http://samples.openweathermap.org/data/2.5/weather?zip=94040,us&appid=b1b15e88fa797225412429c1c50c122a1")
	if err != nil {
		panic(err.Error())
	}
	jsonByteArray, err := ioutil.ReadAll(respons.Body)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf("%q", body)
	konvertere(jsonByteArray)
}

func konvertere(jsonByteArray []byte) {
	var w Weather
	err := json.Unmarshal(jsonByteArray, &w) //vet ikke nøyaktig hva '&' gjør, men får feilmelding uten.
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("-------------------\n")
	fmt.Printf("Lengdegradene er: %+v.\nBreddegradene er: %+v\n", w.Coord.Lon, w.Coord.Lat) //%+v er for å få verdiene, \n er bare ny linje
}
