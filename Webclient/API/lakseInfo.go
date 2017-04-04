package API

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Laks struct {
	Dataset struct {
		Dimension struct {
			VareGrupper2 struct {
				Label    string `json:"label"`
				Category struct {
					Index struct {
						Num01 int `json:"01"`
						Num02 int `json:"02"`
					} `json:"index"`
					Label struct {
						Num01 string `json:"01"`
						Num02 string `json:"02"`
					} `json:"label"`
				} `json:"category"`
			} `json:"VareGrupper2"`
			Tid struct {
				Label    string `json:"label"`
				Category struct {
					Index struct {
						Two017U11 int `json:"2017U11"`
						Two017U12 int `json:"2017U12"`
					} `json:"index"`
					Label struct {
						Two017U11 string `json:"2017U11"`
						Two017U12 string `json:"2017U12"`
					} `json:"label"`
				} `json:"category"`
			} `json:"Tid"`
			ContentsCode struct {
				Label    string `json:"label"`
				Category struct {
					Index struct {
						Vekt     int `json:"Vekt"`
						Kilopris int `json:"Kilopris"`
					} `json:"index"`
					Label struct {
						Vekt     string `json:"Vekt"`
						Kilopris string `json:"Kilopris"`
					} `json:"label"`
					Unit struct {
						Vekt struct {
							Base string `json:"base"`
						} `json:"Vekt"`
						Kilopris struct {
							Base string `json:"base"`
						} `json:"Kilopris"`
					} `json:"unit"`
				} `json:"category"`
			} `json:"ContentsCode"`
			ID   []string `json:"id"`
			Size []int    `json:"size"`
			Role struct {
				Time   []string `json:"time"`
				Metric []string `json:"metric"`
			} `json:"role"`
		} `json:"dimension"`
		Label   string    `json:"label"`
		Source  string    `json:"source"`
		Updated time.Time `json:"updated"`
		Value   []float64 `json:"value"`
	} `json:"dataset"`
}

func LakseInfo(urlLakseStats string) {
	respons, err := http.Get(urlLakseStats)
	if err != nil {
		panic(err.Error())
	}
	LakseinfoJSON, err := ioutil.ReadAll(respons.Body)
	if err != nil {
		panic(err.Error())
	}
	convertLakseInfoJSON(LakseinfoJSON)
}
func convertLakseInfoJSON(LakseinfoJSON []byte) {
	var l Laks
	err := json.Unmarshal(LakseinfoJSON, &l)
	if err != nil {
		panic(err.Error())
	}
	time.Sleep(2 * time.Second)
	fmt.Println("------LakseInfo------")
	fmt.Printf("Value: %+v\nDataset Label: %+v\nSource: %+v\nUpdated: %+v\nID: %+v\nSize: %+v\nContents code Label: %+v\nIndex - Kilopris: %+v, Vekt: %+v\nLabel - Kilopris: %+v, Vekt: %+v\nUnit - Kilopris %+v, Vekt: %+v\nRole - Metric: %+v, Time: %+v\nTid Label: %+v\nIndex - Two017U11: %+v, Two017U12: %+v\nLabel - Two017U11: %+v, Two017U12: %+v\nVaregrupper2 label: %+v\nVaregrupper2 - Index: Num01: %+v, Num02: %+v\nVaregrupper2 - Label: Num01: %+v, Num02: %+v\n", l.Dataset.Value, l.Dataset.Label, l.Dataset.Source, l.Dataset.Updated, l.Dataset.Dimension.ID, l.Dataset.Dimension.Size, l.Dataset.Dimension.ContentsCode.Label, l.Dataset.Dimension.ContentsCode.Category.Index.Kilopris, l.Dataset.Dimension.ContentsCode.Category.Index.Vekt, l.Dataset.Dimension.ContentsCode.Category.Label.Kilopris, l.Dataset.Dimension.ContentsCode.Category.Label.Vekt, l.Dataset.Dimension.ContentsCode.Category.Unit.Kilopris, l.Dataset.Dimension.ContentsCode.Category.Unit.Vekt, l.Dataset.Dimension.Role.Metric, l.Dataset.Dimension.Role.Time, l.Dataset.Dimension.Tid.Label, l.Dataset.Dimension.Tid.Category.Index.Two017U11, l.Dataset.Dimension.Tid.Category.Index.Two017U12, l.Dataset.Dimension.Tid.Category.Label.Two017U11, l.Dataset.Dimension.Tid.Category.Label.Two017U12, l.Dataset.Dimension.VareGrupper2.Label, l.Dataset.Dimension.VareGrupper2.Category.Index.Num01, l.Dataset.Dimension.VareGrupper2.Category.Index.Num02, l.Dataset.Dimension.VareGrupper2.Category.Label.Num01, l.Dataset.Dimension.VareGrupper2.Category.Label.Num02)
	fmt.Println("---------------------------")

}
