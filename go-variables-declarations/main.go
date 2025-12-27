package main

import "fmt"

func main() {
	// 1. Explicit type with initialization
	var a int = 10
	fmt.Println("a:", a)

	// 2. Explicit type without initialization (zero value)
	var b int
	fmt.Println("b (zero value):", b)

	// 3. 'var' with type inferred from value
	var c = 3.14
	fmt.Println("c:", c)

	// 4. Short declaration with :=
	d := "Go"
	fmt.Println("d:", d)

	// 5. Multiple variables of the same type in one declaration
	var x, y int = 1, 2
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	// 6. Multiple variables with short declaration and inference
	e, f := true, "test"
	fmt.Println("e:", e)
	fmt.Println("f:", f)

	// 7. Reassignment (same type)
	e = false
	fmt.Println("e after reassignment:", e)

	// 8. Grouped variable declaration (var block) with zero values and explicit init
	var (
		name   string
		age    int
		active bool = true
	)
	fmt.Println("name (zero value):", name)
	fmt.Println("age (zero value):", age)
	fmt.Println("active:", active)

	// 9. Type printing for a few variables
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of c: %T\n", c)
	fmt.Printf("Type of d: %T\n", d)
	fmt.Printf("Type of e: %T\n", e)
	fmt.Printf("Type of name: %T\n", name)
}
