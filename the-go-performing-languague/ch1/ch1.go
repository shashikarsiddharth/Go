package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Purpose:To echo the command line arguments
func echo() {
	var s string
	// var sep string
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	s = strings.Join(os.Args[1:], " ")
	fmt.Println(s)
}

// Purpose:To print count of chars in string
func charCounter(input string) map[string]int {
	var char = make(map[string]int)
	for i := range input {
		char[string(input[i])]++
	}
	return char
}

// Purpose: To print duplicate input lines along with count
func printDuplicate() {
	var count = make(map[string]int)
	inp := bufio.NewScanner(os.Stdin)
	for inp.Scan() {
		if inp.Err() != nil {
			log.Printf("failed to scan input from stdin: %v", inp.Err())
		}
		count[inp.Text()]++
	}
	fmt.Println(count)
}

// Purpose: To fetch data from an endpoint
func fetchData(endpoint string) {
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Panicf("failed to get response from endpoint: %v", err)
	}

	r, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("failed to read response: %v", err)
	}
	resp.Body.Close()
	fmt.Printf("%s", r)
}

// Purpose: To start a simple http server on port 8080
func startServer() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World\n")
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func main() {
	// fmt.Println(charCounter("HelloWorld"))
	// printDuplicate()
	// fetchData("https://example.com/")
	startServer()
}
