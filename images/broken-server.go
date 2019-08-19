package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

const VERSION = "0"

func status_handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "I'm alive on %s / Version: %v", hostname, VERSION)
	fmt.Printf("200 %s\n", r.URL.Path)
}

func error_handler(w http.ResponseWriter, r *http.Request) {
	if rand.Intn(2)%2 == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal server error"))
	        fmt.Printf("500 %s\n", r.URL.Path)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("503 - Service Unavailable"))
	        fmt.Printf("503 %s\n", r.URL.Path)
	}
}

func main() {
	fmt.Printf("broken server started - version %v\n", VERSION)
	http.HandleFunc("/solr/index/admin/ping", status_handler)
	http.HandleFunc("/solr/admin/info/system", status_handler)
	http.HandleFunc("/", error_handler)
	http.ListenAndServe(":8983", nil)
}
