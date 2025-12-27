### Go — Arrays and Slices (Step 5: fixed-size vs dynamic collections)

---

## 1. New folder and file

Example on Windows (PowerShell or CMD):

```bash
mkdir C:\code\go-arrays-slices
cd C:\code\go-arrays-slices
```

Create `main.go` with the following content.

---

## 2. Program: arrays, slices, `len`, `append`, printed one by one

```go
package main

import "fmt"

func main() {
	// 1. Fixed-size array
	var arr [4]int
	fmt.Println("arr (zero values):", arr)
	fmt.Println("len(arr):", len(arr))

	// 2. Array with initialization
	scores := [4]int{10, 20, 30, 40}
	fmt.Println("\nscores:", scores)
	fmt.Println("len(scores):", len(scores))

	// 3. Indexing and updating array elements
	fmt.Println("\nAccess elements of scores:")
	fmt.Println("scores[0]:", scores[0])
	fmt.Println("scores[1]:", scores[1])

	scores[2] = 35
	fmt.Println("scores after update (scores[2] = 35):", scores)

	// 4. Slice from array (view over part of the array)
	sub := scores[1:3]
	fmt.Println("\nsub slice (scores[1:3]):", sub)
	fmt.Println("len(sub):", len(sub))

	// 5. Slice literal (direct slice creation, backing array created automatically)
	nums := []int{100, 200, 300}
	fmt.Println("\nnums slice:", nums)
	fmt.Println("len(nums):", len(nums))

	// 6. Append to slice (dynamic growth)
	nums = append(nums, 400)
	fmt.Println("nums after append 400:", nums)

	nums = append(nums, 500, 600)
	fmt.Println("nums after append 500, 600:", nums)

	// 7. Make slice with make (length and capacity)
	data := make([]int, 3) // length 3, capacity 3
	fmt.Println("\ndata slice created with make:", data)
	fmt.Println("len(data):", len(data))

	data[0] = 7
	data[1] = 14
	data[2] = 21
	fmt.Println("data after assignments:", data)

	// 8. Append beyond initial length
	data = append(data, 28)
	fmt.Println("data after append 28:", data)
	fmt.Println("len(data):", len(data))

	// 9. Loop over slice with range
	fmt.Println("\nLoop over nums slice with range:")
	for index, value := range nums {
		fmt.Println("index:", index, "value:", value)
	}
}
```

Run:

```bash
go run main.go
```

---

## 3. Explanation of every part

---

### Package and import

```go
package main
```

* Declares that this file is part of the `main` package.
* A `main` package with a `main` function builds into an executable program.

```go
import "fmt"
```

* Imports the `fmt` package from the standard library.
* Provides formatted I/O functions: `Println`, `Printf`, etc.
* Used throughout the file for printing values and labels.
* Go requires imported packages to be used; unused imports cause compile-time errors.

---

### 1. Fixed-size array

```go
var arr [4]int
fmt.Println("arr (zero values):", arr)
fmt.Println("len(arr):", len(arr))
```

* `var arr [4]int`:

  * Declares a variable named `arr`.
  * Type is `[4]int` → array of **exactly 4 integers**.
  * Fixed size; length is part of the type.
* No explicit values are assigned, so elements receive zero values.

  * For `int`, zero value is `0`.
  * So `arr` is `[0 0 0 0]`.
* `fmt.Println("arr (zero values):", arr)`:

  * Prints a label and the contents of the array.
* `len(arr)`:

  * Built-in function `len` returns the length of arrays, slices, strings, etc.
  * For this array, `len(arr)` is `4`.

Arrays in Go have:

* Fixed length.
* Value semantics (assigning an array creates a copy).

---

### 2. Array with initialization

```go
scores := [4]int{10, 20, 30, 40}
fmt.Println("\nscores:", scores)
fmt.Println("len(scores):", len(scores))
```

* `scores := [4]int{10, 20, 30, 40}`:

  * Short variable declaration.
  * Array literal `[4]int{...}`:

    * Type: array of 4 integers.
    * Values: `10`, `20`, `30`, `40`.
* `"\n"` at the beginning of the string:

  * Newline for visual separation in output.
* `fmt.Println("scores:", scores)`:

  * Prints `[10 20 30 40]`.
* `len(scores)`:

  * Returns `4` again.

Arrays are best when size is known and constant.

---

### 3. Indexing and updating array elements

```go
fmt.Println("\nAccess elements of scores:")
fmt.Println("scores[0]:", scores[0])
fmt.Println("scores[1]:", scores[1])
```

* Arrays and slices use **zero-based indexing**.
* `scores[0]`:

  * First element → `10`.
* `scores[1]`:

  * Second element → `20`.
* Each `Println` call prints label and accessed element.

Updating:

```go
scores[2] = 35
fmt.Println("scores after update (scores[2] = 35):", scores)
```

* `scores[2]` was `30`.
* Assignment sets it to `35`.
* After update, `scores` becomes `[10 20 35 40]`.
* Index access and assignment rely on valid index range:

  * Valid indices: `0` to `len(scores)-1`.
  * Out-of-range access panics at runtime.

---

### 4. Slice from array (subrange view)

```go
sub := scores[1:3]
fmt.Println("\nsub slice (scores[1:3]):", sub)
fmt.Println("len(sub):", len(sub))
```

* `scores[1:3]`:

  * **Slice expression**.
  * Creates a **slice view** from index `1` (inclusive) to `3` (exclusive).
  * Elements: `scores[1]`, `scores[2]`.
* With `scores` as `[10 20 35 40]`, `sub` becomes `[20 35]`.
* Type of `sub` is `[]int` (slice of int).
* `len(sub)` is `2`.

Important concept:

* Slices reference the same underlying array.
* Changes through the slice can affect the array (and vice versa) if the same elements are involved.

---

### 5. Slice literal (direct slice creation)

```go
nums := []int{100, 200, 300}
fmt.Println("\nnums slice:", nums)
fmt.Println("len(nums):", len(nums))
```

* `[]int{100, 200, 300}`:

  * Slice literal.
  * Under the hood, Go allocates an array and creates a slice referring to it.
* `nums` is a slice (`[]int`), not an array.
* `len(nums)` is `3`.

Slices are more common than arrays in everyday Go code due to their dynamic size behavior and lighter semantics.

---

### 6. Appending to a slice (dynamic growth)

```go
nums = append(nums, 400)
fmt.Println("nums after append 400:", nums)

nums = append(nums, 500, 600)
fmt.Println("nums after append 500, 600:", nums)
```

* `append` is a built-in function that adds elements to a slice.
* First call:

  * `append(nums, 400)`:

    * Returns a new slice containing previous elements plus `400`.
  * Assignment `nums = ...` updates `nums` to this new slice.
* Second call:

  * `append(nums, 500, 600)`:

    * Appends `500` and `600` in a single call.
* After both appends:

  * Start: `[100 200 300]`
  * After `400`: `[100 200 300 400]`
  * After `500, 600`: `[100 200 300 400 500 600]`

Important semantics:

* `append` may allocate a new underlying array if capacity is exceeded.
* Always assign the result of `append` back to a slice variable to capture possible new backing storage.

---

### 7. Creating slice with `make` (length and capacity)

```go
data := make([]int, 3) // length 3, capacity 3
fmt.Println("\ndata slice created with make:", data)
fmt.Println("len(data):", len(data))
```

* `make([]int, 3)`:

  * Creates a slice of type `[]int` with:

    * Length: `3`
    * Capacity: `3` (here, same as length, because capacity argument is omitted)
* Elements are initialized to zero value:

  * `[0 0 0]`.
* `len(data)` returns `3`.

`make` is used for creating:

* Slices (`[]T`)
* Maps (`map[K]V`)
* Channels (`chan T`)

with controlled initial size/capacity.

Assignments:

```go
data[0] = 7
data[1] = 14
data[2] = 21
fmt.Println("data after assignments:", data)
```

* Index-based assignment to fill the slice.
* Result: `[7 14 21]`.

---

### 8. Append beyond initial length

```go
data = append(data, 28)
fmt.Println("data after append 28:", data)
fmt.Println("len(data):", len(data))
```

* Initial `data`: `[7 14 21]` with length `3`.
* `append(data, 28)`:

  * New slice with `28` added.
  * Under the hood, capacity may be expanded; Go handles allocation automatically.
* After append:

  * `data` becomes `[7 14 21 28]`.
  * `len(data)` is `4`.

Key idea:

* Slices provide dynamic behavior: length can grow via `append` even if they were created with a smaller initial length.
* Arrays cannot change size.

---

### 9. Looping over a slice with `range`

```go
fmt.Println("\nLoop over nums slice with range:")
for index, value := range nums {
	fmt.Println("index:", index, "value:", value)
}
```

* `range nums`:

  * Iterates over the slice `nums`.
  * On each iteration:

    * `index` gets the position (`0`, `1`, `2`, ...).
    * `value` gets the element at that position.
* Printing:

  * Each line shows index and value pair.

Typical output pattern:

* `index: 0 value: 100`
* `index: 1 value: 200`
* etc., depending on `nums` contents at that point.

`range` + slices is the standard way to inspect elements sequentially.

---

## 4. Summary of core ideas from this step

* Arrays:

  * Syntax: `[N]T` (e.g., `[4]int`).
  * Fixed size, length is part of the type.
  * Zero values fill all elements if not explicitly initialized.
* Slices:

  * Syntax: `[]T` (e.g., `[]int`).
  * Dynamic view on an underlying array.
  * Variable length; commonly used in Go code.
  * Created via:

    * Slice literals: `[]int{1, 2, 3}`
    * Slicing arrays or other slices: `arr[1:3]`
    * `make([]int, length)` (optionally with capacity).
* `len`:

  * Works on arrays, slices, strings.
  * Returns current length.
* `append`:

  * Adds elements to a slice.
  * May allocate new backing storage.
  * Result should be assigned back to a slice variable.
* `range`:

  * Convenient iteration over slices:

    ```go
    for index, value := range slice { ... }
    ```
* Arrays resemble low-level, fixed-size containers.
* Slices are the primary dynamic collection abstraction used in typical Go programs.