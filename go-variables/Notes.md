### Go — Variables & Types (Step 1)

---

## 1. New folder and file

Example structure on Windows (PowerShell or CMD):

```bash
mkdir C:\code\go-variables
cd C:\code\go-variables
```

Create a file named `main.go` in that folder.

---

## 2. Program: basic variables and types (printed one by one)

```go
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
	count := 10          // inferred as int
	ratio := 0.75        // inferred as float64
	title := "Go notes"  // inferred as string

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
```

Run:

```bash
go run main.go
```

Output will show each variable’s value, then each variable’s type.

---

## 3. Explanation of concepts (line by line, in simple terms)

### Package and imports

```go
package main
```

* Declares that this file belongs to the `main` package.
* `package main` + `func main()` means “this is an executable program”.
* Without `package main`, the file would be a library, not a directly runnable program.

```go
import "fmt"
```

* Imports the `fmt` package from Go’s standard library.
* `fmt` provides formatting and printing functions such as `Println` and `Printf`.
* The imported package is used later as `fmt.Println` and `fmt.Printf`.
* Go does not allow unused imports; if `fmt` was not used anywhere, compilation would fail with an error.

---

### The `main` function

```go
func main() {
    ...
}
```

* `func` → keyword to declare a function.
* `main` → function name; in `package main` this is the entry point of the program.
* `()` → empty parameter list (no arguments).
* `{ ... }` → the body of the function.
* The program starts running from `main`, similar to Java’s `public static void main`.

---

### Booleans

```go
var isGoFun bool = true
fmt.Println("isGoFun:", isGoFun)
```

* `var` → keyword to declare a variable.
* `isGoFun` → variable name.
* `bool` → type; can hold `true` or `false`.
* `= true` → assign the initial value `true` to the variable.
* `fmt.Println("isGoFun:", isGoFun)`:

  * Prints the text `"isGoFun:"` and then the value of `isGoFun`.
  * `Println` automatically adds a space between arguments and a newline at the end.

Why this is useful:

* Boolean variables are key for conditions, flags, feature toggles, etc.
* Very common in `if` conditions and loops.

---

### Integers

```go
var age int = 25
fmt.Println("age:", age)
```

* `age` → variable name.
* `int` → integer type; size is platform-dependent (usually 64-bit on modern systems).
* `= 25` → assigns the integer 25 to `age`.
* `fmt.Println("age:", age)` prints the label and the value.

Notes:

* `int` is generally enough for counting, indexing, and simple arithmetic.
* Other integer types exist (`int8`, `int16`, `int32`, `int64`, `uint` variants), but `int` is the default workhorse for most apps.

---

### Floating point numbers (decimals)

```go
var pi float64 = 3.14159
fmt.Println("pi:", pi)
```

* `pi` → variable name.
* `float64` → 64-bit floating point type (double precision).
* `= 3.14159` → decimal value.
* `fmt.Println("pi:", pi)` prints the floating-point value.

Why `float64`:

* Go defaults to `float64` for floating point literals in most cases.
* Higher precision than `float32`, so it is preferred for general numerical calculations.

---

### Strings (text)

```go
var language string = "Go language"
fmt.Println("language:", language)
```

* `language` → variable name.
* `string` → type for text data.
* `"Go language"` → string literal.
* `fmt.Println("language:", language)` prints the label and the string value.

Strings are immutable in Go: once created, their contents cannot be changed in-place; any “change” produces a new string.

---

### Short variable declaration (`:=`)

```go
count := 10
ratio := 0.75
title := "Go notes"
```

* `:=` is called the **short variable declaration** operator.
* Declares a new variable **and** assigns a value in one step.
* The type is inferred automatically from the value on the right side:

  * `count := 10` → `count` has type `int`.
  * `ratio := 0.75` → `ratio` has type `float64`.
  * `title := "Go notes"` → `title` has type `string`.

Important rules:

* `:=` can only be used **inside** functions (not at package level).
* At least one variable on the left-hand side must be new in that scope.

Printing them:

```go
fmt.Println("count:", count)
fmt.Println("ratio:", ratio)
fmt.Println("title:", title)
```

Each line prints one variable’s value with a label.

---

### Printing the type of each variable

```go
fmt.Printf("Type of isGoFun: %T\n", isGoFun)
fmt.Printf("Type of age: %T\n", age)
fmt.Printf("Type of pi: %T\n", pi)
fmt.Printf("Type of language: %T\n", language)
fmt.Printf("Type of count: %T\n", count)
fmt.Printf("Type of ratio: %T\n", ratio)
fmt.Printf("Type of title: %T\n", title)
```

Key ideas here:

* `fmt.Printf` → formatted print.
* The format string contains **verbs** such as `%T` and `\n`:

  * `%T` → prints the type of the value.
  * `\n` → newline character.
* For example:

  * `fmt.Printf("Type of age: %T\n", age)` might output:

    ```text
    Type of age: int
    ```

This is handy for learning and debugging, especially when types are inferred using `:=`.

---

## 4. Summary of rules learned in this step

* Every Go file starts with `package <name>`.
* `package main` + `func main()` defines an executable program entry point.
* Imports are declared with `import "packageName"` and must be used.
* `var name type = value` → explicit type declaration.
* `name := value` → short declaration with type inference (inside functions only).
* Basic types seen:

  * `bool`
  * `int`
  * `float64`
  * `string`
* `fmt.Println` → prints values with a newline.
* `fmt.Printf` with `%T` → prints the type of a value.
