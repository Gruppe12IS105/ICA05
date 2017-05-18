// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
package main

import (
	"./generell_IE"
	"./siste13mnd_IE"
	"./vaapen_IE"
	"./levendeDyr_IE"
	"./klaer_IE"
)

func main() {
//API-er
	var urlImportEksport = "http://data.ssb.no/api/v0/dataset/1136.json?lang=no"
	var urlImportEksport13mnd = "http://data.ssb.no/api/v0/dataset/1138.json?lang=no"
	var urlImportEksportVaapen = "http://data.ssb.no/api/v0/dataset/1178.json?lang=no"
	var urlImportEksportKlaer = "http://data.ssb.no/api/v0/dataset/1164.json?lang=no"
	var urlImportEksportDyr = "http://data.ssb.no/api/v0/dataset/321830.json?lang=no"
//Kjører de andre go-filene
	go siste13mnd_IE.HentData(urlImportEksport13mnd)
	go vaapen_IE.HentData(urlImportEksportVaapen)
	go levendeDyr_IE.HentData(urlImportEksportDyr)
	go klaer_IE.HentData(urlImportEksportKlaer)
	generell_IE.HentData(urlImportEksport)
}
