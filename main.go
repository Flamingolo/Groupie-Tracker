package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//Web design
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/artist", artistHandler)
	fmt.Println("Opening application at port http://localhost:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
