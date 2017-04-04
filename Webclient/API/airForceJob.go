package API

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type AirForceJobs []struct {
	ID               string   `json:"id"`
	PositionTitle    string   `json:"position_title"`
	OrganizationName string   `json:"organization_name"`
	RateIntervalCode string   `json:"rate_interval_code"`
	Minimum          int      `json:"minimum"`
	Maximum          int      `json:"maximum"`
	StartDate        string   `json:"start_date"`
	EndDate          string   `json:"end_date"`
	Locations        []string `json:"locations"`
	URL              string   `json:"url"`
}

func USAAirForceJobs(urlAirForce string) {
	respons, err := http.Get(urlAirForce)
	if err != nil {
		panic(err.Error())
	}
	usaAirForceJobsJSON, err := ioutil.ReadAll(respons.Body)
	if err != nil {
		panic(err.Error())
	}
	convertAirForce(usaAirForceJobsJSON)
}
func convertAirForce(usaAirForceJobsJSON []byte) {
	s := string(usaAirForceJobsJSON[:])
	dec := json.NewDecoder(strings.NewReader(s))
	var a AirForceJobs
	if err := dec.Decode(&a); err == io.EOF {
	} else if err != nil {
		log.Fatal(err)
	}
	time.Sleep(4 * time.Second)
	var loper = 1
	for _, airForceJobs := range a {
		fmt.Println("AirForce-API nr", loper)
		fmt.Printf("ID: %+v\nPosition title: %+v\nOrganization name: %+v\nRate Interval Code: %+v\nMinimum: %+v\nMaximum: %+v\nStart date: %+v\nEnd date: %+v\nLocations: %+v\nURL: %+v\n", airForceJobs.ID, airForceJobs.PositionTitle, airForceJobs.OrganizationName, airForceJobs.RateIntervalCode, airForceJobs.Minimum, airForceJobs.Maximum, airForceJobs.StartDate, airForceJobs.EndDate, airForceJobs.Locations, airForceJobs.URL)
		fmt.Println("---------------------------")
		loper++
	}
}
