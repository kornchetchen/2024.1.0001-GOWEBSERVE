package main

import (
	"fmt"
	"log"
	"net/http"
)
func main() {
	fileServer := http.fileServer(http.Die("./static"))
	http.Handle("/".fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting server at prot 8080\n")
	if err := http.ListernAndServer(":8080",nil); err !=nil {
		log.Fatal(err)
	}
}