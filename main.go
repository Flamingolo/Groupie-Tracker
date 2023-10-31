package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/artist", artistHandler)
	fmt.Println("Opening application at port :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
