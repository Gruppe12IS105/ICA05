package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {

	const jsonStream = "https://search.digitalgov.gov/developer/jobs.html"
	/*
		res, err := http.Get("https://search.digitalgov.gov/developer/jobs.html")
		if err != nil {
			panic(err.Error())
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err.Error())
		}
	*/
	type Parameters struct {
		querry           string
		orginazation_ids string
		hl               string
		size             int64
		from             string
		tags             string
		lat_lon          float64
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var p Parameters
		if err := dec.Decode(&p); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", p.tags, p.orginazation_ids)
		fmt.Printf("Tags are: %.2f\n",
			p.tags)
		fmt.Printf("Orginazation id: %.f\n", p.orginazation_ids)
	}
}
