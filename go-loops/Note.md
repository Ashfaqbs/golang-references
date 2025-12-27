### Go — `for` Loops (Step 4: counting, conditions, break, continue)

---

## 1. New folder and file

Example on Windows (PowerShell or CMD):

```bash
mkdir C:\code\go-loops
cd C:\code\go-loops
```

Create a file named `main.go`.

---

## 2. Program: different `for` loop forms, printed one by one

```go
package main

import "fmt"

func main() {
	// 1. Classic counting for-loop
	fmt.Println("Classic for loop (0 to 4):")
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	// 2. While-style loop (condition-only for)
	fmt.Println("\nWhile-style loop (countdown from 3):")
	j := 3
	for j > 0 {
		fmt.Println("j:", j)
		j--
	}

	// 3. Infinite loop with break
	fmt.Println("\nInfinite loop with break (stop at k == 3):")
	k := 0
	for {
		fmt.Println("k:", k)
		if k == 3 {
			fmt.Println("breaking loop at k == 3")
			break
		}
		k++
	}

	// 4. Loop with continue (skip odd numbers)
	fmt.Println("\nLoop with continue (print only even numbers from 0 to 5):")
	for n := 0; n <= 5; n++ {
		if n%2 != 0 {
			continue
		}
		fmt.Println("even n:", n)
	}

	// 5. Loop over a slice with range
	fmt.Println("\nLoop over slice with range:")
	numbers := []int{10, 20, 30, 40}

	for index, value := range numbers {
		fmt.Println("index:", index, "value:", value)
	}

	// 6. Loop over slice with range, ignoring index
	fmt.Println("\nLoop over slice with range (values only):")
	for _, value := range numbers {
		fmt.Println("value:", value)
	}
}
```

Command to run:

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

* Declares the `main` package.
* In a `main` package, the Go runtime looks for `func main()` as the entry point.

```go
import "fmt"
```

* Imports the `fmt` standard library package.
* `fmt` provides text printing functions such as `Println`.
* Required in this file because `fmt.Println` is used multiple times.
* Go does not allow unused imports; removal of all `fmt` usage would cause a compile error.

---

### 1. Classic counting `for` loop

```go
fmt.Println("Classic for loop (0 to 4):")
for i := 0; i < 5; i++ {
	fmt.Println("i:", i)
}
```

Breakdown of the `for` line:

```go
for i := 0; i < 5; i++ {
```

* `i := 0`

  * Initialization.
  * Declares `i` and sets it to `0`.
  * This initialization runs once, before the loop starts.
* `i < 5`

  * Condition checked before each iteration.
  * As long as this is `true`, the loop body runs.
* `i++`

  * Post statement.
  * Runs after each iteration of the loop body.
  * Equivalent to `i = i + 1`.

Loop behavior:

* Iteration 1: `i = 0` → condition `0 < 5` → `true` → prints `i: 0` → `i` becomes `1`.
* Iteration 2: `i = 1` → condition `1 < 5` → `true` → prints `i: 1` → `i` becomes `2`.
* Continues until `i` reaches `5`.
* When `i = 5`, condition `5 < 5` → `false`, loop stops.

This form closely matches the classic C/Java `for` syntax.

---

### 2. While-style `for` loop (condition only)

```go
fmt.Println("\nWhile-style loop (countdown from 3):")
j := 3
for j > 0 {
	fmt.Println("j:", j)
	j--
}
```

* `j := 3`

  * Declares and initializes `j` with `3`.
* `for j > 0 { ... }`

  * `for` with only a condition.
  * This is Go’s version of a `while` loop.

Loop behavior:

* Iteration 1: `j = 3` → `j > 0` → `true` → prints `j: 3` → `j--` → `j = 2`.
* Iteration 2: `j = 2` → `true` → prints `j: 2` → `j = 1`.
* Iteration 3: `j = 1` → `true` → prints `j: 1` → `j = 0`.
* Next check: `j = 0` → `j > 0` → `false` → loop finishes.

In Go, `for` is the only loop keyword; there is no separate `while` keyword.
A `while`-style loop is simply `for` with a condition.

`"\n"` in `fmt.Println` calls:

* `"\n"` is a newline character.
* Placed at the start of the string to visually separate different sections in the output.

---

### 3. Infinite loop with `break`

```go
fmt.Println("\nInfinite loop with break (stop at k == 3):")
k := 0
for {
	fmt.Println("k:", k)
	if k == 3 {
		fmt.Println("breaking loop at k == 3")
		break
	}
	k++
}
```

`for { ... }`:

* `for` with no condition and no init/post parts.
* This is an infinite loop by design.
* It must be controlled with `break`, `return`, or some exit condition inside the body.

Flow:

1. `k := 0`
2. `for {` starts an infinite loop.
3. `fmt.Println("k:", k)` prints the current value of `k`.
4. `if k == 3 { ... }`:

   * Condition is checked.
   * When `k` equals `3`, the `break` statement runs.
   * `break` exits the nearest enclosing `for` loop.
5. If `k` is not `3`, `k++` increments `k` by `1`, and the loop repeats.

Values printed for `k`: `0`, `1`, `2`, `3`, and then the loop breaks after printing the message.

---

### 4. Loop with `continue` (skip odd numbers)

```go
fmt.Println("\nLoop with continue (print only even numbers from 0 to 5):")
for n := 0; n <= 5; n++ {
	if n%2 != 0 {
		continue
	}
	fmt.Println("even n:", n)
}
```

* `for n := 0; n <= 5; n++ { ... }`:

  * `n` starts at `0`.
  * Loop continues while `n <= 5`.
  * `n++` increments `n` after each iteration.

Inside the loop:

```go
if n%2 != 0 {
	continue
}
fmt.Println("even n:", n)
```

* `n % 2`:

  * Remainder when `n` is divided by `2`.
* `n%2 != 0`:

  * Condition is `true` for odd numbers.
* `continue`:

  * Skips the rest of the loop body for the current iteration.
  * Control jumps to the next iteration (`n++` happens, then condition is checked again).

Effect:

* For `n = 0` → `0%2 != 0` → `false` → prints `even n: 0`.
* For `n = 1` → `1%2 != 0` → `true` → `continue` → skip print.
* For `n = 2` → prints `even n: 2`.
* For `n = 3` → skipped.
* For `n = 4` → printed.
* For `n = 5` → skipped.

Only even numbers from `0` to `5` are printed.

---

### 5. Loop over a slice with `range` (index and value)

```go
fmt.Println("\nLoop over slice with range:")
numbers := []int{10, 20, 30, 40}

for index, value := range numbers {
	fmt.Println("index:", index, "value:", value)
}
```

First, the slice:

```go
numbers := []int{10, 20, 30, 40}
```

* `[]int{...}` creates a **slice** of `int`.
* A slice is a flexible view over an underlying array (resizable sequence abstraction).
* Here, `numbers` has four elements: `10`, `20`, `30`, `40`.

The `range` loop:

```go
for index, value := range numbers {
	...
}
```

* `range numbers`:

  * Iterates over the slice.
  * Returns two values per iteration:

    * Index (position, starting from `0`).
    * Value at that index.
* `index, value :=`:

  * Declares both `index` and `value`.
  * Types:

    * `index` → `int`
    * `value` → `int` (same as slice element type)

Iterations:

* 1st: `index = 0`, `value = 10`
* 2nd: `index = 1`, `value = 20`
* 3rd: `index = 2`, `value = 30`
* 4th: `index = 3`, `value = 40`

Each iteration prints both index and value.

---

### 6. Loop over slice with `range`, ignoring index

```go
fmt.Println("\nLoop over slice with range (values only):")
for _, value := range numbers {
	fmt.Println("value:", value)
}
```

* `_` is the **blank identifier** in Go.
* When `_` is used on the left side of a `range`, it means:

  * “Ignore this returned value.”
* `range numbers` still returns index and value, but the index is discarded.

Reason for `_`:

* Go does not allow unused variables.
* If the index variable was declared and not used, compilation would fail.
* `_` solves this by explicitly discarding the unused value.

Effect:

* Same values as before are printed, but only values are shown:

  * `10`, `20`, `30`, `40`.

---

## 4. Summary of `for` loop patterns in Go

* Go has a single loop keyword: `for`.
  It covers:

  * Classic `for` loop:

    ```go
    for i := 0; i < n; i++ { ... }
    ```
  * While-style loop:

    ```go
    for condition { ... }
    ```
  * Infinite loop:

    ```go
    for { ... }
    ```
* `break` exits the nearest enclosing `for` loop.
* `continue` skips to the next iteration of the loop.
* `range` provides a convenient way to loop over elements of slices, arrays, maps, strings, and channels (later topics), usually as:

  ```go
  for index, value := range collection { ... }
  ```
* `_` (blank identifier) discards an unwanted value, avoiding “unused variable” errors.
