package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CovidDeath struct {
	Id             string  `json:"id"`
	CutOffDate     string  `json:"cut_off_date"`
	DeathDate      string  `json:"death_date"`
	Age            string  `json:"age"`
	Sex            string  `json:"sex"`
	Classification string  `json:"classification"`
	Ubigeo         string  `json:"ubigeo"`
	Region         string  `json:"region"`
    Province       string  `json:"province"`
	District       string  `json:"district"`
}

func readCSVFromUrl(url string)([][]string, error){
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	reader := csv.NewReader(resp.Body)
	reader.LazyQuotes = true
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getCovidDeathsFromData(data[][] string) []CovidDeath {
	var covidDeaths []CovidDeath
	for _, line := range data{
		covidDeath := CovidDeath{
			CutOffDate:     line[0],
			Id:             line[1],
			DeathDate:      line[2],
			Age:            line[3],
			Sex:            line[4],
			Classification: line[5],
			Ubigeo:         line[6],
			Region:         line[7],
			Province:       line[8],
			District:       line[9],
		}
		covidDeaths = append(covidDeaths, covidDeath)
	}
	return covidDeaths
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"message": "This is the TB2 API"}`))
		if err != nil {
			return
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(`{"message": "not found"}`))
		if err != nil {
			return
		}
	}
}

func covidDeathsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"message": "This is the covid deaths"}`))
		if err != nil {
			return
		}
	case "POST":
		w.WriteHeader(http.StatusCreated)
		_, err := w.Write([]byte(`{"message": "This is the covid deaths post"}`))
		if err != nil {
			return
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(`{"message": "not found"}`))
		if err != nil {
			return
		}
	}
}

func main() {
	url := "https://raw.githubusercontent.com/MarioTataje/tb2-dataset/main/fallecidos_covid.csv"
	data, err := readCSVFromUrl(url)
	if err != nil {
		panic(err)
	}
	covidDeaths := getCovidDeathsFromData(data)
	fmt.Println(covidDeaths)
	http.HandleFunc("/", home)
	http.HandleFunc("/covid-deaths", covidDeathsController)
	log.Fatal(http.ListenAndServe(":8090", nil))
}