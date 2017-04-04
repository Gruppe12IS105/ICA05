package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./Decode"
	//"os"
	//Påminnelse til senere: Importér pakken med API
)

// NOTE: denne må startes som en goroutine for hver URL/datakilde
// som dere skal hente inn-data fra.

func doGet(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%q", contents)
		fmt.Println("   ", response.StatusCode)
		hdr := response.Header
		for key, value := range hdr {
			fmt.Println("   ", key, ":", value)
		}
	}
	return contents
}

func main() {
	//var url = "https://pozzad-email-validator.p.mashape.com/emailvalidator/validateEmail/{email}"
	//nøkkel: J0GY3KAb6pmshcL8IJrNubPtNrKZp1rEZLojsn4lkQoBngfOGi
	var url = "http://samples.openweathermap.org/data/2.5/weather?zip=94040,us&appid=b1b15e88fa797225412429c1c50c122a1"
	go Decode.ExDecodeMine(url)
	//doGet(url)
}
