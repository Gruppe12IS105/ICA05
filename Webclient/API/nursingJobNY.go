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

type NursingJobs_NY []struct {
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

func USANursingJobs_NY(urlNursing_NY string) {
	respons, err := http.Get(urlNursing_NY)
	if err != nil {
		panic(err.Error())
	}
	usaNursingJobsJSON_NY, err := ioutil.ReadAll(respons.Body)
	if err != nil {
		panic(err.Error())
	}
	convertNursingJobs_NY(usaNursingJobsJSON_NY)
}
func convertNursingJobs_NY(usaNursingJobsJSON_NY []byte) {
	s := string(usaNursingJobsJSON_NY[:])
	dec := json.NewDecoder(strings.NewReader(s))
	var ny NursingJobs_NY
	if err := dec.Decode(&ny); err == io.EOF {
	} else if err != nil {
		log.Fatal(err)
	}
	var loper = 1
	time.Sleep(time.Second)
	for _, nursingJobs_NY := range ny {
		fmt.Println("Nursing-API i New York nr", loper)
		fmt.Printf("ID: %+v\nPosition title: %+v\nOrganization name: %+v\nRate Interval Code: %+v\nMinimum: %+v\nMaximum: %+v\nStart date: %+v\nEnd date: %+v\nLocations: %+v\nURL: %+v\n", nursingJobs_NY.ID, nursingJobs_NY.PositionTitle, nursingJobs_NY.OrganizationName, nursingJobs_NY.RateIntervalCode, nursingJobs_NY.Minimum, nursingJobs_NY.Maximum, nursingJobs_NY.StartDate, nursingJobs_NY.EndDate, nursingJobs_NY.Locations, nursingJobs_NY.URL)
		fmt.Println("---------------------------")
		loper++
	}
}
