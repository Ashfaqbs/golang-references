### Go — `nil`, Slices, Maps, Pointers, Interfaces (Step 12: reference behavior)

---

## 1. New folder and file

Example on Windows:

```bash
mkdir C:\code\go-nil-and-references
cd C:\code\go-nil-and-references
```

Create `main.go` with the following content.

---

## 2. Program: `nil` slices, `nil` maps, pointers, `nil` interfaces

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 1. Nil vs empty slice
	var s1 []int
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
```

---

## 3. Explanation of each concept

---

### 1. Nil vs empty slices

```go
var s1 []int
s2 := []int{}
s3 := make([]int, 0)
```

* `var s1 []int`

  * Declares a slice variable with type `[]int`.
  * Zero value for a slice is `nil`.
* `s2 := []int{}`

  * Slice literal with zero elements.
  * Non-nil slice whose length is `0`.
* `s3 := make([]int, 0)`

  * Uses `make` to create a slice with length `0`.
  * Non-nil slice.

```go
fmt.Println("s1 (nil slice):", s1, "len:", len(s1))
fmt.Println("s2 (empty literal):", s2, "len:", len(s2))
fmt.Println("s3 (make with length 0):", s3, "len:", len(s3))

fmt.Println("s1 == nil:", s1 == nil)
fmt.Println("s2 == nil:", s2 == nil)
fmt.Println("s3 == nil:", s3 == nil)
```

* All have `len == 0`.
* Only `s1` is actually `nil`.
* `s2` and `s3` are distinct (non-nil) empty slices.

Important slice facts:

* `len(nilSlice)` is `0`.
* `range` over a nil slice does nothing (no iterations).
* `append` on a nil slice works and returns a new non-nil slice.

```go
s1 = append(s1, 10, 20)
fmt.Println("s1 after append:", s1, "len:", len(s1))
```

* Appending to a nil slice:

  * Allocates underlying array.
  * Returns a valid slice containing the appended elements.

Range examples:

```go
fmt.Println("\nRange over s1:")
for i, v := range s1 {
	fmt.Println("index:", i, "value:", v)
}

fmt.Println("\nRange over s2 (empty):")
for i, v := range s2 {
	fmt.Println("index:", i, "value:", v)
}
```

* Ranging over a non-empty slice (`s1`) iterates over each element.
* Ranging over an empty slice (`s2`) runs zero times; body does not execute.

---

### 2. Nil vs initialized maps

```go
var m1 map[string]int
m2 := map[string]int{}
m3 := make(map[string]int)
```

* `var m1 map[string]int`

  * Declares a map variable.
  * Zero value of a map is `nil`.
* `m2 := map[string]int{}`

  * Map literal, initialized as empty but non-nil.
* `m3 := make(map[string]int)`

  * `make` creates an empty but ready-to-use map.

```go
fmt.Println("\n== Maps ==")
fmt.Println("m1 (nil map):", m1)
fmt.Println("m2 (empty literal):", m2)
fmt.Println("m3 (make):", m3)

fmt.Println("m1 == nil:", m1 == nil)
fmt.Println("m2 == nil:", m2 == nil)
fmt.Println("m3 == nil:", m3 == nil)
```

* `m1` is `nil`.
* `m2`, `m3` are non-nil maps, even though currently empty.

Map read from nil map:

```go
fmt.Println(`m1["x"] from nil map:`, m1["x"])
```

* Reading from a `nil` map:

  * Safe.
  * Returns zero value for the map’s value type.
  * Here, zero value for `int` is `0`.

Map write to nil map:

```go
// m1["x"] = 1
```

* Writing to a `nil` map causes a runtime panic.
* Comment left to indicate this behavior without breaking the program.

Writing to initialized maps:

```go
m2["a"] = 10
m3["b"] = 20
fmt.Println("m2 after write:", m2)
fmt.Println("m3 after write:", m3)
```

* Valid; maps created via literal or `make` are ready to accept writes.

Summary:

* Slices: `nil` slice is safe to read length, range, and append.
* Maps: `nil` map is safe to read, but **not** safe to write.

---

### 3. Pointers and scalar mutation

```go
x := 10
y := x
y = 20

fmt.Println("x (original scalar):", x)
fmt.Println("y (copy of x):", y)
```

* `y := x` copies the value `10`.
* Updating `y` does not affect `x`.
* Demonstrates value semantics for basic types.

```go
z := 10
zPtr := &z
fmt.Println("z before:", z, "zPtr:", zPtr, "*zPtr:", *zPtr)

*zPtr = 99
fmt.Println("z after *zPtr = 99:", z, "zPtr:", zPtr, "*zPtr:", *zPtr)
```

* `&z`:

  * Address-of operator.
  * `zPtr` is of type `*int`.
* `*zPtr`:

  * Dereferences `zPtr` to access the `int` value it points to.
* Assigning through the pointer (`*zPtr = 99`) updates the original `z`.

This shows:

* Value copy vs reference via pointer.
* Pointer enables shared, mutable state over a variable.

---

### 4. Struct pointer sharing

```go
p := Person{Name: "Alice", Age: 30}
pCopy := p
pPtr := &p

pCopy.Age = 40
pPtr.Age = 35

fmt.Println("p      (original):", p)
fmt.Println("pCopy  (value copy):", pCopy)
fmt.Println("pPtr-> (via pointer):", *pPtr)
```

* `p`:

  * Original struct.
* `pCopy := p`:

  * Value copy of the struct.
  * Changing `pCopy` does not affect `p`.
* `pPtr := &p`:

  * Pointer to original `p`.
* `pCopy.Age = 40`:

  * Changes only the copy.
* `pPtr.Age = 35`:

  * Through pointer, changes original `p`.
  * Go automatically dereferences pointer in field selectors.

Result:

* `p`’s `Age` becomes `35`.
* `pCopy`’s `Age` becomes `40`.
* `*pPtr` prints same content as `p`.

---

### 5. Nil pointer

```go
var pNil *Person
fmt.Println("\n== Nil pointer ==")
fmt.Println("pNil:", pNil)
fmt.Println("pNil == nil:", pNil == nil)
```

* `var pNil *Person`:

  * Declares a pointer to `Person`.
  * Zero value is `nil`.
* `pNil == nil`:

  * `true`.

Dereferencing or accessing fields via `pNil`:

```go
// fmt.Println(pNil.Name)
```

* Would cause a runtime panic:

  * `panic: runtime error: invalid memory address or nil pointer dereference`.

So pointer variables should either be checked for `nil` or guaranteed non-nil before use.

---

### 6. Interfaces and `nil`

```go
var i1 interface{}
fmt.Println("i1 (uninitialized interface):", i1)
fmt.Println("i1 == nil:", i1 == nil)
```

* `var i1 interface{}`:

  * Declares a variable of the empty interface type.
  * Zero value is an interface with:

    * No dynamic type.
    * No dynamic value.
* `i1 == nil`:

  * `true`.

Now, interface holding a nil pointer:

```go
var pNil2 *Person = nil
var i2 interface{} = pNil2
fmt.Println("i2 (interface holding nil *Person):", i2)
fmt.Println("i2 == nil:", i2 == nil)
```

* `pNil2`:

  * Nil pointer to `Person`.
* `i2`:

  * Interface value whose **dynamic type** is `*Person` and **dynamic value** is `nil`.
* `i2 == nil`:

  * `false`, because the interface is not empty:

    * It has a type (`*Person`) even though the underlying value is `nil`.

This is a subtle but important point:

* An interface is `nil` only if both its dynamic type and dynamic value are `nil`.
* An interface holding a typed `nil` pointer is **not** itself `nil`.

Type assertion:

```go
if pFromI2, ok := i2.(*Person); ok {
	fmt.Println("Type assertion i2.(*Person) succeeded, value:", pFromI2)
	fmt.Println("pFromI2 == nil:", pFromI2 == nil)
}
```

* Uses type assertion `i2.(*Person)`:

  * Asks the interface whether its dynamic type is `*Person`.
* `ok`:

  * `true` if assertion succeeded.
* `pFromI2`:

  * Resulting `*Person`, which is `nil` in this case.
* `pFromI2 == nil`:

  * `true`.

This pattern is common when inspecting underlying values stored in interfaces.

---

## 4. Summary of core ideas from this step

* Slices:

  * Zero value: `nil`, but safe for `len`, `range`, and `append`.
  * Empty but non-nil slices can be created with `[]T{}` or `make([]T, 0)`.
* Maps:

  * Zero value: `nil`, safe for reads, **not** safe for writes.
  * Must be created via literal or `make` before writing.
* Pointers:

  * `&x` gets the address of a variable.
  * `*ptr` dereferences the pointer.
  * Enable shared mutation of the same value.
* Struct values vs copies vs pointers:

  * Direct assignment copies the struct.
  * Pointer points to a single shared instance.
* Nil pointers:

  * Zero value of a pointer type.
  * Must not be dereferenced or used for field access without checks.
* Interfaces and nil:

  * Plain `var i interface{}` is `nil`.
  * Interface holding a typed nil pointer is **not** `nil.
  * Correct handling sometimes requires type assertions and explicit checks.