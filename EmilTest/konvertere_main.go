package main

import "fmt"
import "./konvertere"

func main() {
	fmt.Println("---------")
	openWeather := konvertere.GetWeather()
	fmt.Println(openWeather)
}
