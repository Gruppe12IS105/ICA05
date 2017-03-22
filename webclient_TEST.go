package main

import (
  //"encoding/json"
  "net/http"
  "log"
  "fmt"
  "io/ioutil"
  //"os"
)

// NOTE: denne må startes som en goroutine for hver URL/datakilde
// som dere skal hente inn-data fra.

func doGet(url string) {
  response, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  } else {
      defer response.Body.Close()
      contents, err := ioutil.ReadAll(response.Body)
      if err != nil {
        log.Fatal(err)
      }
      fmt.Println("%q", contents)
      fmt.Println("   ", response.StatusCode)
      hdr := response.Header
      for key, value := range hdr {
        fmt.Println("   ", key, ":", value)
    }
  }
}

func main() {
  url := "http://samples.openweathermap.org/data/2.5/weather?zip=94040,us&appid=b1b15e88fa797225412429c1c50c122a1"
//  url := os.Args[1]
  doGet(url)
}
