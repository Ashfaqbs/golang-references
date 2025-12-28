package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Reading an environment variable using os.Getenv
	API_KEY := os.Getenv("GEMINI_API_KEY") // Get DB_HOST environment variable
	// Get DB_PORT environment variable
	if API_KEY == "" {
		log.Fatal("Environment variables GEMINI_API_KEY is not set")
	}

	// Printing the values of the environment variables
	fmt.Println("GEMINI_API_KEY :", API_KEY)

	// Reading command-line arguments using os.Args
	// os.Args[0] is the name of the program itself
	// os.Args[1:] contains the actual arguments passed
	if len(os.Args) < 2 {
		log.Fatal("Please provide a command-line argument")
	}
	arg := os.Args[1] // Get the first command-line argument
	fmt.Println("Command-line argument:", arg)

	// Example logic based on the CLI argument
	if arg == "start" {
		fmt.Println("Starting the application...")
	} else if arg == "stop" {
		fmt.Println("Stopping the application...")
	} else {
		fmt.Println("Invalid command-line argument. Use 'start' or 'stop'.")
	}
}

/*
OP

❯❯ go-env-vars-arguments git:(main)  20:55 go run .\main.go start
GEMINI_API_KEY : APIKEY
Command-line argument: start
Starting the application...


*/
