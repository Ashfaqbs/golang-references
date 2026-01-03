package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 1. Nil vs empty slice
	var s1 []int // not defining size {} so nil
	s2 := []int{}
	s3 := make([]int, 0)

	fmt.Println("== Slices ==")
	fmt.Println("s1 (nil slice):", s1, "len:", len(s1))
	fmt.Println("s2 (empty literal):", s2, "len:", len(s2))
	fmt.Println("s3 (make with length 0):", s3, "len:", len(s3))

	fmt.Println("s1 == nil:", s1 == nil)
	fmt.Println("s2 == nil:", s2 == nil)
	fmt.Println("s3 == nil:", s3 == nil)

	// Append to nil slice (works, returns new slice)
	s1 = append(s1, 10, 20)
	fmt.Println("s1 after append:", s1, "len:", len(s1))

	// Range over slices (nil and empty)
	fmt.Println("\nRange over s1:")
	for i, v := range s1 {
		fmt.Println("index:", i, "value:", v)
	}

	fmt.Println("\nRange over s2 (empty):")
	for i, v := range s2 {
		fmt.Println("index:", i, "value:", v)
	}
	// no output from this loop

	// 2. Nil vs initialized map
	var m1 map[string]int
	m2 := map[string]int{}
	m3 := make(map[string]int)

	fmt.Println("\n== Maps ==")
	fmt.Println("m1 (nil map):", m1)
	fmt.Println("m2 (empty literal):", m2)
	fmt.Println("m3 (make):", m3)

	fmt.Println("m1 == nil:", m1 == nil)
	fmt.Println("m2 == nil:", m2 == nil)
	fmt.Println("m3 == nil:", m3 == nil)

	// Reading from nil map (safe, returns zero value)
	fmt.Println(`m1["x"] from nil map:`, m1["x"])

	// Writing to nil map (panics)
	fmt.Println("\nAttempting write to nil map (m1)...")
	// Uncommenting the next line will cause a runtime panic:
	// m1["x"] = 1

	// Writing to initialized maps (ok)
	m2["a"] = 10
	m3["b"] = 20
	fmt.Println("m2 after write:", m2)
	fmt.Println("m3 after write:", m3)

	// 3. Pointer basics and mutation
	fmt.Println("\n== Pointers and mutation ==")
	x := 10
	y := x
	y = 20

	fmt.Println("x (original scalar):", x)
	fmt.Println("y (copy of x):", y)

	z := 10
	zPtr := &z
	fmt.Println("z before:", z, "zPtr:", zPtr, "*zPtr:", *zPtr)

	*zPtr = 99
	fmt.Println("z after *zPtr = 99:", z, "zPtr:", zPtr, "*zPtr:", *zPtr)

	// 4. Pointers to struct and shared mutation
	fmt.Println("\n== Struct pointer sharing ==")
	p := Person{Name: "Alice", Age: 30}
	pCopy := p
	pPtr := &p

	pCopy.Age = 40
	pPtr.Age = 35

	fmt.Println("p      (original):", p)
	fmt.Println("pCopy  (value copy):", pCopy)
	fmt.Println("pPtr-> (via pointer):", *pPtr)

	// 5. Nil pointer
	var pNil *Person
	fmt.Println("\n== Nil pointer ==")
	fmt.Println("pNil:", pNil)
	fmt.Println("pNil == nil:", pNil == nil)

	// Accessing fields via nil pointer would panic:
	// fmt.Println(pNil.Name) // runtime panic: dereference of nil pointer

	// 6. Nil interface vs non-nil interface holding nil pointer
	fmt.Println("\n== Interfaces and nil ==")

	var i1 interface{}
	fmt.Println("i1 (uninitialized interface):", i1)
	fmt.Println("i1 == nil:", i1 == nil)

	var pNil2 *Person = nil
	var i2 interface{} = pNil2
	fmt.Println("i2 (interface holding nil *Person):", i2)
	fmt.Println("i2 == nil:", i2 == nil)

	// Type assertion from interface to *Person
	if pFromI2, ok := i2.(*Person); ok {
		fmt.Println("Type assertion i2.(*Person) succeeded, value:", pFromI2)
		fmt.Println("pFromI2 == nil:", pFromI2 == nil)
	}
}
