package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func updateConfig(conf *Config, next string, previous string) {
	conf.Next = next
	conf.Previous = previous
}

func getLocationAreas(url string, conf *Config) []LocationResults {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var myLocationAreas locationAreasAPIResponse
	err = json.Unmarshal(body, &myLocationAreas)
	if err != nil {
		log.Fatal(err)
	}
	updateConfig(conf, myLocationAreas.Next, myLocationAreas.Previous)
	return myLocationAreas.Results
}

type locationAreasAPIResponse struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []LocationResults `json:"results"`
}

type LocationResults struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
