### Go — Interfaces (Step 9: behavior contracts)

---

## 1. New folder and file

Example on Windows (PowerShell or CMD):

```bash
mkdir C:\code\go-interfaces
cd C:\code\go-interfaces
```

Create `main.go` with the following content.

---

## 2. Program: interface, multiple implementations, function using interface, slice of interface, nil interface

```go
package main

import "fmt"

// 1. Interface definition (behavior contract)
type Describer interface {
	Describe() string
}

// 2. Concrete type 1
type Person struct {
	Name string
	Age  int
}

// Person implements Describer with a value receiver method
func (p Person) Describe() string {
	return fmt.Sprintf("Person(Name=%s, Age=%d)", p.Name, p.Age)
}

// 3. Concrete type 2
type Product struct {
	Name  string
	Price float64
}

// Product implements Describer with a value receiver method
func (p Product) Describe() string {
	return fmt.Sprintf("Product(Name=%s, Price=%.2f)", p.Name, p.Price)
}

// 4. Function that works with any Describer
func PrintDescription(d Describer) {
	fmt.Println("Description:", d.Describe())
}

func main() {
	// 5. Create concrete values
	person := Person{
		Name: "Alice",
		Age:  30,
	}
	product := Product{
		Name:  "Laptop",
		Price: 1299.99,
	}

	fmt.Println("Concrete values:")
	fmt.Println("person:", person)
	fmt.Println("product:", product)

	// 6. Assign concrete values to interface variable
	var d Describer

	d = person
	fmt.Println("\nDescriber holding Person:")
	PrintDescription(d)

	d = product
	fmt.Println("\nDescriber holding Product:")
	PrintDescription(d)

	// 7. Slice of Describer (polymorphic collection)
	items := []Describer{
		person,
		product,
		Product{Name: "Phone", Price: 699.50},
		Person{Name: "Bob", Age: 25},
	}

	fmt.Println("\nLoop over slice of Describer:")
	for index, item := range items {
		fmt.Println("index:", index)
		PrintDescription(item)
	}

	// 8. Zero value (nil) of an interface
	var empty Describer
	fmt.Println("\nZero value of interface (empty):", empty)
	fmt.Println("Is empty == nil?", empty == nil)
}
```

Run:

```bash
go run main.go
```

---

## 3. Explanation of every part (simple language)

---

### Package and import

```go
package main
```

* Declares that this file belongs to the `main` package.
* A `main` package with a `main` function builds into an executable binary.

```go
import "fmt"
```

* Imports the `fmt` package from the standard library.
* Provides functions such as `Println` and `Sprintf`.
* Required in this file for printing and string formatting.

Unused imports are not allowed; all imported packages must be referenced somewhere in the code.

---

### 1. Interface definition (behavior contract)

```go
type Describer interface {
	Describe() string
}
```

* `type Describer interface { ... }`:

  * Declares a new interface type named `Describer`.
* Inside the interface:

  ```go
  Describe() string
  ```

  * Method signature.
  * Any type that has a method with this exact signature:

    * Name: `Describe`
    * Parameters: none
    * Return type: `string`
  * …is said to **implement** the `Describer` interface.

Important details:

* There is **no** `implements` keyword in Go.
* Implementation is **implicit**:

  * If a type has all methods listed in the interface, it implements that interface.
  * No extra declaration is needed on the type.

Interfaces in Go represent **behavior**, not data.
`Describer` means “anything that can describe itself as a string.”

---

### 2. Concrete type 1: `Person`

```go
type Person struct {
	Name string
	Age  int
}
```

* Struct definition as before.
* `Person` has two fields: `Name` and `Age`.

Method:

```go
func (p Person) Describe() string {
	return fmt.Sprintf("Person(Name=%s, Age=%d)", p.Name, p.Age)
}
```

* `func (p Person) Describe() string`:

  * Method attached to `Person`.
  * Receiver: `(p Person)` → value receiver (copy of `Person`).
  * Method name: `Describe`.
  * Return type: `string`.
* Body:

  * Uses `fmt.Sprintf`:

    * `Sprintf` formats a string and returns it (no printing to console).
    * Format string: `"Person(Name=%s, Age=%d)"`

      * `%s` → placeholder for string.
      * `%d` → placeholder for integer.
  * `p.Name` and `p.Age` are passed into `Sprintf`.
  * Returned string describes the person.

Since `Person` has a `Describe() string` method, `Person` **implements** `Describer` by the rules of the interface.

---

### 3. Concrete type 2: `Product`

```go
type Product struct {
	Name  string
	Price float64
}
```

* Another struct type.
* Fields:

  * `Name` (string).
  * `Price` (float64).

Method:

```go
func (p Product) Describe() string {
	return fmt.Sprintf("Product(Name=%s, Price=%.2f)", p.Name, p.Price)
}
```

* Value receiver: `(p Product)`.
* Returns a formatted string with the product’s name and price.
* `%.2f`:

  * Floating point format with 2 decimal places.

`Product` also has a `Describe() string` method, so `Product` also implements `Describer`.

Now there are two different concrete types that both satisfy the same interface.

---

### 4. Function that works with any `Describer`

```go
func PrintDescription(d Describer) {
	fmt.Println("Description:", d.Describe())
}
```

* Parameter: `d Describer`.

  * `d` can hold **any value of any type** that implements `Describer`.
* Inside:

  * Calls `d.Describe()`.
  * Prints `"Description:"` followed by the returned string.

This function **does not care** whether `d` is a `Person`, `Product`, or any other type, as long as it has the required method.
That is the core polymorphism mechanism in Go.

---

### 5. Create concrete values

```go
person := Person{
	Name: "Alice",
	Age:  30,
}
product := Product{
	Name:  "Laptop",
	Price: 1299.99,
}

fmt.Println("Concrete values:")
fmt.Println("person:", person)
fmt.Println("product:", product)
```

* `person`:

  * Struct literal with field names.
* `product`:

  * Struct literal with field names and float price.
* Direct prints of `person` and `product`:

  * Use the default `fmt.Println` formatting for structs:

    * Example: `{Alice 30}` or `{Laptop 1299.99}`.

At this stage, these are just plain struct values.

---

### 6. Assign concrete values to interface variable

```go
var d Describer
```

* Declares `d` of type `Describer`.
* Zero value of an interface variable is `nil`:

  * Before assignment, `d` holds no concrete value.
  * Later printed explicitly for demonstration.

First assignment:

```go
d = person
fmt.Println("\nDescriber holding Person:")
PrintDescription(d)
```

* `person` (type `Person`) implements `Describer`.
* Assignment is valid: a `Person` value fits into a `Describer` variable.
* After assignment:

  * `d` internally holds:

    * Dynamic type: `Person`
    * Dynamic value: the `Person` instance `{Name: "Alice", Age: 30}`
* `PrintDescription(d)`:

  * Calls `d.Describe()` → invokes `Person.Describe` implementation.

Second assignment:

```go
d = product
fmt.Println("\nDescriber holding Product:")
PrintDescription(d)
```

* Now `d` holds:

  * Dynamic type: `Product`
  * Dynamic value: the `Product` instance.
* `PrintDescription(d)` calls `Product.Describe`.

The same function `PrintDescription` works with different concrete types by relying entirely on the interface contract.

---

### 7. Slice of `Describer` (polymorphic collection)

```go
items := []Describer{
	person,
	product,
	Product{Name: "Phone", Price: 699.50},
	Person{Name: "Bob", Age: 25},
}
```

* `[]Describer{ ... }`:

  * Slice where each element is of type `Describer`.
* Elements:

  * Existing variables: `person`, `product`.
  * New literals: a `Product` and a `Person` created inline.
* Each element must have a type that implements `Describer`.

Loop:

```go
fmt.Println("\nLoop over slice of Describer:")
for index, item := range items {
	fmt.Println("index:", index)
	PrintDescription(item)
}
```

* `range items`:

  * Iterates over each `Describer` in the slice.
* For each `item`:

  * Concrete dynamic type may differ (`Person` or `Product`).
  * `PrintDescription(item)` prints the correct description based on the actual type.

This is the idiomatic Go equivalent of a polymorphic collection in OOP:
a slice of interface type where each element is a different concrete type.

---

### 8. Zero value of an interface (nil)

```go
var empty Describer
fmt.Println("\nZero value of interface (empty):", empty)
fmt.Println("Is empty == nil?", empty == nil)
```

* `var empty Describer`:

  * Declares `empty` of type `Describer`.
  * No assignment is made.
* Zero value of an interface:

  * Both its dynamic type and dynamic value are `nil`.
* Printing:

  * First print typically shows `<nil>` or nothing for the value.
  * Second print confirms that `empty == nil` is `true`.

This is important later when checking for non-initialized interfaces or error-like patterns.

---

## 4. Summary of core interface concepts from this step

* Interface definition:

  ```go
  type InterfaceName interface {
      Method1(...)
      Method2(...) ReturnType
  }
  ```

  * Defines a **behavior contract** (set of methods).
* Implementation is implicit:

  * Any type that has all listed methods implements the interface.
  * No `implements` keyword.
* Interface variables:

  * Can hold values of any type that implements the interface.
  * Support polymorphism: same function can work with many concrete types.
* Functions with interface parameters:

  * Decoupled from specific concrete types.
  * Depend only on required behavior (methods).
* Slices of interfaces:

  * Collections containing heterogeneous values that share a common behavior.
* Zero value of an interface:

  * `nil` → no dynamic type, no value.
