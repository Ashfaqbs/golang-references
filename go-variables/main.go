package main

import "fmt"

func main() {
	// 1. Boolean
	var isGoFun bool = true
	fmt.Println("isGoFun:", isGoFun)

	// 2. Integer
	var age int = 25
	fmt.Println("age:", age)

	// 3. Float (decimal number)
	var pi float64 = 3.14159
	fmt.Println("pi:", pi)

	// 4. String (text)
	var language string = "Go language"
	fmt.Println("language:", language)

	// 5. Short variable declaration with :=
	//    Type is inferred from the value on the right-hand side.
	count := 10         // inferred as int
	ratio := 0.75       // inferred as float64
	title := "Go notes" // inferred as string

	fmt.Println("count:", count)
	fmt.Println("ratio:", ratio)
	fmt.Println("title:", title)

	// 6. Printing type information as well
	fmt.Printf("Type of isGoFun: %T\n", isGoFun)
	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of pi: %T\n", pi)
	fmt.Printf("Type of language: %T\n", language)
	fmt.Printf("Type of count: %T\n", count)
	fmt.Printf("Type of ratio: %T\n", ratio)
	fmt.Printf("Type of title: %T\n", title)
}
