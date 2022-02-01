package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "That`s okay")
		if err != nil {
			log.Print(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
