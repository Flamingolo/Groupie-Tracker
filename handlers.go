package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relation     Relation `json:"relation"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	artistURL := "https://groupietrackers.herokuapp.com/api/artists"

	if r.URL.Path != "/"{
		http.Error(w, "Bad request: Invalid URL path", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(artistURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusInternalServerError {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			http.Error(w, fmt.Sprintf("API request failed: %s", resp.Status), resp.StatusCode)
		}
		return
	}

	var artist []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temp, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	if err := temp.Execute(w, artist); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.Query().Get("id")
	relationURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", URL)
	artistURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", URL)

	rel, err := http.Get(relationURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer rel.Body.Close()

	art, err := http.Get(artistURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer art.Body.Close()

	var artist Artist
	var relations Relation
	if err := json.NewDecoder(rel.Body).Decode(&relations); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewDecoder(art.Body).Decode(&artist); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	formatedLocations := make(map[string][]string)
	for location, dates := range relations.DatesLocations {
		formatedLoc := formatArtist(location)
		formatedLocations[formatedLoc] = dates
	}
	relations.DatesLocations = formatedLocations

	tmpl, err := template.ParseFiles("static/artistpage.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, struct {
		Artist    Artist
		Relations Relation
	}{
		Artist:    artist,
		Relations: relations,
	}); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func formatArtist(input string) string {
	parts := strings.Split(input, "_")
	var formatted []string
	for _, part := range parts {
		if part != "" {
			formatted = append(formatted, strings.Title(part))
		}
	}
	formattedName := strings.Join(formatted, " ")
	return formattedName
}
