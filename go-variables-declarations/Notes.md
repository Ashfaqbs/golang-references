### Go — Variables & Declarations (Step 2: zero values, styles, reassignment)

---

## 1. New folder and file

Example on Windows:

```bash
mkdir C:\code\go-variables-declarations
cd C:\code\go-variables-declarations
```

Create `main.go` with the following content.

---

## 2. Program: different declarations + zero values, printed one by one

```go
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
```

Run:

```bash
go run main.go
```

---

## 3. Explanation of each concept (simple language)

### `package main`

```go
package main
```

* Declares the package name.
* `main` indicates that this package builds into an executable.
* The Go runtime will look for `func main()` inside this package and start execution there.

---

### `import "fmt"`

```go
import "fmt"
```

* Imports the standard library package `fmt`.
* `fmt` provides formatted I/O functions:

  * `Println` → print values with newline.
  * `Printf` → formatted printing.
* The code uses `fmt.Println` and `fmt.Printf`, so this import is required.
* Go enforces “no unused imports”; without any use of `fmt`, compilation would fail.

---

### `func main() { ... }`

```go
func main() {
	...
}
```

* Declares the `main` function.
* Signature:

  * No parameters.
  * No return type.
* This is the entry point for the program in a `main` package.

---

### 1. Explicit type with initialization

```go
var a int = 10
fmt.Println("a:", a)
```

* `var` → begins a variable declaration.
* `a` → variable name.
* `int` → type of the variable.
* `= 10` → initial value assigned.
* Type is explicit; no inference here.
* This style is useful when clarity of type is more important than brevity.

`fmt.Println("a:", a)` prints the label and the value: `a: 10`.

---

### 2. Explicit type without initialization (zero value)

```go
var b int
fmt.Println("b (zero value):", b)
```

* `var b int` declares `b` as an `int` but does not provide a value.
* In Go, every variable always has a value; if not explicitly set, it gets a **zero value**.

Zero values for basic types:

* `int`, `float*` → `0` or `0.0`
* `bool` → `false`
* `string` → `""` (empty string)
* Pointers, slices, maps, interfaces, channels, functions → `nil` (later topics)

Here:

* `b` is an `int`, so its zero value is `0`.
* `fmt.Println("b (zero value):", b)` prints `b (zero value): 0`.

Zero values make it safe to use variables immediately without worrying about “uninitialized” memory.

---

### 3. `var` with type inference

```go
var c = 3.14
fmt.Println("c:", c)
```

* `var c = 3.14`:

  * No explicit type is written.
  * The compiler infers the type from the value `3.14`.
  * Here, `c` becomes `float64` (default floating-point type).

This pattern keeps code slightly shorter while still being explicit that `c` is a package-local or function-local variable declared with `var`.

`fmt.Println("c:", c)` prints the value of `c`.

---

### 4. Short declaration with `:=`

```go
d := "Go"
fmt.Println("d:", d)
```

* `:=` is the short variable declaration syntax.
* Declares `d` and assigns `"Go"` to it in one step.
* Type is inferred from the right side:

  * `d` becomes `string`.
* Short declarations are only valid **inside functions**.

This is the most commonly used style in local scopes when the type is obvious from context.

---

### 5. Multiple variables of the same type

```go
var x, y int = 1, 2
fmt.Println("x:", x)
	fmt.Println("y:", y)
```

* Declares `x` and `y` in one line.
* Both have type `int`.
* `= 1, 2` assigns `1` to `x` and `2` to `y`.
* Useful when several related variables share the same type.

The prints show both values one by one.

---

### 6. Multiple variables with short declaration

```go
e, f := true, "test"
fmt.Println("e:", e)
fmt.Println("f:", f)
```

* Uses `:=` to declare **two** variables at once.
* `e` is inferred as `bool` (`true`).
* `f` is inferred as `string` (`"test"`).
* This style is idiomatic for returning multiple values from functions later.

---

### 7. Reassignment (same type)

```go
e = false
fmt.Println("e after reassignment:", e)
```

* `e` was already declared as `bool`.
* This line changes its value from `true` to `false`.
* Type must remain the same:

  * Assigning an `int` to `e` would cause a compile-time error (`cannot use 10 (untyped int constant) as type bool`).
* Go enforces type safety strictly at compile time.

---

### 8. Grouped variable declaration (var block)

```go
var (
	name   string
	age    int
	active bool = true
)
fmt.Println("name (zero value):", name)
fmt.Println("age (zero value):", age)
fmt.Println("active:", active)
```

* `var ( ... )` is a **var block**, grouping several variable declarations.
* `name string`:

  * No explicit value → zero value is `""` (empty string).
* `age int`:

  * No explicit value → zero value is `0`.
* `active bool = true`:

  * Explicit initial value `true`.

This pattern is common for grouping related configuration or state variables.

Printing:

* `name (zero value):` → empty string printed after the label.
* `age (zero value):` → `0`.
* `active:` → `true`.

---

### 9. Printing types

```go
fmt.Printf("Type of a: %T\n", a)
fmt.Printf("Type of c: %T\n", c)
fmt.Printf("Type of d: %T\n", d)
fmt.Printf("Type of e: %T\n", e)
fmt.Printf("Type of name: %T\n", name)
```

* `%T` is a format verb that prints the type of the value.
* `\n` is a newline character.
* These lines help confirm how the compiler inferred types:

  * `a` → `int`
  * `c` → `float64`
  * `d` → `string`
  * `e` → `bool`
  * `name` → `string`

---

## 4. Summary of core ideas from this step

* Three main declaration styles inside functions:

  1. `var x int = 10` → explicit type and value.
  2. `var x = 10` → type inferred, still using `var`.
  3. `x := 10` → short declaration, type inferred, local scope.
* Zero values:

  * `int` → `0`
  * `float64` → `0`
  * `bool` → `false`
  * `string` → `""`
* Reassignment must respect the original type.
* Multiple variables can be declared in a single statement.
* `var` blocks group related declarations neatly.
* `%T` with `fmt.Printf` inspects types at runtime, useful for learning and debugging.
