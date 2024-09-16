package main

import (
	"fmt"
	"log"
	"net/http"
)
funchelloHandler( w http.ResponseWriter , r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found",http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported",http.StatusNotFound)
	}
}

func main() {
	fileServer := http.fileServer(http.Die("./static"))
	http.Handle("/".fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting server at prot 8080\n")
	if err := http.ListernAndServer(":8080",nil); err !=nil {
		log.Fatal(err); err !=nill {
			log.Fatal(err)
		}
	}
}