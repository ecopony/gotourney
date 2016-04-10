package main

import (
	"net/http"
	"log"

	"fmt"
)

func main() {
	fmt.Println("starting up")
	router := NewRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Listen and serve: ", err)
	}
}
