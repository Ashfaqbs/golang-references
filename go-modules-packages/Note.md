### Go — Modules and Packages (Step 10: multi-file, multi-package layout)

---

## 1. New folder and basic module setup

Example on Windows (PowerShell or CMD):

```bash
mkdir C:\code\go-modules-packages
cd C:\code\go-modules-packages
```

Initialize a Go module:

```bash
go mod init example.com/go-notes
```

This creates a `go.mod` file.
Example content:

```go
module example.com/go-notes

go 1.24
```

Explanation:

* `module example.com/go-notes`

  * Declares the **module path**.
  * This is the base path used in import statements inside this module.
  * Does not need to be a real domain during local development; any unique string works.
* `go 1.24`

  * Declares the Go language version the module is written for.
  * Tools use this to enable the right language features.

The module is now ready for multiple packages and files.

---

## 2. Project structure for this step

Target structure:

```text
go-modules-packages/
  go.mod
  main.go
  greet/
    greet.go
```

* `main.go` → entrypoint (`package main`).
* `greet/greet.go` → internal package with reusable functions.

---

## 3. `main.go`: use a custom package

Create `main.go` in `C:\code\go-modules-packages`:

```go
package main

import (
	"fmt"

	"example.com/go-notes/greet"
)

func main() {
	fmt.Println("In main function")

	message := greet.Hello("Go learner")
	fmt.Println("greet.Hello returned:", message)

	byeMessage := greet.Goodbye("Go learner")
	fmt.Println("greet.Goodbye returned:", byeMessage)
}
```

### Explanation of `main.go`

#### `package main`

```go
package main
```

* Declares that this file belongs to the `main` package.
* A program with a `main` package and a `main` function builds into an executable binary.
* `package main` is reserved for executables, not for reusable libraries.

---

#### `import` block

```go
import (
	"fmt"

	"example.com/go-notes/greet"
)
```

* `import (...)`:

  * Import block for multiple packages.
* `"fmt"`:

  * Standard library package for formatted I/O.
  * Provides `Println` used later in `main`.
* `"example.com/go-notes/greet"`:

  * Import path for the custom package.
  * Built as: `<module path>/<folder name>`.
  * `module` path is from `go.mod`: `example.com/go-notes`.
  * Folder `greet` contains `package greet`.

This matches the pattern:

```text
module: example.com/go-notes
folder: greet
import path: example.com/go-notes/greet
```

---

#### `func main()`

```go
func main() {
	fmt.Println("In main function")

	message := greet.Hello("Go learner")
	fmt.Println("greet.Hello returned:", message)

	byeMessage := greet.Goodbye("Go learner")
	fmt.Println("greet.Goodbye returned:", byeMessage)
}
```

* `func main()`:

  * Program’s entrypoint.
* `fmt.Println("In main function")`:

  * Confirms that `main` is running.

Calling the custom package:

```go
message := greet.Hello("Go learner")
fmt.Println("greet.Hello returned:", message)
```

* `greet.Hello(...)`:

  * Calls the `Hello` function from the `greet` package.
  * Returns a string.
* Assigned to `message` and printed.

Second call:

```go
byeMessage := greet.Goodbye("Go learner")
fmt.Println("greet.Goodbye returned:", byeMessage)
```

* `greet.Goodbye(...)`:

  * Another function in the same `greet` package.
  * Returns a farewell message.
* Assigned to `byeMessage` and printed.

The `main` package depends on `greet` as a reusable internal library.

---

## 4. `greet/greet.go`: define the custom package

Create the folder:

```bash
mkdir greet
```

Then create `greet\greet.go` with content:

```go
package greet

import "fmt"

// Hello returns a greeting message for the given name.
func Hello(name string) string {
	message := fmt.Sprintf("Hello, %s! Welcome to Go packages.", name)
	return message
}

// Goodbye returns a farewell message for the given name.
func Goodbye(name string) string {
	message := fmt.Sprintf("Goodbye, %s! See you again.", name)
	return message
}
```

### Explanation of `greet/greet.go`

#### `package greet`

```go
package greet
```

* Declares a new package named `greet`.
* All `.go` files in the `greet` folder that start with `package greet` belong to the same package.
* The package name is usually the same as the folder name, but does not have to be. Matching names are idiomatic.

---

#### Import inside the `greet` package

```go
import "fmt"
```

* Imports `fmt` here because `greet` needs `fmt.Sprintf`.
* Package imports are local to each file:

  * `main.go` and `greet.go` both import `fmt` separately.
* Unused imports in any file cause a compile-time error.

---

#### Exported function: `Hello`

```go
// Hello returns a greeting message for the given name.
func Hello(name string) string {
	message := fmt.Sprintf("Hello, %s! Welcome to Go packages.", name)
	return message
}
```

Important details:

* Function name `Hello` starts with a **capital letter**.

  * In Go, capitalized identifiers are **exported** from a package.
  * Exported = visible to other packages that import this one.
* Signature:

  * Parameter: `name string`
  * Return type: `string`.
* Body:

  * `fmt.Sprintf` builds a string without printing.
  * `message` holds the formatted string.
  * `return message` returns the value to the caller.

This function is accessible from other packages as `greet.Hello`.

---

#### Exported function: `Goodbye`

```go
// Goodbye returns a farewell message for the given name.
func Goodbye(name string) string {
	message := fmt.Sprintf("Goodbye, %s! See you again.", name)
	return message
}
```

* Same structure as `Hello`.
* Different text in the formatted string.
* Also exported due to the capitalized name `Goodbye`.

Exports vs non-exports rule:

* Identifiers starting with an uppercase letter → exported (public).
* Identifiers starting with a lowercase letter → unexported (package-internal).

---

## 5. Running the module

From inside `C:\code\go-modules-packages`:

```bash
go run ./...
```

or:

```bash
go run main.go
```

`go run ./...`:

* Builds and runs the `main` package.
* Automatically resolves imports within the module, including `example.com/go-notes/greet`.

Output will show:

* Initial message from `main`.
* Return value from `greet.Hello`.
* Return value from `greet.Goodbye`.

---

## 6. Multiple files in the same package

Packages often span several files. For example, the `greet` package can be split into multiple files.

Add another file `greet/more_greet.go`:

```go
package greet

// InternalMessage is not exported because the name starts with a lowercase letter.
func internalMessage() string {
	return "internal only"
}

// ShoutHello is exported and can be used from other packages.
func ShoutHello(name string) string {
	msg := Hello(name)
	return msg + "!!!"
}
```

Key points:

* `package greet` again: same package, different file.
* `internalMessage`:

  * Starts with lowercase.
  * Only code **inside** the `greet` package can call it.
  * `main` cannot access `greet.internalMessage()`.
* `ShoutHello`:

  * Starts with uppercase.
  * Exported and can be used from `main` as `greet.ShoutHello`.

`main.go` can now call:

```go
shout := greet.ShoutHello("Go learner")
fmt.Println("greet.ShoutHello returned:", shout)
```

This demonstrates:

* One package can be split across multiple `.go` files.
* Export control is done by capitalization, not by keywords like `public` / `private`.

---

## 7. Summary of modules and packages

* `go mod init <module-path>`:

  * Creates `go.mod`.
  * Module path becomes the root for import paths.
* `package main`:

  * Defines an executable program.
* Non-main packages (like `package greet`):

  * Hold reusable logic.
  * Imported using the module path plus folder name.
* Export rules:

  * Capitalized names (e.g., `Hello`, `Person`, `Service`) are exported.
  * Lowercase names are internal to the package.
* Multiple files with the same `package` name:

  * Form a single package.
  * Allow splitting related logic across files.
* Imports are per file and must only include used packages.

This step completes the core idea of multi-file, multi-package Go projects, mirroring the “multiple .java files + packages” stage from Java.