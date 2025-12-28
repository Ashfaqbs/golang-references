### Go — Functions and Methods (Step 8: reusable logic + behavior on structs)

---

## 1. New folder and file

Example on Windows:

```bash
mkdir C:\code\go-functions-methods
cd C:\code\go-functions-methods
```

Create `main.go` with the following content.

---

## 2. Program: functions, multiple returns, methods on struct

```go
package main

import "fmt"

// 1. Struct type reused from previous step
type Person struct {
	Name string
	Age  int
}

// 2. Simple function with one parameter, no return value
func greet(name string) {
	fmt.Println("Hello,", name)
}

// 3. Function with parameters and a single return value
func add(a int, b int) int {
	sum := a + b
	return sum
}

// 4. Function with multiple return values
func divideAndRemainder(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// 5. Method with value receiver (does not modify original)
func (p Person) Info() {
	fmt.Println("Person Info -> Name:", p.Name, "Age:", p.Age)
}

// 6. Method with pointer receiver (can modify original)
func (p *Person) HaveBirthday() {
	p.Age = p.Age + 1
	fmt.Println("Happy birthday,", p.Name, "New age:", p.Age)
}

func main() {
	// Call simple function
	fmt.Println("Calling greet:")
	greet("Go learner")

	// Call function with return value
	fmt.Println("\nCalling add:")
	result := add(10, 20)
	fmt.Println("add(10, 20):", result)

	// Call function with multiple return values
	fmt.Println("\nCalling divideAndRemainder:")
	q, r := divideAndRemainder(17, 5)
	fmt.Println("divideAndRemainder(17, 5) -> quotient:", q, "remainder:", r)

	// Create a Person value
	fmt.Println("\nCreating Person:")
	person := Person{
		Name: "Alice",
		Age:  29,
	}
	fmt.Println("Initial person:", person)

	// Call method with value receiver
	fmt.Println("\nCalling person.Info():")
	person.Info()

	// Call method with pointer receiver (modifies Age)
	fmt.Println("\nCalling person.HaveBirthday():")
	person.HaveBirthday()
	fmt.Println("Person after HaveBirthday:", person)

	// Explicit pointer call (same effect)
	fmt.Println("\nCalling HaveBirthday via pointer:")
	personPtr := &person
	personPtr.HaveBirthday()
	fmt.Println("Person after pointer HaveBirthday:", person)
}
```

Run:

```bash
go run main.go
```

---

## 3. Explanation of each part

---

### Package and import

```go
package main
```

* File belongs to the `main` package.
* A `main` package with a `main` function compiles into an executable program.

```go
import "fmt"
```

* Imports the `fmt` standard library package.
* Provides `Println` for printing.
* Required because many values and messages are printed.

---

### 1. Struct type

```go
type Person struct {
	Name string
	Age  int
}
```

* `type Person struct { ... }`:

  * Declares a custom type named `Person`.
  * `Name` and `Age` are fields on this struct.
* Same idea as in previous step, used here to attach methods and show behavior.

---

### 2. Simple function: one parameter, no return

```go
func greet(name string) {
	fmt.Println("Hello,", name)
}
```

Breakdown:

* `func`:

  * Keyword that begins a function declaration.
* `greet`:

  * Function name.
* `(name string)`:

  * Function parameter list.
  * One parameter:

    * `name` → parameter name.
    * `string` → type of that parameter.
  * Parameter type is written **after** the name in Go.
* Return type:

  * No return type is written after the parameter list.
  * That means the function returns nothing.
* Body:

  ```go
  {
      fmt.Println("Hello,", name)
  }
  ```

  * Prints `"Hello,"` followed by the `name` value.

This function encapsulates a single operation: printing a greeting.

---

### 3. Function with parameters and a single return value

```go
func add(a int, b int) int {
	sum := a + b
	return sum
}
```

Breakdown:

* `func add(a int, b int) int`:

  * `add` is the function name.
  * Parameters:

    * `a int`
    * `b int`
  * Return type: `int`.
* Inside:

  ```go
  sum := a + b
  return sum
  ```

  * `sum` is a local variable storing the result.
  * `return sum` sends the value back to the caller.

In Go, the return type is written **after** the parameter list.
Single return value → just one type after the `)`.

Call site in `main`:

```go
result := add(10, 20)
fmt.Println("add(10, 20):", result)
```

* `add(10, 20)` calls the function with arguments `10` and `20`.
* Result is assigned to `result` and printed.

---

### 4. Function with multiple return values

```go
func divideAndRemainder(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}
```

* Parameter list:

  * `a int`
  * `b int`
* Return types:

  * `(int, int)` → two `int` values.
* Inside:

  * `quotient`:

    * Result of integer division `a / b`.
  * `remainder`:

    * Result of `a % b`.

`return quotient, remainder`:

* Returns two values in order.
* Caller receives both values.

Call site:

```go
q, r := divideAndRemainder(17, 5)
fmt.Println("divideAndRemainder(17, 5) -> quotient:", q, "remainder:", r)
```

* `q` will hold `3`.
* `r` will hold `2`.
* Pattern `(value1, value2 := functionCall())` is very common in Go.

Go’s multiple return feature is heavily used for results paired with error values in more advanced code.

---

### 5. Method with value receiver

```go
func (p Person) Info() {
	fmt.Println("Person Info -> Name:", p.Name, "Age:", p.Age)
}
```

Key ideas:

* `func (p Person) Info()`:

  * `(p Person)` is the **receiver**.
  * `Info` is a method attached to the `Person` type.
  * This method can be called with the dot syntax: `person.Info()`.

* Receiver semantics:

  * `p` is a **copy** of the `Person` value on which the method is called.
  * Changes to `p` inside this method do not modify the original struct.

* Inside body:

  * Accesses fields via `p.Name` and `p.Age`.
  * Simply prints information; no mutation.

Methods provide behavior associated with a type, in the same spirit as methods on Java classes, but the association is specified explicitly via the receiver.

Call site:

```go
person.Info()
```

---

### 6. Method with pointer receiver

```go
func (p *Person) HaveBirthday() {
	p.Age = p.Age + 1
	fmt.Println("Happy birthday,", p.Name, "New age:", p.Age)
}
```

Differences from `Info`:

* Receiver: `(p *Person)`:

  * Receiver is a pointer to `Person`.
  * `p` holds an address pointing to an actual `Person` value.
* `p.Age = p.Age + 1`:

  * Increments the `Age` field of the underlying `Person`.
  * Because `p` is a pointer, the original `Person` is modified.

Automatic dereferencing:

* Even though `p` is a pointer, field access uses `p.Age`, not `(*p).Age`.
* Go automatically dereferences the pointer when accessing fields.

Pointer receivers are used when:

* Methods need to modify the receiver’s state.
* Avoiding copying large structs on every method call is important.

Call site:

```go
person.HaveBirthday()
```

Even though `person` is a value, Go automatically takes its address when calling a method with a pointer receiver.
This is syntactic sugar for `(&person).HaveBirthday()`.

---

### 7. `main` function

```go
func main() {
	// Call simple function
	fmt.Println("Calling greet:")
	greet("Go learner")
```

* `func main()`:

  * Entry point of the Go program.
  * No parameters, no return type.

First part:

* Prints a label.
* Calls `greet("Go learner")`.
* `greet` prints a greeting message.

---

```go
	fmt.Println("\nCalling add:")
	result := add(10, 20)
	fmt.Println("add(10, 20):", result)
```

* Prints label.
* Calls `add(10, 20)`.
* Assigns returned value to `result`.
* Prints the value.

---

```go
	fmt.Println("\nCalling divideAndRemainder:")
	q, r := divideAndRemainder(17, 5)
	fmt.Println("divideAndRemainder(17, 5) -> quotient:", q, "remainder:", r)
```

* Calls `divideAndRemainder(17, 5)`.
* Uses multiple assignment `q, r := ...` to capture both return values.
* Prints both.

---

```go
	// Create a Person value
	fmt.Println("\nCreating Person:")
	person := Person{
		Name: "Alice",
		Age:  29,
	}
	fmt.Println("Initial person:", person)
```

* Creates a `Person` with a struct literal using field names.
* Prints the initial struct state.

---

```go
	// Call method with value receiver
	fmt.Println("\nCalling person.Info():")
	person.Info()
```

* Calls `Info` on `person`.
* Uses value receiver.
* Only reads and prints fields, no mutation.

---

```go
	// Call method with pointer receiver (modifies Age)
	fmt.Println("\nCalling person.HaveBirthday():")
	person.HaveBirthday()
	fmt.Println("Person after HaveBirthday:", person)
```

* Calls `HaveBirthday` on `person`.
* Pointer receiver increments `Age`.
* Prints updated `person`.

---

```go
	// Explicit pointer call (same effect)
	fmt.Println("\nCalling HaveBirthday via pointer:")
	personPtr := &person
	personPtr.HaveBirthday()
	fmt.Println("Person after pointer HaveBirthday:", person)
}
```

* `personPtr := &person`:

  * Takes address of `person`.
  * Type: `*Person`.
* `personPtr.HaveBirthday()`:

  * Calls method via pointer explicitly.
  * Same effect as `person.HaveBirthday()`.
* Final print shows `person` with `Age` incremented again.

---

## 4. Summary of core ideas from this step

* Functions:

  * Declared with `func` at package level.
  * Parameters: `name type`.
  * Return types are written after the parameter list.
  * Multiple return values are a first-class feature:

    ```go
    func f() (int, string) { ... }
    ```

* Methods:

  * Functions attached to a type using a receiver:

    ```go
    func (r ReceiverType) MethodName(...) ...
    ```
  * Value receiver:

    * `(p Person)` → method receives a copy.
    * Does not modify original value (unless fields are reference types inside).
  * Pointer receiver:

    * `(p *Person)` → method receives a pointer.
    * Can modify the original value.
    * Avoids copying large structs.

* Method calls:

  * `value.Method()` and `pointer.Method()` both work.
  * Go adds automatic address/dereference where appropriate.

* Structs + methods:

  * Form the basic building block for modeling entities and their behavior in Go.
  * Similar conceptual role to classes with methods in Java, but composition and interfaces are preferred instead of inheritance.