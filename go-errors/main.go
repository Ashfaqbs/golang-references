package main

import (
	"errors"
	"fmt"
	"os"
)

// 1. Function that can fail, returns (result, error) similiar to throws in JAVA
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	result := a / b
	return result, nil
}

// understanding Defer :

func exampleFunction() {
	defer fmt.Println("First deferred statement.")
	defer fmt.Println("Second deferred statement.")
	defer fmt.Println("Third deferred statement.")

	fmt.Println("Function starts here.")
}

// Understanding panic :

// openConfigFile tries to open the config file. If it fails, it panics with a message.
// Note that when panic is called, the function does not return normally; it immediately stops execution.
func openConfigFile(filename string) *os.File {
	// Try to open the configuration file
	file, err := os.Open(filename)
	if err != nil {
		// If there's an error, panic with a message.
		// This stops the function execution immediately and does not return a value.
		panic(fmt.Sprintf("Critical error: cannot open config file '%s': %v", filename, err))
	}
	// If successful, return the file pointer (this won't be reached if panic is called).
	return file
}

// Understanding recover and how defer should be used with Recover
// safeOpenConfigFile tries to open a config file and recovers from panic if it occurs.
func safeOpenConfigFile(filename string) (file *os.File, errMsg string) {
	// Defer function will be executed when safeOpenConfigFile exits, even if it's due to a panic.
	defer func() {
		// If a panic occurs, recover will capture it.
		if r := recover(); r != nil {
			// Handle the panic (recover from it)
			errMsg = fmt.Sprintf("Recovered from panic: %v", r)
		}
	}()

	// Attempt to open the file (this could panic)
	file = openConfigFile(filename)
	return file, ""
}

func main() {

	// A. Normal error handling with (value, error)
	fmt.Println("== divide with valid input ==")
	value, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", value)
	}

	fmt.Println("\n== divide with division by zero ==")
	value, err = divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", value)
	}

	// Understanding defer :

	exampleFunction()

	// Understanding Panic
	// Specify the configuration file
	configFile := "config.txt"

	// Attempt to open the config file. If panic occurs, the program will terminate here.
	file := openConfigFile(configFile)
	// Ensure the file gets closed when the program exits
	defer file.Close()

	// Proceed with the rest of the program, assuming the file is successfully opened
	fmt.Println("Config file opened successfully!")
	// Additional logic like reading from the file could go here

	/*
	   panic: Critical error: cannot open config file 'config.txt': open config.txt: The system cannot find the file specified.

	   goroutine 1 [running]:
	   main.openConfigFile({0xc43678, 0xa})
	           C:/tmp/projects/golang-references/go-errors/main.go:38 +0xb9
	   main.main()
	           C:/tmp/projects/golang-references/go-errors/main.go:99 +0x196
	   exit status 2
	*/

	// Understanding recover and how defer should be used with Recover

	// Attempt to open the config file with recovery in place
	file, errMsg := safeOpenConfigFile(configFile)
	if errMsg != "" {
		// Panic was recovered, print the error message
		fmt.Println(errMsg)
	} else {
		// If file opened successfully, proceed as normal
		defer file.Close()
		fmt.Println("Config file opened successfully!")
	}
}
