package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Artist []struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Relations    map[string][]string `json:"relations"`
}

type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func mainHandler(w http.ResponseWriter, r *http.Request) {

}

func artistHandler(w http.ResponseWriter, r *http.Request) {

}

func getJSON(URL string, target interface{}) {
	resp, err := http.Get(URL)
	checkErr(err)
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&target)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
