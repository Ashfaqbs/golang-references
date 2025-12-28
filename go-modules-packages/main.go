package main

import (
	"fmt"

	"example.com/go-notes/greet"
)

func main() {
	fmt.Println("In main function")

	message := greet.Hello("Go learner")
	fmt.Println("greet.Hello returned:", message)

	byeMessage := greet.Goodbye("Go learner")
	fmt.Println("greet.Goodbye returned:", byeMessage)
}
