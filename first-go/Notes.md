
```bash
go version go1.24.11 windows/amd64
```

The toolchain is ready. The next step is to mirror a Java learning journey, but in Go-land.

The approach is **very step-by-step**, with explanations for **every line** and **why it exists**.

---

## Step 1: First Go Program (Go’s “Hello, World” + structure)

### 1. Create a folder

Pick a folder for this experiment, for example:

```powershell
mkdir C:\code\go-basics
cd C:\code\go-basics
```

### 2. Create `main.go`

Create a file named `main.go` with the following content:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello from Go!")
}
```

The sections below unpack **every single piece**.

---

## File structure & rules

### `package main`

```go
package main
```

* In Go, **every file belongs to a package**.
* A **package** is similar to a namespace or module in Java.
* `main` is **special**:

  * If a program has a `package main` and a function `main`, Go treats it as:

    > “An executable program. A runnable binary will be built.”

So:

* `package main` → the file is part of a directly runnable program.
* Comparison with Java:

  * Java requires a class with `public static void main(String[] args)`.
  * Go requires a **package named `main`** and a **function named `main`**.

Without `package main`, no runnable binary is produced.

---

### `import "fmt"`

```go
import "fmt"
```

* `import` is similar to `import java.util.*;` in Java.
* `"fmt"` is the name of a **standard library package**.
* `fmt` stands for **format** (input/output formatting).
* It contains functions such as:

  * `fmt.Println` → print with newline
  * `fmt.Printf` → formatted print (`printf`-style)

The `"fmt"` package is imported because `fmt.Println` is used later.

Important Go behavior:

* Importing something and not using it causes a compiler error.
* The same rule applies to variables—unused variables are not allowed.
* This enforces **no dead code** by default.

Since `fmt.Println` is used, importing `"fmt"` is required and valid.

---

### `func main() { ... }`

```go
func main() {
	fmt.Println("Hello from Go!")
}
```

Breakdown:

#### `func`

* Keyword used to define a **function** in Go.
* Comparable to method definitions in Java, with different syntax.

#### `main`

* Name of the function.
* This exact name is significant:

  * Inside `package main`, `func main()` acts as the **entry point**, similar to `public static void main` in Java.
* Entry-point signature rules:

  * No parameters
  * No return values
  * Defined strictly as: `func main() { ... }`

#### `()`

* Empty parentheses indicate **no arguments**.

#### `{ ... }`

* Curly braces define the function body.
* Go enforces a strict formatting rule:

  * `{` **must** be on the same line as `func main()`.

  * The following does not compile:

    ```go
    func main()
    {
        // not allowed
    }
    ```

  * This enforces consistent formatting across Go code.

---

### `fmt.Println("Hello from Go!")`

```go
fmt.Println("Hello from Go!")
```

* `fmt` → imported package.
* `.` → access a member of the package.
* `Println` → function inside the `fmt` package.

Key details:

* Functions starting with a **capital letter** (e.g., `Println`) are **exported**:

  * Accessible outside the package.
  * Comparable to `public` methods in Java.
* Lowercase names would be package-private.

Behavior of `Println`:

* Prints arguments to **standard output**.
* Appends a newline.
* Equivalent to `System.out.println` in Java.

---

## Step 2: Running the program

From the directory containing `main.go`:

### Option 1: Run directly (similar to `java Main.java`)

```bash
go run main.go
```

What happens:

* Code is compiled into a **temporary binary**.
* The binary is executed immediately.
* The binary is deleted after execution.

Expected output:

```bash
Hello from Go!
```

---

### Option 2: Build a binary (similar to `javac` + `java`)

```bash
go build
```

* Compiles all `.go` files in the current directory.
* Produces an executable:

  * On Windows: `go-basics.exe` (or `main.exe`, depending on folder name).
* The executable can then be run:

```bash
.\go-basics.exe
```

This is comparable to:

* `javac Main.java` → generates `Main.class`
* `java Main` → runs it

Go produces a **single native binary**, without requiring a JVM at runtime.

---

## Tiny extension (optional, Java-style approach)

In Java, a common first step is defining a variable and printing it.
The same concept can be applied in Go.

Update `main.go`:

```go
package main

import "fmt"

func main() {
	message := "Hello from Go variables!"
	fmt.Println(message)
}
```

Explanation of new elements:

---

### `message := "Hello from Go variables!"`

* `message` → variable name.
* `:=` → **short variable declaration**:

  * Declares and assigns simultaneously.
  * Type is inferred from the right-hand side.
  * In this case, the inferred type is `string`.

Comparable Java code:

```java
String message = "Hello from Go variables!";
```

Go infers `string` automatically in this context.

---

### `fmt.Println(message)`

* Prints the **value stored in the variable**.
* Since `message` is used, no unused-variable error occurs.
* Since `fmt` is used, the import remains valid.

Run again:

```bash
go run main.go
```

Expected output:

```bash
Hello from Go variables!
```

---

## What has been established so far

A working mental model now exists for:

* **Packages**: Every file starts with `package <name>`.
* **Executable vs library code**:

  * `package main` + `func main()` → runnable program.
  * Other packages (e.g., `package mylib`) → reusable library code.
* **Imports**:

  * Always use string-based imports: `import "fmt"`.
  * Unused imports result in compile-time errors.
* **Basic function structure in Go**:

  ```go
  func name(params) returnType {
      // body
  }
  ```

  (Parameters and return types are explored later.)
* **Standard output and libraries**: `fmt.Println` as the Go equivalent of `System.out.println`.
