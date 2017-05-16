package klaer_IE

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Vareeksport struct {
	Dataset struct {
		Dimension struct {
			ImpEks struct {
				Label    string `json:"label"`
				Category struct {
					Index struct {
						Num2 int `json:"2"`
					} `json:"index"`
					Label struct {
						Num2 string `json:"2"`
					} `json:"label"`
				} `json:"category"`
			} `json:"ImpEks"`
			Land struct {
				Label    string `json:"label"`
				Category struct {
					Index struct {
						Num00 int `json:"00"`
						AF    int `json:"AF"`
						AL    int `json:"AL"`
						DZ    int `json:"DZ"`
						AS    int `json:"AS"`
						AD    int `json:"AD"`
						AO    int `json:"AO"`
						AI    int `json:"AI"`
						AG    int `json:"AG"`
						AN    int `json:"AN"`
						AR    int `json:"AR"`
						AM    int `json:"AM"`
						AW    int `json:"AW"`
						AZ    int `json:"AZ"`
						AU    int `json:"AU"`
						BS    int `json:"BS"`
						BH    int `json:"BH"`
						BD    int `json:"BD"`
						BB    int `json:"BB"`
						BE    int `json:"BE"`
						BZ    int `json:"BZ"`
						BJ    int `json:"BJ"`
						BM    int `json:"BM"`
						BT    int `json:"BT"`
						BO    int `json:"BO"`
						BQ    int `json:"BQ"`
						BA    int `json:"BA"`
						BW    int `json:"BW"`
						BR    int `json:"BR"`
						BN    int `json:"BN"`
						BG    int `json:"BG"`
						BF    int `json:"BF"`
						BU    int `json:"BU"`
						BI    int `json:"BI"`
						CA    int `json:"CA"`
						KY    int `json:"KY"`
						XC    int `json:"XC"`
						CL    int `json:"CL"`
						CX    int `json:"CX"`
						CO    int `json:"CO"`
						CK    int `json:"CK"`
						CR    int `json:"CR"`
						CU    int `json:"CU"`
						CW    int `json:"CW"`
						DK    int `json:"DK"`
						VI    int `json:"VI"`
						VG    int `json:"VG"`
						AE    int `json:"AE"`
						DO    int `json:"DO"`
						CF    int `json:"CF"`
						IO    int `json:"IO"`
						DJ    int `json:"DJ"`
						DM    int `json:"DM"`
						EC    int `json:"EC"`
						EG    int `json:"EG"`
						GQ    int `json:"GQ"`
						SV    int `json:"SV"`
						CI    int `json:"CI"`
						ER    int `json:"ER"`
						EE    int `json:"EE"`
						ET    int `json:"ET"`
						FK    int `json:"FK"`
						FJ    int `json:"FJ"`
						PH    int `json:"PH"`
						FI    int `json:"FI"`
						FR    int `json:"FR"`
						GF    int `json:"GF"`
						PF    int `json:"PF"`
						TF    int `json:"TF"`
						FO    int `json:"FO"`
						GA    int `json:"GA"`
						GM    int `json:"GM"`
						GE    int `json:"GE"`
						GH    int `json:"GH"`
						GI    int `json:"GI"`
						GD    int `json:"GD"`
						GL    int `json:"GL"`
						GP    int `json:"GP"`
						GU    int `json:"GU"`
						GT    int `json:"GT"`
						GG    int `json:"GG"`
						GN    int `json:"GN"`
						GW    int `json:"GW"`
						GY    int `json:"GY"`
						HT    int `json:"HT"`
						HM    int `json:"HM"`
						GR    int `json:"GR"`
						HN    int `json:"HN"`
						HK    int `json:"HK"`
						BY    int `json:"BY"`
						IN    int `json:"IN"`
						ID    int `json:"ID"`
						IQ    int `json:"IQ"`
						IR    int `json:"IR"`
						IE    int `json:"IE"`
						IS    int `json:"IS"`
						IM    int `json:"IM"`
						IL    int `json:"IL"`
						IT    int `json:"IT"`
						JM    int `json:"JM"`
						JP    int `json:"JP"`
						YE    int `json:"YE"`
						JE    int `json:"JE"`
						JO    int `json:"JO"`
						YU    int `json:"YU"`
						KH    int `json:"KH"`
						CM    int `json:"CM"`
						XB    int `json:"XB"`
						CV    int `json:"CV"`
						KZ    int `json:"KZ"`
						KE    int `json:"KE"`
						CN    int `json:"CN"`
						KG    int `json:"KG"`
						KI    int `json:"KI"`
						CC    int `json:"CC"`
						KM    int `json:"KM"`
						CD    int `json:"CD"`
						CG    int `json:"CG"`
						XK    int `json:"XK"`
						HR    int `json:"HR"`
						KW    int `json:"KW"`
						CY    int `json:"CY"`
						LA    int `json:"LA"`
						LV    int `json:"LV"`
						LS    int `json:"LS"`
						LB    int `json:"LB"`
						LR    int `json:"LR"`
						LY    int `json:"LY"`
						LI    int `json:"LI"`
						LT    int `json:"LT"`
						LU    int `json:"LU"`
						MO    int `json:"MO"`
						MG    int `json:"MG"`
						MK    int `json:"MK"`
						MW    int `json:"MW"`
						MY    int `json:"MY"`
						MV    int `json:"MV"`
						ML    int `json:"ML"`
						MT    int `json:"MT"`
						MA    int `json:"MA"`
						MH    int `json:"MH"`
						MQ    int `json:"MQ"`
						MR    int `json:"MR"`
						MU    int `json:"MU"`
						YT    int `json:"YT"`
						MX    int `json:"MX"`
						FM    int `json:"FM"`
						MD    int `json:"MD"`
						MC    int `json:"MC"`
						MN    int `json:"MN"`
						ME    int `json:"ME"`
						MS    int `json:"MS"`
						MZ    int `json:"MZ"`
						MM    int `json:"MM"`
						NA    int `json:"NA"`
						NR    int `json:"NR"`
						NL    int `json:"NL"`
						NP    int `json:"NP"`
						NZ    int `json:"NZ"`
						NI    int `json:"NI"`
						NE    int `json:"NE"`
						NG    int `json:"NG"`
						NU    int `json:"NU"`
						KP    int `json:"KP"`
						MP    int `json:"MP"`
						NF    int `json:"NF"`
						NC    int `json:"NC"`
						NT    int `json:"NT"`
						OM    int `json:"OM"`
						PK    int `json:"PK"`
						PW    int `json:"PW"`
						PA    int `json:"PA"`
						PG    int `json:"PG"`
						PY    int `json:"PY"`
						PE    int `json:"PE"`
						PN    int `json:"PN"`
						PL    int `json:"PL"`
						PT    int `json:"PT"`
						PR    int `json:"PR"`
						QA    int `json:"QA"`
						RE    int `json:"RE"`
						RO    int `json:"RO"`
						RU    int `json:"RU"`
						RW    int `json:"RW"`
						BL    int `json:"BL"`
						KN    int `json:"KN"`
						LC    int `json:"LC"`
						MF    int `json:"MF"`
						PM    int `json:"PM"`
						VC    int `json:"VC"`
						SB    int `json:"SB"`
						WS    int `json:"WS"`
						SM    int `json:"SM"`
						ST    int `json:"ST"`
						SA    int `json:"SA"`
						SN    int `json:"SN"`
						RS    int `json:"RS"`
						CS    int `json:"CS"`
						SC    int `json:"SC"`
						SL    int `json:"SL"`
						SG    int `json:"SG"`
						SX    int `json:"SX"`
						SK    int `json:"SK"`
						SI    int `json:"SI"`
						SO    int `json:"SO"`
						SU    int `json:"SU"`
						ES    int `json:"ES"`
						LK    int `json:"LK"`
						SH    int `json:"SH"`
						GB    int `json:"GB"`
						SD    int `json:"SD"`
						SR    int `json:"SR"`
						CH    int `json:"CH"`
						SE    int `json:"SE"`
						SZ    int `json:"SZ"`
						SY    int `json:"SY"`
						GS    int `json:"GS"`
						ZA    int `json:"ZA"`
						KR    int `json:"KR"`
						SS    int `json:"SS"`
						TJ    int `json:"TJ"`
						TW    int `json:"TW"`
						TZ    int `json:"TZ"`
						TH    int `json:"TH"`
						TG    int `json:"TG"`
						TK    int `json:"TK"`
						TO    int `json:"TO"`
						TT    int `json:"TT"`
						TD    int `json:"TD"`
						CZ    int `json:"CZ"`
						TN    int `json:"TN"`
						TM    int `json:"TM"`
						TC    int `json:"TC"`
						TV    int `json:"TV"`
						TR    int `json:"TR"`
						DE    int `json:"DE"`
						UG    int `json:"UG"`
						UA    int `json:"UA"`
						HU    int `json:"HU"`
						UY    int `json:"UY"`
						US    int `json:"US"`
						UM    int `json:"UM"`
						UZ    int `json:"UZ"`
						VU    int `json:"VU"`
						VA    int `json:"VA"`
						VE    int `json:"VE"`
						XI    int `json:"XI"`
						PS    int `json:"PS"`
						EH    int `json:"EH"`
						VN    int `json:"VN"`
						WF    int `json:"WF"`
						YD    int `json:"YD"`
						ZR    int `json:"ZR"`
						ZM    int `json:"ZM"`
						ZW    int `json:"ZW"`
						AT    int `json:"AT"`
						TP    int `json:"TP"`
						TL    int `json:"TL"`
						DD    int `json:"DD"`
						AX    int `json:"AX"`
						ZZ    int `json:"ZZ"`
					} `json:"index"`
					Label struct {
						Num00 string `json:"00"`
						AF    string `json:"AF"`
						AL    string `json:"AL"`
						DZ    string `json:"DZ"`
						AS    string `json:"AS"`
						AD    string `json:"AD"`
						AO    string `json:"AO"`
						AI    string `json:"AI"`
						AG    string `json:"AG"`
						AN    string `json:"AN"`
						AR    string `json:"AR"`
						AM    string `json:"AM"`
						AW    string `json:"AW"`
						AZ    string `json:"AZ"`
						AU    string `json:"AU"`
						BS    string `json:"BS"`
						BH    string `json:"BH"`
						BD    string `json:"BD"`
						BB    string `json:"BB"`
						BE    string `json:"BE"`
						BZ    string `json:"BZ"`
						BJ    string `json:"BJ"`
						BM    string `json:"BM"`
						BT    string `json:"BT"`
						BO    string `json:"BO"`
						BQ    string `json:"BQ"`
						BA    string `json:"BA"`
						BW    string `json:"BW"`
						BR    string `json:"BR"`
						BN    string `json:"BN"`
						BG    string `json:"BG"`
						BF    string `json:"BF"`
						BU    string `json:"BU"`
						BI    string `json:"BI"`
						CA    string `json:"CA"`
						KY    string `json:"KY"`
						XC    string `json:"XC"`
						CL    string `json:"CL"`
						CX    string `json:"CX"`
						CO    string `json:"CO"`
						CK    string `json:"CK"`
						CR    string `json:"CR"`
						CU    string `json:"CU"`
						CW    string `json:"CW"`
						DK    string `json:"DK"`
						VI    string `json:"VI"`
						VG    string `json:"VG"`
						AE    string `json:"AE"`
						DO    string `json:"DO"`
						CF    string `json:"CF"`
						IO    string `json:"IO"`
						DJ    string `json:"DJ"`
						DM    string `json:"DM"`
						EC    string `json:"EC"`
						EG    string `json:"EG"`
						GQ    string `json:"GQ"`
						SV    string `json:"SV"`
						CI    string `json:"CI"`
						ER    string `json:"ER"`
						EE    string `json:"EE"`
						ET    string `json:"ET"`
						FK    string `json:"FK"`
						FJ    string `json:"FJ"`
						PH    string `json:"PH"`
						FI    string `json:"FI"`
						FR    string `json:"FR"`
						GF    string `json:"GF"`
						PF    string `json:"PF"`
						TF    string `json:"TF"`
						FO    string `json:"FO"`
						GA    string `json:"GA"`
						GM    string `json:"GM"`
						GE    string `json:"GE"`
						GH    string `json:"GH"`
						GI    string `json:"GI"`
						GD    string `json:"GD"`
						GL    string `json:"GL"`
						GP    string `json:"GP"`
						GU    string `json:"GU"`
						GT    string `json:"GT"`
						GG    string `json:"GG"`
						GN    string `json:"GN"`
						GW    string `json:"GW"`
						GY    string `json:"GY"`
						HT    string `json:"HT"`
						HM    string `json:"HM"`
						GR    string `json:"GR"`
						HN    string `json:"HN"`
						HK    string `json:"HK"`
						BY    string `json:"BY"`
						IN    string `json:"IN"`
						ID    string `json:"ID"`
						IQ    string `json:"IQ"`
						IR    string `json:"IR"`
						IE    string `json:"IE"`
						IS    string `json:"IS"`
						IM    string `json:"IM"`
						IL    string `json:"IL"`
						IT    string `json:"IT"`
						JM    string `json:"JM"`
						JP    string `json:"JP"`
						YE    string `json:"YE"`
						JE    string `json:"JE"`
						JO    string `json:"JO"`
						YU    string `json:"YU"`
						KH    string `json:"KH"`
						CM    string `json:"CM"`
						XB    string `json:"XB"`
						CV    string `json:"CV"`
						KZ    string `json:"KZ"`
						KE    string `json:"KE"`
						CN    string `json:"CN"`
						KG    string `json:"KG"`
						KI    string `json:"KI"`
						CC    string `json:"CC"`
						KM    string `json:"KM"`
						CD    string `json:"CD"`
						CG    string `json:"CG"`
						XK    string `json:"XK"`
						HR    string `json:"HR"`
						KW    string `json:"KW"`
						CY    string `json:"CY"`
						LA    string `json:"LA"`
						LV    string `json:"LV"`
						LS    string `json:"LS"`
						LB    string `json:"LB"`
						LR    string `json:"LR"`
						LY    string `json:"LY"`
						LI    string `json:"LI"`
						LT    string `json:"LT"`
						LU    string `json:"LU"`
						MO    string `json:"MO"`
						MG    string `json:"MG"`
						MK    string `json:"MK"`
						MW    string `json:"MW"`
						MY    string `json:"MY"`
						MV    string `json:"MV"`
						ML    string `json:"ML"`
						MT    string `json:"MT"`
						MA    string `json:"MA"`
						MH    string `json:"MH"`
						MQ    string `json:"MQ"`
						MR    string `json:"MR"`
						MU    string `json:"MU"`
						YT    string `json:"YT"`
						MX    string `json:"MX"`
						FM    string `json:"FM"`
						MD    string `json:"MD"`
						MC    string `json:"MC"`
						MN    string `json:"MN"`
						ME    string `json:"ME"`
						MS    string `json:"MS"`
						MZ    string `json:"MZ"`
						MM    string `json:"MM"`
						NA    string `json:"NA"`
						NR    string `json:"NR"`
						NL    string `json:"NL"`
						NP    string `json:"NP"`
						NZ    string `json:"NZ"`
						NI    string `json:"NI"`
						NE    string `json:"NE"`
						NG    string `json:"NG"`
						NU    string `json:"NU"`
						KP    string `json:"KP"`
						MP    string `json:"MP"`
						NF    string `json:"NF"`
						NC    string `json:"NC"`
						NT    string `json:"NT"`
						OM    string `json:"OM"`
						PK    string `json:"PK"`
						PW    string `json:"PW"`
						PA    string `json:"PA"`
						PG    string `json:"PG"`
						PY    string `json:"PY"`
						PE    string `json:"PE"`
						PN    string `json:"PN"`
						PL    string `json:"PL"`
						PT    string `json:"PT"`
						PR    string `json:"PR"`
						QA    string `json:"QA"`
						RE    string `json:"RE"`
						RO    string `json:"RO"`
						RU    string `json:"RU"`
						RW    string `json:"RW"`
						BL    string `json:"BL"`
						KN    string `json:"KN"`
						LC    string `json:"LC"`
						MF    string `json:"MF"`
						PM    string `json:"PM"`
						VC    string `json:"VC"`
						SB    string `json:"SB"`
						WS    string `json:"WS"`
						SM    string `json:"SM"`
						ST    string `json:"ST"`
						SA    string `json:"SA"`
						SN    string `json:"SN"`
						RS    string `json:"RS"`
						CS    string `json:"CS"`
						SC    string `json:"SC"`
						SL    string `json:"SL"`
						SG    string `json:"SG"`
						SX    string `json:"SX"`
						SK    string `json:"SK"`
						SI    string `json:"SI"`
						SO    string `json:"SO"`
						SU    string `json:"SU"`
						ES    string `json:"ES"`
						LK    string `json:"LK"`
						SH    string `json:"SH"`
						GB    string `json:"GB"`
						SD    string `json:"SD"`
						SR    string `json:"SR"`
						CH    string `json:"CH"`
						SE    string `json:"SE"`
						SZ    string `json:"SZ"`
						SY    string `json:"SY"`
						GS    string `json:"GS"`
						ZA    string `json:"ZA"`
						KR    string `json:"KR"`
						SS    string `json:"SS"`
						TJ    string `json:"TJ"`
						TW    string `json:"TW"`
						TZ    string `json:"TZ"`
						TH    string `json:"TH"`
						TG    string `json:"TG"`
						TK    string `json:"TK"`
						TO    string `json:"TO"`
						TT    string `json:"TT"`
						TD    string `json:"TD"`
						CZ    string `json:"CZ"`
						TN    string `json:"TN"`
						TM    string `json:"TM"`
						TC    string `json:"TC"`
						TV    string `json:"TV"`
						TR    string `json:"TR"`
						DE    string `json:"DE"`
						UG    string `json:"UG"`
						UA    string `json:"UA"`
						HU    string `json:"HU"`
						UY    string `json:"UY"`
						US    string `json:"US"`
						UM    string `json:"UM"`
						UZ    string `json:"UZ"`
						VU    string `json:"VU"`
						VA    string `json:"VA"`
						VE    string `json:"VE"`
						XI    string `json:"XI"`
						PS    string `json:"PS"`
						EH    string `json:"EH"`
						VN    string `json:"VN"`
						WF    string `json:"WF"`
						YD    string `json:"YD"`
						ZR    string `json:"ZR"`
						ZM    string `json:"ZM"`
						ZW    string `json:"ZW"`
						AT    string `json:"AT"`
						TP    string `json:"TP"`
						TL    string `json:"TL"`
						DD    string `json:"DD"`
						AX    string `json:"AX"`
						ZZ    string `json:"ZZ"`
					} `json:"label"`
				} `json:"category"`
			} `json:"Land"`
			SITC struct {
				Label    string `json:"label"`
				Category struct {
					Index struct {
						NAMING_FAILED int `json:"-"`
					} `json:"index"`
					Label struct {
						NAMING_FAILED string `json:"-"`
					} `json:"label"`
				} `json:"category"`
			} `json:"SITC"`
			Tid struct {
				Label    string `json:"label"`
				Category struct {
					Index struct {
						Two016M03 int `json:"2016M03"`
						Two016M04 int `json:"2016M04"`
						Two016M05 int `json:"2016M05"`
						Two016M06 int `json:"2016M06"`
						Two016M07 int `json:"2016M07"`
						Two016M08 int `json:"2016M08"`
						Two016M09 int `json:"2016M09"`
						Two016M10 int `json:"2016M10"`
						Two016M11 int `json:"2016M11"`
						Two016M12 int `json:"2016M12"`
						Two017M01 int `json:"2017M01"`
						Two017M02 int `json:"2017M02"`
						Two017M03 int `json:"2017M03"`
					} `json:"index"`
					Label struct {
						Two016M03 string `json:"2016M03"`
						Two016M04 string `json:"2016M04"`
						Two016M05 string `json:"2016M05"`
						Two016M06 string `json:"2016M06"`
						Two016M07 string `json:"2016M07"`
						Two016M08 string `json:"2016M08"`
						Two016M09 string `json:"2016M09"`
						Two016M10 string `json:"2016M10"`
						Two016M11 string `json:"2016M11"`
						Two016M12 string `json:"2016M12"`
						Two017M01 string `json:"2017M01"`
						Two017M02 string `json:"2017M02"`
						Two017M03 string `json:"2017M03"`
					} `json:"label"`
				} `json:"category"`
			} `json:"Tid"`
			ContentsCode struct {
				Label    string `json:"label"`
				Category struct {
					Index struct {
						Verdi int `json:"Verdi"`
					} `json:"index"`
					Label struct {
						Verdi string `json:"Verdi"`
					} `json:"label"`
					Unit struct {
						Verdi struct {
							Base string `json:"base"`
						} `json:"Verdi"`
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
		Label   string        `json:"label"`
		Source  string        `json:"source"`
		Updated time.Time     `json:"updated"`
		Value   []interface{} `json:"value"`
	} `json:"dataset"`
}

func HentData(urlImportEksportKlaer string) {
	for {
		respons, err := http.Get(urlImportEksportKlaer)
		if err != nil {
			panic(err.Error())
		}
		importEksportJSON, err := ioutil.ReadAll(respons.Body)
		if err != nil {
			panic(err.Error())
		}
		go mellomLagringJSON(importEksportJSON)
		go mellomLagringTXT(importEksportJSON)
		go convertJSON(importEksportJSON)
		time.Sleep(5 * time.Minute)
	}
}

func mellomLagringJSON(importEksportJSON []byte) {
	err := ioutil.WriteFile("klaer_IE/importEksportKlær.json", importEksportJSON, 0666)
	if err != nil {
		panic(err.Error())
	}
}

func mellomLagringTXT(importEksportJSON []byte) {
	err := ioutil.WriteFile("klaer_IE/importEksportKlær.txt", importEksportJSON, 0666)
	if err != nil {
		panic(err.Error())
	}
}

func convertJSON(importEksportJSON []byte) {
	var i Vareeksport
	err := json.Unmarshal(importEksportJSON, &i)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("----Import/Eksport Klær----")
	fmt.Printf("Label: %+v\nSource: %+v\nOppdatert: %+v\nStatistikk for: %+v\nContentscode label: %+v\n--Category--\nIndex: %+v\nLabel: %+v\nUnit: %+v\n", i.Dataset.Label, i.Dataset.Source, i.Dataset.Updated, i.Dataset.Dimension.ImpEks.Label, i.Dataset.Dimension.ContentsCode.Label, i.Dataset.Dimension.ContentsCode.Category.Index.Verdi, i.Dataset.Dimension.ContentsCode.Category.Label.Verdi, i.Dataset.Dimension.ContentsCode.Category.Unit.Verdi)
	fmt.Printf("---%+v:---\n", i.Dataset.Dimension.Land.Label)
	fmt.Printf("Land: %+v, Index: %+v\n", i.Dataset.Dimension.Land.Category.Label.AF, i.Dataset.Dimension.Land.Category.Index.AF)
	fmt.Printf("Land: %+v, Index: %+v\n", i.Dataset.Dimension.Land.Category.Label.GB, i.Dataset.Dimension.Land.Category.Index.GB)
	fmt.Println("----Import/Eksport----")
}
