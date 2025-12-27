### Go — Maps (Step 6: key–value storage)

---

## 1. New folder and file

Example on Windows (PowerShell or CMD):

```bash
mkdir C:\code\go-maps
cd C:\code\go-maps
```

Create `main.go` with the following content.

---

## 2. Program: map creation, access, update, delete, `range`

```go
package main

import "fmt"

func main() {
	// 1. Map literal with string keys and int values
	scores := map[string]int{
		"Alice": 90,
		"Bob":   75,
	}
	fmt.Println("scores map:", scores)

	// 2. Access value by key
	fmt.Println("\nAccess values by key:")
	fmt.Println(`scores["Alice"]:`, scores["Alice"])
	fmt.Println(`scores["Bob"]:`, scores["Bob"])

	// 3. Add new key–value pair
	scores["Charlie"] = 82
	fmt.Println("\nAfter adding Charlie:")
	fmt.Println("scores map:", scores)

	// 4. Update existing value
	scores["Bob"] = 80
	fmt.Println("\nAfter updating Bob to 80:")
	fmt.Println("scores map:", scores)

	// 5. Access non-existing key (zero value behavior)
	fmt.Println("\nAccess non-existing key:")
	fmt.Println(`scores["David"]:`, scores["David"])

	// 6. Check if key exists using the comma-ok idiom
	value, ok := scores["David"]
	fmt.Println("\nCheck presence of key \"David\":")
	fmt.Println("value:", value)
	fmt.Println("exists (ok):", ok)

	value2, ok2 := scores["Alice"]
	fmt.Println("\nCheck presence of key \"Alice\":")
	fmt.Println("value:", value2)
	fmt.Println("exists (ok):", ok2)

	// 7. Delete a key
	delete(scores, "Charlie")
	fmt.Println("\nAfter deleting Charlie:")
	fmt.Println("scores map:", scores)

	// 8. Create map with make
	ages := make(map[string]int)
	ages["Tom"] = 30
	ages["Jerry"] = 28

	fmt.Println("\nages map created with make:")
	fmt.Println("ages:", ages)

	// 9. Loop over map with range
	fmt.Println("\nLoop over scores map with range:")
	for name, score := range scores {
		fmt.Println("name:", name, "score:", score)
	}
}
```

Run:

```bash
go run main.go
```

Output will list map values and changes step by step.

---

## 3. Explanation of every part

---

### Package and import

```go
package main
```

* Declares that this file belongs to the `main` package.
* A `main` package with a `main` function is compiled as an executable program.

```go
import "fmt"
```

* Imports the `fmt` standard library package.
* Provides printing functions, such as `Println` and `Printf`.
* This import is necessary because the code prints many intermediate values.
* Unused imports are not allowed; removing all `fmt` usage would cause a compile error.

---

### 1. Map literal with string keys and int values

```go
scores := map[string]int{
	"Alice": 90,
	"Bob":   75,
}
fmt.Println("scores map:", scores)
```

* `map[string]int`:

  * Type: map with keys of type `string` and values of type `int`.
* `{ "Alice": 90, "Bob": 75 }`:

  * Map literal that initializes the map with two entries.
* `scores`:

  * Variable holding this map.
* A map in Go:

  * Is a built-in associative container.
  * Stores key–value pairs with fast lookups by key.
* `fmt.Println("scores map:", scores)` prints the current contents, for example:

  * `scores map: map[Alice:90 Bob:75]`
  * The order of keys is **not guaranteed**; maps in Go are unordered.

---

### 2. Access value by key

```go
fmt.Println("\nAccess values by key:")
fmt.Println(`scores["Alice"]:`, scores["Alice"])
fmt.Println(`scores["Bob"]:`, scores["Bob"])
```

* `scores["Alice"]`:

  * Retrieves the value associated with key `"Alice"`.
  * Returns `90` with the current initialization.
* `scores["Bob"]`:

  * Retrieves `75` initially (before later updates).
* If a key exists, the map returns the stored value.
* If a key does not exist, the map returns the **zero value** for the value type (explored further below).

The backticks around the string in `fmt.Println` (e.g. `` `scores["Alice"]:` ``):

* Represent a raw string literal, making it easier to include quotes without escaping.

---

### 3. Add a new key–value pair

```go
scores["Charlie"] = 82
fmt.Println("\nAfter adding Charlie:")
fmt.Println("scores map:", scores)
```

* `scores["Charlie"] = 82`:

  * Inserts a new entry with key `"Charlie"` and value `82`.
  * If `"Charlie"` did not exist earlier, it is added.
* Maps in Go grow dynamically as new key–value pairs are added.
* The `fmt.Println` call displays the updated map, now containing three keys.

---

### 4. Update existing value

```go
scores["Bob"] = 80
fmt.Println("\nAfter updating Bob to 80:")
fmt.Println("scores map:", scores)
```

* Using an existing key assigns a new value for that key.
* `"Bob"` previously had value `75`, now it becomes `80`.
* This is how map entries are updated in-place.

---

### 5. Access non-existing key (zero value behavior)

```go
fmt.Println("\nAccess non-existing key:")
fmt.Println(`scores["David"]:`, scores["David"])
```

* `"David"` has not been added to the map.
* When a map is indexed with a key that does not exist:

  * The result is the **zero value** of the value type.
  * Here, the value type is `int`, so the zero value is `0`.
* This makes it impossible to distinguish between:

  * “Key is present with value `0`”
  * “Key is not present and zero value is returned”
* The comma-ok idiom fixes this by providing a second boolean result.

---

### 6. Check if key exists (comma-ok idiom)

```go
value, ok := scores["David"]
fmt.Println("\nCheck presence of key \"David\":")
fmt.Println("value:", value)
fmt.Println("exists (ok):", ok)
```

* `scores["David"]` in this form:

  * Returns two values:

    * `value` → `0` (zero value for `int`).
    * `ok` → `false`, indicating key not present.
* `ok`:

  * Boolean flag that is `true` only if the key exists in the map.

```go
value2, ok2 := scores["Alice"]
fmt.Println("\nCheck presence of key \"Alice\":")
fmt.Println("value:", value2)
fmt.Println("exists (ok):", ok2)
```

* For `"Alice"`:

  * `value2` → `90`.
  * `ok2` → `true`.
* This two-value form is the standard way to safely check both value and existence.

This pattern is often referred to as the “comma-ok” idiom in Go.

---

### 7. Delete a key

```go
delete(scores, "Charlie")
fmt.Println("\nAfter deleting Charlie:")
fmt.Println("scores map:", scores)
```

* `delete` is a built-in function for maps.
* `delete(scores, "Charlie")`:

  * Removes the entry with key `"Charlie"`, if it exists.
  * If the key does not exist, `delete` does nothing (no panic).
* After deletion, `"Charlie"` no longer appears in the map.

---

### 8. Create map with `make`

```go
ages := make(map[string]int)
ages["Tom"] = 30
ages["Jerry"] = 28

fmt.Println("\nages map created with make:")
fmt.Println("ages:", ages)
```

* `make(map[string]int)`:

  * Creates and initializes an empty map with key type `string` and value type `int`.
  * The map is ready to use; no further initialization is required.
* `ages["Tom"] = 30` and `ages["Jerry"] = 28`:

  * Add two entries into the `ages` map.
* `make` is the standard way to allocate and initialize:

  * Maps
  * Slices
  * Channels

Map creation options:

* Map literal with initial data:

  ```go
  m := map[string]int{"A": 1, "B": 2}
  ```
* Empty map with make:

  ```go
  m := make(map[string]int)
  ```

---

### 9. Loop over map with `range`

```go
fmt.Println("\nLoop over scores map with range:")
for name, score := range scores {
	fmt.Println("name:", name, "score:", score)
}
```

* `range scores`:

  * Iterates over all key–value pairs in the map.
  * Each iteration returns:

    * A key (`name`).
    * The corresponding value (`score`).
* `name` type:

  * `string` (same as map key type).
* `score` type:

  * `int` (same as map value type).
* Output shows each pair.

Important detail:

* Map iteration order in Go is **not guaranteed**.
* Different runs may produce different key orders.
* Any logic that depends on map order is considered incorrect.

If only keys or only values are needed:

* Keys only:

  ```go
  for name := range scores {
      fmt.Println("name:", name)
  }
  ```
* Values only (with blank identifier):

  ```go
  for _, score := range scores {
      fmt.Println("score:", score)
  }
  ```

---

## 4. Summary of core ideas from this step

* Map type:

  * `map[K]V` where `K` is key type, `V` is value type.
* Construction:

  * Literal: `map[string]int{"A": 1}`
  * With `make`: `make(map[string]int)`
* Access:

  * `m[key]` returns value or zero value if key does not exist.
* Comma-ok idiom:

  * `value, ok := m[key]`

    * `ok` is `true` if key exists, `false` otherwise.
* Update and insert:

  * `m[key] = value`
* Delete:

  * `delete(m, key)`
* Iteration:

  * `for k, v := range m { ... }`
  * Order is not defined and may change from run to run.