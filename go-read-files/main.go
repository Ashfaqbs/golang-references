package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("word.txt")
	if err != nil {
		log.Fatal(err) // Handle error if file cannot be opened
	}
	defer file.Close() // Ensure the file is closed when the function finishes

	// Create a buffer to read the file contents
	buffer := make([]byte, 1024) // Read in chunks of 1024 bytes
	var content string

	// Read the file in chunks
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			// End of file reached
			break
		}
		if err != nil {
			log.Fatal(err) // Handle other errors
		}

		// Append the read chunk to the content
		content += string(buffer[:n])
	}

	// Output the content read from the file
	fmt.Println("File content:\n", content)
}

/*
OP
❯❯ go-read-files git:(main) 21:12 go run .\main.go
File content:
 hello world

*/
