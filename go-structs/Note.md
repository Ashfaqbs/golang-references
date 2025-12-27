### Go — Structs (Step 7: custom types with fields)

---

## 1. New folder and file

Example on Windows (PowerShell or CMD):

```bash
mkdir C:\code\go-structs
cd C:\code\go-structs
```

Create `main.go` with the following content.

---

## 2. Program: define struct, create values, update, pointer, slice of structs

```go
package main

import "fmt"

// 1. Define a struct type
type Person struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	// 2. Zero value of a struct
	var p1 Person
	fmt.Println("p1 (zero value):", p1)
	fmt.Println("p1 fields -> Name:", p1.Name, "Age:", p1.Age, "Active:", p1.Active)

	// 3. Struct literal with field names
	p2 := Person{
		Name:   "Alice",
		Age:    30,
		Active: true,
	}
	fmt.Println("\np2 (with field names):", p2)

	// 4. Update struct fields
	p2.Age = 31
	p2.Active = false
	fmt.Println("p2 after updates:", p2)

	// 5. Struct literal without field names (positional)
	p3 := Person{"Bob", 25, true}
	fmt.Println("\np3 (positional literal):", p3)

	// 6. Pointer to struct and modification via pointer
	p4 := Person{
		Name:   "Charlie",
		Age:    40,
		Active: true,
	}
	fmt.Println("\np4 before pointer modification:", p4)

	p4Ptr := &p4
	p4Ptr.Age = 41
	p4Ptr.Active = false

	fmt.Println("p4 after pointer modification:", p4)
	fmt.Println("p4Ptr points to:", *p4Ptr)

	// 7. Slice of structs
	people := []Person{
		p2,
		p3,
		p4,
	}
	fmt.Println("\npeople slice:", people)

	// 8. Loop over slice of structs with range
	fmt.Println("\nLoop over people slice:")
	for index, person := range people {
		fmt.Println("index:", index, "Name:", person.Name, "Age:", person.Age, "Active:", person.Active)
	}
}
```

Run:

```bash
go run main.go
```

---

## 3. Explanation of every part (line by line, in simple terms)

---

### Package and import

```go
package main
```

* Declares that this file belongs to the `main` package.
* A `main` package plus a `main` function forms an executable program.

```go
import "fmt"
```

* Imports the standard library package `fmt`.
* `fmt` supplies formatted printing functions such as `Println`.
* Required because the code prints many intermediate values.
* Go requires imported packages to be used; unused imports cause compile-time errors.

---

### 1. Struct type definition

```go
// 1. Define a struct type
type Person struct {
	Name   string
	Age    int
	Active bool
}
```

* `type Person struct { ... }`:

  * Declares a new type named `Person`.
  * `struct` groups several fields together under one type.
* Fields:

  * `Name   string` → field `Name` of type `string`.
  * `Age    int` → field `Age` of type `int`.
  * `Active bool` → field `Active` of type `bool`.

Conceptually:

* `Person` is similar to a Java class that only has fields and no methods yet.
* Instances of `Person` will store related data as one value.

Capitalization:

* `Person` (type) and `Name`, `Age`, `Active` (fields) all start with a capital letter.
* In Go, identifiers starting with a capital letter are **exported** from the package (similar to `public` in Java).
* That detail becomes important when packages are separated later.

---

### 2. Zero value of a struct

```go
var p1 Person
fmt.Println("p1 (zero value):", p1)
fmt.Println("p1 fields -> Name:", p1.Name, "Age:", p1.Age, "Active:", p1.Active)
```

* `var p1 Person`:

  * Declares `p1` as a variable of type `Person`.
  * No explicit value is given.
* As with basic types, structs also receive a **zero value**:

  * Each field is set to its own zero value:

    * `Name` (`string`) → `""` (empty string)
    * `Age` (`int`) → `0`
    * `Active` (`bool`) → `false`
* Printing:

  * `p1` is printed as something like `{ 0 false}` (exact formatting may vary, but field order follows the struct definition).
  * The second `Println` call prints each field explicitly.

Zero values guarantee that a freshly declared struct variable is always in a valid, predictable state.

---

### 3. Struct literal with field names

```go
p2 := Person{
	Name:   "Alice",
	Age:    30,
	Active: true,
}
fmt.Println("\np2 (with field names):", p2)
```

* `Person{ ... }`:

  * Creates a new `Person` value.
* `Name: "Alice"`:

  * Sets the `Name` field to `"Alice"`.
* `Age: 30`:

  * Sets `Age` to `30`.
* `Active: true`:

  * Sets `Active` to `true`.

Key points:

* Field names are explicitly written, which makes the code self-documenting.
* Field order is flexible when using `fieldName: value` syntax.
* Very resilient to changes:

  * If the struct gains a new field later, the compiler will point out missing fields in literals, aiding maintenance.

`fmt.Println` prints the `Person` value as a struct, for example: `{Alice 30 true}`.

---

### 4. Updating struct fields

```go
p2.Age = 31
p2.Active = false
fmt.Println("p2 after updates:", p2)
```

* `p2.Age`:

  * Field selector syntax: `<variable>.<FieldName>`.
  * Accesses the `Age` field of `p2`.
* `p2.Age = 31`:

  * Assigns a new value to `Age`.
* `p2.Active = false`:

  * Changes `Active` to `false`.

Structs in Go use **value semantics**:

* `p2` is a value containing its fields.
* Assigning `p2` to another variable would copy all fields.

Here, no copy is made; only field updates on the same variable `p2`.

---

### 5. Struct literal without field names (positional)

```go
p3 := Person{"Bob", 25, true}
fmt.Println("\np3 (positional literal):", p3)
```

* This form uses the field order defined in the struct:

  * `"Bob"` → `Name`
  * `25` → `Age`
  * `true` → `Active`
* No field names are written, so the order must match the struct definition exactly.

Notes:

* This style is shorter but more fragile:

  * If the field order changes in the struct definition, this literal might silently change meaning or fail.
* For public APIs and long-lived code, field-name-based literals are usually preferred.

---

### 6. Pointer to struct and modification via pointer

```go
p4 := Person{
	Name:   "Charlie",
	Age:    40,
	Active: true,
}
fmt.Println("\np4 before pointer modification:", p4)

p4Ptr := &p4
p4Ptr.Age = 41
p4Ptr.Active = false

fmt.Println("p4 after pointer modification:", p4)
fmt.Println("p4Ptr points to:", *p4Ptr)
```

Breakdown:

* `p4 := Person{ ... }`:

  * Creates another `Person` value.
* `p4Ptr := &p4`:

  * `&` operator → “address of”.
  * `p4Ptr` is of type `*Person` (pointer to `Person`).
  * `p4Ptr` holds the memory address of `p4`.

Field access via pointer:

* `p4Ptr.Age = 41`:

  * Go automatically dereferences `p4Ptr` when accessing fields.
  * This line modifies the `Age` field of the original `p4`.
* `p4Ptr.Active = false`:

  * Same idea, sets `Active` to `false`.

After these updates:

* `p4` is changed (because pointer refers to the same underlying `Person`).
* `fmt.Println("p4 after pointer modification:", p4)` shows the updated struct.
* `*p4Ptr`:

  * `*` dereferences the pointer.
  * `fmt.Println("p4Ptr points to:", *p4Ptr)` prints the value stored at that address (same as `p4`).

Why this is important:

* Structs have value semantics; copying a struct copies its fields.
* Pointers allow shared, mutable access to the same struct instance.
* Methods often use pointer receivers (`func (p *Person) ...`) to modify struct state; this pattern is a foundation for that later topic.

---

### 7. Slice of structs

```go
people := []Person{
	p2,
	p3,
	p4,
}
fmt.Println("\npeople slice:", people)
```

* `[]Person{ ... }`:

  * Slice literal where each element is a `Person`.
* Elements:

  * `p2` (`Person`).
  * `p3` (`Person`).
  * `p4` (`Person`).
* `people` is of type `[]Person` (slice of `Person`).

Conceptually:

* Combines the idea of slices from earlier steps with custom struct types.
* Very similar to `List<Person>` in Java, but syntax and semantics are native to Go.

`fmt.Println("people slice:", people)` prints the slice as `[ {…} {…} {…} ]` with each struct’s content inline.

---

### 8. Loop over slice of structs with `range`

```go
fmt.Println("\nLoop over people slice:")
	for index, person := range people {
	fmt.Println("index:", index, "Name:", person.Name, "Age:", person.Age, "Active:", person.Active)
}
```

* `range people`:

  * Iterates over the slice.
  * Provides:

    * Index (position).
    * Element value at that position (a `Person`).
* `index`:

  * `int` index into the `people` slice.
* `person`:

  * Copy of each `Person` value in the slice.

Each iteration prints:

* Index (0-based).
* `Name`, `Age`, `Active` fields of the `Person`.

Pattern:

```go
for i, v := range sliceOfStructs {
	// work with v.FieldName
}
```

is a common way to traverse a collection of rich data structures.

---

## 4. Summary of core ideas from this step

* `struct` introduces custom composite types with named fields:

  ```go
  type Person struct {
      Name string
      Age  int
  }
  ```
* Zero value of a struct:

  * Each field has its own zero value.
* Initialization styles:

  * With field names (recommended):

    ```go
    Person{Name: "Alice", Age: 30}
    ```
  * Positional (shorter, but order-dependent):

    ```go
    Person{"Bob", 25, true}
    ```
* Field access and mutation:

  * `p.Field` for reading and writing.
* Pointers to structs:

  * `&p` gets a pointer.
  * `ptr.Field` automatically dereferences for field access.
  * Allows mutating the original struct through the pointer.
* Slices of structs:

  * `[]Person{...}` models collections of complex objects.
  * `range` over `[]Person` combines loop semantics with struct field access.
