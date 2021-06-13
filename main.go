package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
)

//https://raw.githubusercontent.com/MarioTataje/tb2-dataset/main/fallecidos_covid.csv

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

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"message": "get called"}`))
		if err != nil {
			return
		}
	case "POST":
		w.WriteHeader(http.StatusCreated)
		_, err := w.Write([]byte(`{"message": "post called"}`))
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
	for _, line := range data{
		fmt.Println(line)
	}
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8090", nil))
}