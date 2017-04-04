package API

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type NursingJobs []struct {
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

func USANursingJobs(urlNursing string) {
	respons, err := http.Get(urlNursing)
	if err != nil {
		panic(err.Error())
	}
	usaNursingJobsJSON, err := ioutil.ReadAll(respons.Body)
	if err != nil {
		panic(err.Error())
	}
	convertNursingJobs(usaNursingJobsJSON)
}
func convertNursingJobs(usaNursingJobsJSON []byte) {
	s := string(usaNursingJobsJSON[:])
	dec := json.NewDecoder(strings.NewReader(s))
	var n NursingJobs
	if err := dec.Decode(&n); err == io.EOF {
	} else if err != nil {
		log.Fatal(err)
	}
	var loper = 1
	for _, nursingJobs := range n {
		fmt.Println("Nursing-API nr", loper)
		fmt.Printf("ID: %+v\nPosition title: %+v\nOrganization name: %+v\nRate Interval Code: %+v\nMinimum: %+v\nMaximum: %+v\nStart date: %+v\nEnd date: %+v\nLocations: %+v\nURL: %+v\n", nursingJobs.ID, nursingJobs.PositionTitle, nursingJobs.OrganizationName, nursingJobs.RateIntervalCode, nursingJobs.Minimum, nursingJobs.Maximum, nursingJobs.StartDate, nursingJobs.EndDate, nursingJobs.Locations, nursingJobs.URL)
		fmt.Println("---------------------------")
		loper++
	}
}
