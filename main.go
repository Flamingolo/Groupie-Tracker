package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/artistpage/", artistHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
