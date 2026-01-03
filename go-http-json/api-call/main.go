package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// API URL
	url := "https://kimiquotes.pages.dev/api/quote"

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set headers if needed
	req.Header.Set("User-Agent", "Go-Http-Client/1.0")

	// Send the request
	// Create a new http.Client and get a pointer to it.
	// The & operator is used to avoid copying the entire struct and to allow sharing and modifying the same client instance across requests
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Use io.ReadAll instead of ioutil.ReadAll
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Print out the response (JSON quote in your case)
	fmt.Println("Response from API:")
	fmt.Println(string(body))
}

/*
OP :

api-call git:(main) 19:21 go run .\main.go
Response from API:
{"id":20,"quote":"OK...","year":2013}

*/
