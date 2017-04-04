package konvertere

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	CITI_BIKE_JSON = "https://www.citibikenyc.com/stations/json"
	OPENWEATHER    = "http://samples.openweathermap.org/data/2.5/weather?zip=94040,us&appid=b1b15e88fa797225412429c1c50c122a1"
)

type Blanding struct {
	Lon float64
	Lat float64
}

/*
type Blanding struct {
	BlandingListe []Coord
}
*/

//type API struct {
//}

/*
   Denne koden er bare kopiert fra https://github.com/josephmisiti/go-citibike
   Utskriften blir kun:
   ------
   &{[]}
*/
func GetWeather() *Blanding {
	body, err := MakeRequest(OPENWEATHER)
	if err != nil {
		fmt.Println(err)
	}
	ferdigParset, err := ParseWeather(body)
	return ferdigParset
}

/*
	Disse to metodene under blir kalt av GetWeather()
*/
func MakeRequest(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf("%q\n", body)
	return []byte(body), err
}

func ParseWeather(body []byte) (*Blanding, error) {
	var blanding = new(Blanding)
	err := json.Unmarshal(body, blanding)
	if err == nil {
		fmt.Println("Det skjedde en feil:", err)
	}
	//fmt.Printf("%q", blanding)
	return blanding, err
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
