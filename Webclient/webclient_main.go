package main

import "./API"

/*
  Airforce: "https://api.usa.gov/jobs/search.json?query=jobs+at+the+va&organization_ids=AF"
  Nursing: "https://api.usa.gov/jobs/search.json?query=nursing+jobs"
	Nursing in New York: "https://api.usa.gov/jobs/search.json?query=nursing+jobs+in+ny"
	Statistikk over Laks i Norge: "https://data.ssb.no/api/v0/dataset/1120.json?lang=en"
	Statistikk over Døddfall i Norge: "https://data.ssb.no/api/v0/dataset/1104.json?lang=en"

*/

func main() {
	var urlNursing = "https://api.usa.gov/jobs/search.json?query=nursing+jobs"
	var urlAirForce = "https://api.usa.gov/jobs/search.json?query=jobs+at+the+va&organization_ids=AF"
	var urlNursing_NY = "https://api.usa.gov/jobs/search.json?query=nursing+jobs+in+ny"
	var urlLakseStats = "https://data.ssb.no/api/v0/dataset/1120.json?lang=en"
	var urlDodStats = "https://data.ssb.no/api/v0/dataset/1104.json?lang=en"

	/*
		Alle go-rutinene venter 1 sek mellom hverandre for at ikke alle dataene
		skal bli skrevet ut over hverandre (opplysninger fra DodStats kunne plutselig ligge
		midt i opplysninger om LakseInfo)
	*/
	go API.USANursingJobs(urlNursing)
	go API.USANursingJobs_NY(urlNursing_NY)
	go API.LakseInfo(urlLakseStats)
	go API.GetDodStats(urlDodStats)
	API.USAAirForceJobs(urlAirForce)
}
