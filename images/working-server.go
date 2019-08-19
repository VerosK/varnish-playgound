package main

import (
	"fmt"
	"net/http"
	"os"
)

const VERSION = "0"

func valid_handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "I'm alive on %s / Version: %v", hostname, VERSION)
	fmt.Printf("200 %s\n", r.URL.Path)
}


func main() {
	fmt.Printf("working server started - version %v\n", VERSION)
	http.HandleFunc("/", valid_handler)
	http.ListenAndServe(":8983", nil)
}
