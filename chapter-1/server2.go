// Server2 é um servidor de "eco" e contador
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	count += 1
	mu.Unlock()
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}

func counter(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Count: %d\n", count)
	mu.Unlock()
}
