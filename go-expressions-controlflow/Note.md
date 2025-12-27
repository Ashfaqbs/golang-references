### Go — Expressions & Basic Control Flow (Step 3)

---

## 1. New folder and file

Example on Windows (PowerShell or CMD):

```bash
mkdir C:\code\go-expressions-controlflow
cd C:\code\go-expressions-controlflow
```

Create a file named `main.go` in that folder.

---

## 2. Program: arithmetic, comparisons, booleans, `if` (printed one by one)

```go
package main

import "fmt"

func main() {
	// 1. Basic integer arithmetic
	a := 10
	b := 3

	sum := a + b
	diff := a - b
	prod := a * b
	quotient := a / b      // integer division
	remainder := a % b     // modulus (remainder)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("sum (a + b):", sum)
	fmt.Println("diff (a - b):", diff)
	fmt.Println("prod (a * b):", prod)
	fmt.Println("quotient (a / b):", quotient)
	fmt.Println("remainder (a % b):", remainder)

	// 2. Floating-point division
	x := 10.0
	y := 3.0
	division := x / y
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("division (x / y):", division)

	// 3. Comparison operators (result is bool)
	isEqual := (a == b)
	isNotEqual := (a != b)
	isGreater := (a > b)
	isLessOrEqual := (a <= b)

	fmt.Println("isEqual (a == b):", isEqual)
	fmt.Println("isNotEqual (a != b):", isNotEqual)
	fmt.Println("isGreater (a > b):", isGreater)
	fmt.Println("isLessOrEqual (a <= b):", isLessOrEqual)

	// 4. Boolean logic
	isAdult := true
	hasTicket := false

	canEnter := isAdult && hasTicket
	canEnterWithPass := isAdult || hasTicket
	denied := !canEnter

	fmt.Println("isAdult:", isAdult)
	fmt.Println("hasTicket:", hasTicket)
	fmt.Println("canEnter (isAdult && hasTicket):", canEnter)
	fmt.Println("canEnterWithPass (isAdult || hasTicket):", canEnterWithPass)
	fmt.Println("denied (!canEnter):", denied)

	// 5. Simple if/else based on boolean
	if canEnter {
		fmt.Println("Entry status: allowed")
	} else {
		fmt.Println("Entry status: denied")
	}

	// 6. If/else-if chain with comparisons
	score := 85

	fmt.Println("score:", score)

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: below C")
	}

	// 7. If with short statement (common Go pattern)
	if bonus := score - 80; bonus > 0 {
		fmt.Println("bonus (score - 80):", bonus)
	} else {
		fmt.Println("bonus (score - 80): 0 or negative")
	}
}
```

Run:

```bash
go run main.go
```

Each value and decision is printed one by one, matching the pattern from previous steps.

---

## 3. Explanation of concepts and lines (simple language)

---

### Package and imports

```go
package main
```

* Declares that this file belongs to the `main` package.
* In `main` package, the Go runtime looks for `func main()` as the starting point of execution.

```go
import "fmt"
```

* Imports the `fmt` package from the standard library.
* Provides `Println` and `Printf` for printing formatted text.
* Required because the code calls `fmt.Println` multiple times.

Unused imports are not allowed; if `fmt` was not used, compilation would fail with an “imported and not used” error.

---

### 1. Basic integer arithmetic

```go
a := 10
b := 3
```

* Two integer variables:

  * `a` with value `10`
  * `b` with value `3`
* Type of both is inferred as `int` because the literals are integers.

```go
sum := a + b
diff := a - b
prod := a * b
quotient := a / b
remainder := a % b
```

* `sum`:

  * `a + b` → `10 + 3` → `13`
* `diff`:

  * `a - b` → `10 - 3` → `7`
* `prod`:

  * `a * b` → `10 * 3` → `30`
* `quotient`:

  * `a / b` → integer division.
  * Both `a` and `b` are `int`, so result is also `int`.
  * `10 / 3` → `3` (fractional part discarded).
* `remainder`:

  * `a % b` → remainder of integer division.
  * `10 % 3` → `1`.

Printing:

```go
fmt.Println("a:", a)
fmt.Println("b:", b)
fmt.Println("sum (a + b):", sum)
fmt.Println("diff (a - b):", diff)
fmt.Println("prod (a * b):", prod)
fmt.Println("quotient (a / b):", quotient)
fmt.Println("remainder (a % b):", remainder)
```

Each line prints a descriptive label and the value, one by one.

Key rule shown here:

* When both operands are integers, division is integer division.
* For fractional results, at least one operand must be a floating-point type.

---

### 2. Floating-point division

```go
x := 10.0
y := 3.0
division := x / y
```

* `10.0` and `3.0` are floating-point literals (`float64`).
* `x` and `y` are inferred as `float64`.
* `division` holds the exact division result as a `float64`.

  * `10.0 / 3.0` → approximately `3.3333333...`.

Printing:

```go
fmt.Println("x:", x)
fmt.Println("y:", y)
fmt.Println("division (x / y):", division)
```

This highlights the difference between integer division and floating-point division in Go.

---

### 3. Comparison operators

```go
isEqual := (a == b)
isNotEqual := (a != b)
isGreater := (a > b)
isLessOrEqual := (a <= b)
```

* Comparison operators in Go:

  * `==` → equal to
  * `!=` → not equal to
  * `>`  → greater than
  * `<`  → less than
  * `>=` → greater than or equal
  * `<=` → less than or equal
* Each comparison returns a `bool` (`true` or `false`).
* Here, `a` is `10` and `b` is `3`, so:

  * `a == b` → `false`
  * `a != b` → `true`
  * `a > b` → `true`
  * `a <= b` → `false`

Printing:

```go
fmt.Println("isEqual (a == b):", isEqual)
fmt.Println("isNotEqual (a != b):", isNotEqual)
	fmt.Println("isGreater (a > b):", isGreater)
fmt.Println("isLessOrEqual (a <= b):", isLessOrEqual)
```

Each print shows the label and the boolean result.

---

### 4. Boolean logic (`&&`, `||`, `!`)

```go
isAdult := true
hasTicket := false
```

* `isAdult` and `hasTicket` are boolean variables.
* Booleans are often used to represent conditions or flags.

```go
canEnter := isAdult && hasTicket
canEnterWithPass := isAdult || hasTicket
denied := !canEnter
```

* `&&` (logical AND):

  * `true && true` → `true`
  * `true && false` → `false`
  * `false && anything` → `false`
  * `canEnter` is `true` only if both `isAdult` and `hasTicket` are `true`.
* `||` (logical OR):

  * `true || anything` → `true`
  * `false || true` → `true`
  * `false || false` → `false`
  * `canEnterWithPass` is `true` if at least one of the conditions is `true`.
* `!` (logical NOT):

  * Flips a boolean.
  * `!true` → `false`
  * `!false` → `true`
  * `denied` is the opposite of `canEnter`.

Printing:

```go
fmt.Println("isAdult:", isAdult)
fmt.Println("hasTicket:", hasTicket)
fmt.Println("canEnter (isAdult && hasTicket):", canEnter)
fmt.Println("canEnterWithPass (isAdult || hasTicket):", canEnterWithPass)
fmt.Println("denied (!canEnter):", denied)
```

This ties raw boolean expressions to a simple real-world scenario (entry control).

---

### 5. Simple `if/else` with a boolean

```go
if canEnter {
	fmt.Println("Entry status: allowed")
} else {
	fmt.Println("Entry status: denied")
}
```

Key points:

* `if` expects a condition of type `bool`.
* No parentheses required around the condition:

  * `if canEnter { ... }` is valid.
  * `if (canEnter) { ... }` also compiles, but idiomatic Go omits the parentheses.
* `{` must be on the same line as the `if` statement.
* `else` must be on the same line as the closing brace of the `if` block.

Logic here:

* If `canEnter` is `true`, the program prints `"Entry status: allowed"`.
* Otherwise, it prints `"Entry status: denied"`.

---

### 6. `if / else if / else` chain with comparisons

```go
score := 85

fmt.Println("score:", score)

if score >= 90 {
	fmt.Println("Grade: A")
} else if score >= 75 {
	fmt.Println("Grade: B")
} else if score >= 60 {
	fmt.Println("Grade: C")
} else {
	fmt.Println("Grade: below C")
}
```

* `score` is an `int` with value `85`.
* The chain is evaluated from top to bottom:

  1. `score >= 90`?

     * For `85`, this is `false`, so the next condition is checked.
  2. `score >= 75`?

     * This is `true`, so `"Grade: B"` is printed.
  3. Remaining `else if` and `else` blocks are skipped after the first match.

Important rules illustrated:

* The first `if` or `else if` whose condition is `true` has its block executed; the rest are skipped.
* An `else` block (if present) catches all remaining cases where none of the previous conditions were `true`.

---

### 7. `if` with a short statement

```go
if bonus := score - 80; bonus > 0 {
	fmt.Println("bonus (score - 80):", bonus)
} else {
	fmt.Println("bonus (score - 80): 0 or negative")
}
```

This line does two things at once:

1. `bonus := score - 80`

   * Short variable declaration inside the `if`.
   * `bonus` exists only inside this `if-else` block (limited scope).
2. `bonus > 0`

   * Condition checked by `if`.

If `score` is `85`, then:

* `bonus := 85 - 80` → `5`
* Condition `bonus > 0` → `true`
* So `"bonus (score - 80): 5"` is printed.

If `score` was `80` or less:

* `bonus` would be `0` or negative.
* The `else` branch would be executed.

Why this pattern matters:

* Keeps the variable `bonus` tightly scoped.
* Common in Go for situations like:

  * Parsing values.
  * Checking errors.
  * Computing temporary values only needed for one `if`.

---

## 4. Summary of core ideas from this step

* Basic arithmetic operators on integers:

  * `+`, `-`, `*`, `/`, `%`
  * Integer division discards the fractional part.
* Floating-point division requires floating-point operands (`float64`).
* Comparison operators:

  * `==`, `!=`, `>`, `<`, `>=`, `<=`
  * Always return `bool`.
* Boolean operators:

  * `&&` (AND), `||` (OR), `!` (NOT).
* `if` statements:

  * Condition must be `bool`.
  * Parentheses around condition are optional but usually omitted.
  * Braces are mandatory and follow strict placement rules.
* `if / else if / else` chain:

  * Evaluated from top to bottom.
  * First matching branch runs; others are skipped.
* `if` with a short statement:

  * Allows local variable declaration that exists only inside the `if` and `else` blocks.