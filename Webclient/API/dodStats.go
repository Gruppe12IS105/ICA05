package API

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type DodStats struct {
	Dataset struct {
		Dimension struct {
			Region struct {
				Label    string `json:"label"`
				Category struct {
					Index struct {
						Num0 int `json:"0"`
					} `json:"index"`
					Label struct {
						Num0 string `json:"0"`
					} `json:"label"`
				} `json:"category"`
			} `json:"Region"`
			ContentsCode struct {
				Label    string `json:"label"`
				Category struct {
					Index struct {
						Folketallet1      int `json:"Folketallet1"`
						Fodte2            int `json:"Fodte2"`
						Dode3             int `json:"Dode3"`
						Fodselsoverskudd4 int `json:"Fodselsoverskudd4"`
						Innvandring5      int `json:"Innvandring5"`
						Utvandring6       int `json:"Utvandring6"`
						Tilflytting7      int `json:"Tilflytting7"`
						Fraflytting8      int `json:"Fraflytting8"`
						Nettoinnflytting9 int `json:"Nettoinnflytting9"`
						Folketilvekst10   int `json:"Folketilvekst10"`
						Folketallet11     int `json:"Folketallet11"`
					} `json:"index"`
					Label struct {
						Folketallet1      string `json:"Folketallet1"`
						Fodte2            string `json:"Fodte2"`
						Dode3             string `json:"Dode3"`
						Fodselsoverskudd4 string `json:"Fodselsoverskudd4"`
						Innvandring5      string `json:"Innvandring5"`
						Utvandring6       string `json:"Utvandring6"`
						Tilflytting7      string `json:"Tilflytting7"`
						Fraflytting8      string `json:"Fraflytting8"`
						Nettoinnflytting9 string `json:"Nettoinnflytting9"`
						Folketilvekst10   string `json:"Folketilvekst10"`
						Folketallet11     string `json:"Folketallet11"`
					} `json:"label"`
					Unit struct {
						Folketallet1 struct {
							Base string `json:"base"`
						} `json:"Folketallet1"`
						Fodte2 struct {
							Base string `json:"base"`
						} `json:"Fodte2"`
						Dode3 struct {
							Base string `json:"base"`
						} `json:"Dode3"`
						Fodselsoverskudd4 struct {
							Base string `json:"base"`
						} `json:"Fodselsoverskudd4"`
						Innvandring5 struct {
							Base string `json:"base"`
						} `json:"Innvandring5"`
						Utvandring6 struct {
							Base string `json:"base"`
						} `json:"Utvandring6"`
						Tilflytting7 struct {
							Base string `json:"base"`
						} `json:"Tilflytting7"`
						Fraflytting8 struct {
							Base string `json:"base"`
						} `json:"Fraflytting8"`
						Nettoinnflytting9 struct {
							Base string `json:"base"`
						} `json:"Nettoinnflytting9"`
						Folketilvekst10 struct {
							Base string `json:"base"`
						} `json:"Folketilvekst10"`
						Folketallet11 struct {
							Base string `json:"base"`
						} `json:"Folketallet11"`
					} `json:"unit"`
				} `json:"category"`
			} `json:"ContentsCode"`
			Tid struct {
				Label    string `json:"label"`
				Category struct {
					Index struct {
						Two016K4 int `json:"2016K4"`
					} `json:"index"`
					Label struct {
						Two016K4 string `json:"2016K4"`
					} `json:"label"`
				} `json:"category"`
			} `json:"Tid"`
			ID   []string `json:"id"`
			Size []int    `json:"size"`
			Role struct {
				Metric []string `json:"metric"`
				Time   []string `json:"time"`
			} `json:"role"`
		} `json:"dimension"`
		Label   string    `json:"label"`
		Source  string    `json:"source"`
		Updated time.Time `json:"updated"`
		Value   []int     `json:"value"`
	} `json:"dataset"`
}

func GetDodStats(urlDodStats string) {
	respons, err := http.Get(urlDodStats)
	if err != nil {
		panic(err.Error())
	}
	dodStatsJSON, err := ioutil.ReadAll(respons.Body)
	if err != nil {
		panic(err.Error())
	}
	convertDodStats(dodStatsJSON)
}
func convertDodStats(dodStatsJSON []byte) {
	var d DodStats
	err := json.Unmarshal(dodStatsJSON, &d)
	if err != nil {
		panic(err.Error())
	}
	time.Sleep(3 * time.Second)
	fmt.Println("----Dødsstatistikk i Norge----")
	fmt.Printf("Label: %+v\nSource: %+v\nUpdatet: %+v\nValue: %+v\n", d.Dataset.Label, d.Dataset.Source, d.Dataset.Updated, d.Dataset.Value)
	fmt.Println("---------------------------")
}
