## Go — Errors, `defer`, `panic`, `recover`
---

## 1. New folder and file

Example on Windows:

```bash
mkdir C:\code\go-errors
cd C:\code\go-errors
```

Create `main.go` with the following content.

---

## 2. Program: error-return pattern, defer, panic/recover

```go
package main

import (
	"errors"
	"fmt"
)

// 1. Function that can fail, returns (result, error)
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	result := a / b
	return result, nil
}

// 2. Function that logs start/end using defer
func doWork(name string) {
	fmt.Println("start work:", name)
	defer fmt.Println("end work:", name)

	fmt.Println("doing something in", name)
}

// 3. Function that panics, and a wrapper that recovers
func mightPanic(trigger bool) {
	if trigger {
		panic("something went really wrong")
	}
	fmt.Println("mightPanic completed without panic")
}

func safeCall(trigger bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()
	fmt.Println("safeCall: before mightPanic")
	mightPanic(trigger)
	fmt.Println("safeCall: after mightPanic")
}

func main() {
	// A. Normal error handling with (value, error)
	fmt.Println("== divide with valid input ==")
	value, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", value)
	}

	fmt.Println("\n== divide with division by zero ==")
	value, err = divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", value)
	}

	// B. defer for cleanup-style logic
	fmt.Println("\n== doWork with defer example ==")
	doWork("task-1")

	// C. panic and recover
	fmt.Println("\n== safeCall with trigger = false (no panic) ==")
	safeCall(false)

	fmt.Println("\n== safeCall with trigger = true (panic + recover) ==")
	safeCall(true)

	fmt.Println("\nProgram completed after safeCall(true)")
}
```

Run:

```bash
go run main.go
```

---

## 3. Explanation of each concept

---

### Package and imports

```go
package main
```

* Declares the `main` package.
* A `main` package with a `main` function builds as an executable.

```go
import (
	"errors"
	"fmt"
)
```

* `fmt`:

  * For printing and formatting.
* `errors`:

  * Standard library package for creating `error` values via `errors.New`.

Go requires every imported package to be used somewhere in the file.

---

### 1. Function that can fail: `(value, error)` pattern

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	result := a / b
	return result, nil
}
```

Key ideas:

* Signature:

  * Parameters: `a, b float64`
  * Return values: `(float64, error)`

    * First: result of division.
    * Second: an `error` value describing failure, or `nil` on success.
* `if b == 0`:

  * Guards against invalid operation.
* `errors.New("cannot divide by zero")`:

  * Constructs a new `error` with the given message.
* `return 0, errors.New(...)`:

  * Returns a placeholder numeric result and a non-nil error.
  * Callers are expected to check the error first.
* On success:

  ```go
  result := a / b
  return result, nil
  ```

  * `nil` means “no error”.

This pattern is the core of Go error handling:

> Return `(value, error)` instead of throwing exceptions.

---

### `main` using the error-return pattern

```go
fmt.Println("== divide with valid input ==")
value, err := divide(10, 2)
if err != nil {
	fmt.Println("error:", err)
} else {
	fmt.Println("result:", value)
}
```

* Calls `divide(10, 2)`.
* Multiple assignment:

  * `value` holds the numeric result.
  * `err` holds the error.
* Check:

  * `if err != nil`:

    * Non-nil error means failure.
  * Else branch runs only when `err == nil`.

Second case:

```go
fmt.Println("\n== divide with division by zero ==")
value, err = divide(10, 0)
if err != nil {
	fmt.Println("error:", err)
} else {
	fmt.Println("result:", value)
}
```

* Calls `divide(10, 0)`:

  * This triggers the `b == 0` branch in `divide`.
  * Returns `0` and a non-nil error.
* Error branch prints the error message.

This is the standard, idiomatic Go style: **explicit error checking at call sites**.

---

### 2. `defer` for cleanup-style logic

```go
func doWork(name string) {
	fmt.Println("start work:", name)
	defer fmt.Println("end work:", name)

	fmt.Println("doing something in", name)
}
```

Behavior:

* `fmt.Println("start work:", name)`:

  * Runs immediately when `doWork` is called.
* `defer fmt.Println("end work:", name)`:

  * Schedules this `Println` to run **when `doWork` returns**.
  * Deferred calls execute **after** normal execution reaches the end of the function, or after a `return` is hit, or after a panic unwinds into this stack frame.
* `fmt.Println("doing something in", name)`:

  * Regular work in the middle.

Call site:

```go
fmt.Println("\n== doWork with defer example ==")
doWork("task-1")
```

Expected order:

1. `start work: task-1`
2. `doing something in task-1`
3. `end work: task-1` (deferred call)

`defer` is commonly used for:

* Closing files:

  * `defer file.Close()`
* Releasing locks:

  * `defer mu.Unlock()`
* Logging entry/exit of functions.





### `defer` in Go?

In Go, the `defer` keyword is used to schedule a function (or statement) to be **executed** right **before** the current function **returns**, no matter what happens (like whether an error occurs or not). It's like a **cleanup** mechanism.

---

### How does `defer` work?

* **When is it executed?**
  The deferred function or statement is **not** executed immediately. It is executed **last** when the function returns.

* **Why use it?**
  It's perfect for things like:

  * Closing files
  * Unlocking mutexes
  * Releasing resources (database connections, network connections)
  * Logging
  * Handling cleanup after a task is done, even if an error occurs.

---

### Go `defer` vs Java `finally`


In Java, we might use the `finally` block to ensure that a certain piece of code runs after a `try-catch` block, no matter if an exception occurs or not.

#### Example in Java:

```java
try {
    // Some code
} catch (Exception e) {
    // Handle exception
} finally {
    // This block will always run, whether an exception occurred or not
    // It’s like cleanup code
    closeResources();
}
```

In Go, `defer` behaves similarly to the `finally` block in Java, except:

* You don't need to explicitly declare a `finally` block.
* `defer` is placed directly before the function exits, and it will always run no matter how the function exits (even if it exits early due to a `return`, `panic`, etc.).

---

### Go Example using `defer`:

Here’s a Go example to show how `defer` works:

```go
package main

import "fmt"

func exampleFunction() {
	// Defer statement
	defer fmt.Println("This is executed last, before the function returns.")

	// Some other code
	fmt.Println("This is executed first.")
}

func main() {
	exampleFunction()
}
```

#### Output:

```
This is executed first.
This is executed last, before the function returns.
```

### How does it compare to Java?

* In Go, `defer` will run just before the function exits, which is conceptually similar to how `finally` in Java ensures certain cleanup happens at the end of a method, regardless of whether an exception was thrown.

#### Defer with multiple statements:

In Go, you can have multiple `defer` statements, and they are executed in **LIFO (Last In, First Out)** order. This means the last `defer` you add will run first.

Example:

```go
package main

import "fmt"

func exampleFunction() {
	defer fmt.Println("First deferred statement.")
	defer fmt.Println("Second deferred statement.")
	defer fmt.Println("Third deferred statement.")

	fmt.Println("Function starts here.")
}

func main() {
	exampleFunction()
}
```

#### Output:

```
Function starts here.
Third deferred statement.
Second deferred statement.
First deferred statement.
```

### When to use `defer`?

In Go, you would typically use `defer` for things like:

* Closing files
* Unlocking mutexes
* Closing database connections
* Releasing resources

For example, when working with files in Go:

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Open a file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Defer closing the file
	defer file.Close()

	// Do something with the file
	fmt.Println("File opened successfully")
}
```

In this case, `file.Close()` is deferred so that it is guaranteed to run when the `main` function exits, ensuring the file is closed even if an error occurs earlier in the function.

---

### Key Takeaways:

1. **Defer in Go** is like **finally** in Java, ensuring that a function or resource is properly cleaned up before the function returns.
2. **Deferred functions** are executed in reverse order (LIFO).
3. It is typically used for things like closing resources (files, network connections, etc.), which need to be closed regardless of the function’s execution flow.


---

### 3. `panic` and `recover`


### Understanding `panic` in Go

In Go, **`panic`** is used to handle unexpected situations that result in an unrecoverable error, causing the program to stop execution. Unlike normal error handling with return values (like using `error` types), `panic` is intended for situations where continuing execution would lead to undefined behavior or an invalid program state. It immediately terminates the normal flow of execution and starts unwinding the stack.

#### Why Not Always Use `panic` for Errors?

In Go, functions commonly return an `error` type alongside the result to indicate any issues or failures. This allows the calling code to handle errors in a controlled manner without halting the program. For example, when attempting to divide two numbers, we can check for errors like division by zero and handle them gracefully.

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	result := a / b
	return result, nil
}
```

In this case, if there's an issue (like dividing by zero), the function returns an error, and the caller can decide how to handle it:

```go
value, err := divide(10, 0)
if err != nil {
    fmt.Println("error:", err)
} else {
    fmt.Println("result:", value)
}
```

This method of handling errors is **preferred** for regular, recoverable errors because it gives you control over the program flow, allowing the program to continue running or take appropriate actions.

#### When to Use `panic`

`panic` is meant for **critical errors** that make it impossible for the program to continue in a meaningful way. For example, trying to open a critical file at the beginning of your program and failing to do so would be an appropriate case for using `panic`.

Consider the following example where `panic` is used to handle a missing or unreadable configuration file:

```go
package main

import (
	"fmt"
	"os"
)

// openConfigFile tries to open the config file. If it fails, it panics with a message.
// Note that when panic is called, the function does not return normally; it immediately stops execution.
func openConfigFile(filename string) *os.File {
	// Try to open the configuration file
	file, err := os.Open(filename)
	if err != nil {
		// If there's an error, panic with a message.
		// This stops the function execution immediately and does not return a value.
		panic(fmt.Sprintf("Critical error: cannot open config file '%s': %v", filename, err))
	}
	// If successful, return the file pointer (this won't be reached if panic is called).
	return file
}

func main() {
	// Specify the configuration file
	configFile := "config.txt"
	
	// Attempt to open the config file. If panic occurs, the program will terminate here.
	file := openConfigFile(configFile)
	// Ensure the file gets closed when the program exits
	defer file.Close()

	// Proceed with the rest of the program, assuming the file is successfully opened
	fmt.Println("Config file opened successfully!")
	// Additional logic like reading from the file could go here
}
```

### Explanation:

1. **Function `openConfigFile`**:

   * This function attempts to open a specified configuration file. If the file can't be opened (e.g., it’s missing or corrupted), the function calls `panic` to stop the program. The `panic` message provides a clear explanation of the error.
   * **Note**: When `panic` is called, the function doesn't return normally. Execution halts immediately, and the program begins the process of unwinding the stack.

2. **Usage in `main`**:

   * If the file can't be opened, the program will panic and print an error message.
   * If the file is successfully opened, the program proceeds, and the file is closed using `defer` to ensure cleanup.

### Key Points to Remember:

* **`panic` Stops Execution**: When `panic` is called, the function doesn't return. Instead, it immediately terminates the execution of the function and starts unwinding the call stack.
* **No Return After `panic`**: Since `panic` halts the execution, there's no need to return any value after calling `panic`. The program will terminate, and any deferred functions will still be executed.
* **Use `panic` for Critical Failures**: `panic` should only be used for errors that are critical and make it impossible for the program to continue (e.g., missing essential files or resources).
* **Graceful Error Handling**: For normal recoverable errors (like user input validation or non-critical file access), it's better to return an error and handle it gracefully rather than panicking.

### Example Output:

If the config file (`config.txt`) is missing, the output will look like this:

```bash
panic: Critical error: cannot open config file 'config.txt': open config.txt: no such file or directory

goroutine 1 [running]:
main.openConfigFile()
        /path/to/your/file.go:12 +0x12c
main.main()
        /path/to/your/file.go:20 +0x6f
```

This output shows the panic message along with a stack trace, indicating where the error occurred.

### Conclusion

While **`panic`** is a powerful tool for handling critical errors, it should be used sparingly and only for situations where continuing execution doesn't make sense. For most cases, Go’s idiomatic error handling (returning an error) is a better approach, as it allows your program to handle issues more gracefully.


### **`recover` in Go**

In Go, **`recover`** is used in conjunction with **`panic`** to handle a panic situation and **recover** from it, preventing the program from terminating. It is called within a **deferred** function, meaning that you can use it to handle panics after they have been triggered, giving you the opportunity to **recover gracefully**.

Think of **`recover`** as a way to "catch" a panic and allow the program to continue running instead of crashing. It can be compared to **Java's `try-catch` blocks** for handling exceptions, but it has some key differences.

### Key Concepts:

1. **`panic`**: Causes the program to stop executing and starts unwinding the call stack.
2. **`recover`**: Allows you to handle a panic if it occurs and **"recover"** from it, enabling the program to continue executing.

However, unlike Java's `try-catch` block, **Go does not have a general `try-catch` mechanism**. Instead, Go uses **`panic` and `recover`** to deal with errors that are critical or unexpected, often in low-level parts of the code. **`recover`** only works in a **deferred function**.

### Example with `recover`:
.

#### Example: Handling `panic` with `recover`

```go
package main

import (
	"fmt"
	"os"
)

// openConfigFile tries to open the config file. If it fails, it panics with a message.
func openConfigFile(filename string) *os.File {
	// Try to open the configuration file
	file, err := os.Open(filename)
	if err != nil {
		// If there's an error, panic with a message.
		// This stops the function execution immediately and does not return a value.
		panic(fmt.Sprintf("Critical error: cannot open config file '%s': %v", filename, err))
	}
	// If successful, return the file pointer (this won't be reached if panic is called).
	return file
}

// safeOpenConfigFile tries to open a config file and recovers from panic if it occurs.
func safeOpenConfigFile(filename string) (file *os.File, errMsg string) {
	defer func() {
		// If a panic occurs, recover will capture it.
		if r := recover(); r != nil {
			// Handle the panic (recover from it)
			errMsg = fmt.Sprintf("Recovered from panic: %v", r)
		}
	}()
	// Attempt to open the file (this could panic)
	file = openConfigFile(filename)
	return file, ""
}

func main() {
	// Specify the configuration file
	configFile := "config.txt"
	
	// Attempt to open the config file with recovery in place
	file, errMsg := safeOpenConfigFile(configFile)
	if errMsg != "" {
		// Panic was recovered, print the error message
		fmt.Println(errMsg)
	} else {
		// If file opened successfully, proceed as normal
		defer file.Close()
		fmt.Println("Config file opened successfully!")
	}
}
```

### Explanation:

* **`openConfigFile(filename)`**: This function attempts to open a file. If it fails, it triggers a **`panic`** with a specific error message.
* **`safeOpenConfigFile(filename)`**: This function wraps the `openConfigFile` call and uses **`defer`** and **`recover`** to handle the panic. If `panic` is triggered, **`recover`** catches the panic, preventing the program from crashing, and the error message is returned instead.
* **In `main()`**: We check if an error message was returned from `safeOpenConfigFile`. If `errMsg` is non-empty, we know that a panic occurred, and we print the recovery message. Otherwise, we proceed normally.

### Output:

If the config file (`config.txt`) is missing or can't be opened, instead of the program terminating, it will output something like:

```bash
Recovered from panic: Critical error: cannot open config file 'config.txt': open config.txt: no such file or directory
```

If the file is opened successfully, it will output:

```bash
Config file opened successfully!
```

### **`recover` vs Java `try-catch`**

In Java, you would use **`try-catch`** blocks to catch exceptions and handle them gracefully:

#### Example in Java:

```java
import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        try {
            // Try to open the file
            Scanner scanner = new Scanner(new File("config.txt"));
            System.out.println("File opened successfully!");
            scanner.close();
        } catch (FileNotFoundException e) {
            // Handle the exception
            System.out.println("Error: " + e.getMessage());
        }
    }
}
```

In this Java example:

* **`try-catch`** blocks are used to handle errors like missing files.
* If the file isn't found, the program catches the `FileNotFoundException` and prints an error message, allowing the program to continue running.

### Key Differences between Go’s `panic-recover` and Java's `try-catch`:

1. **Mechanism**:

   * **Java**: You use **`try-catch`** blocks for exceptions, which are a form of structured error handling. The program flow is explicitly interrupted when an exception occurs, but you can catch and handle exceptions at any point.
   * **Go**: Go doesn't have a `try-catch` mechanism. Instead, **`panic`** is used for unrecoverable errors, and **`recover`** is used in a **deferred function** to catch and handle those panics. It’s more like a way to catch a critical error and allow the program to continue.

2. **Use Case**:

   * **Java**: `try-catch` is used for general exception handling. It can catch a variety of exceptions (e.g., `IOException`, `FileNotFoundException`) and allows you to handle different types of exceptions based on their types.
   * **Go**: `panic-recover` is used specifically for critical errors that cannot be recovered from in the normal flow of the program (e.g., a missing file that the program cannot operate without). It's not meant for regular error handling.

3. **Error Handling Philosophy**:

   * **Go**: Go prefers handling errors via return values (i.e., `error` type) rather than exceptions or panics. **Panic-recover** is meant to be used for serious problems, not regular error handling.
   * **Java**: Java's approach is more reliant on exceptions for handling errors in the flow, and it uses `try-catch` blocks for both expected and unexpected errors.

### Summary of **`panic` and `recover`** in Go:

* **`panic`** is used for critical errors where the program cannot continue. It stops execution and unwinds the stack.
* **`recover`** is used to catch and handle a panic inside a deferred function, allowing the program to continue executing instead of crashing.
* **In Java**, we would use `try-catch` to handle exceptions, allowing for controlled error handling and recovery.

### Conclusion:

While Go doesn't use exceptions in the same way as Java, **`panic` and `recover`** provide a similar mechanism to catch and handle critical errors. However, in Go, you typically handle regular errors using the `error` type and return values, reserving `panic` for truly exceptional conditions.

###  What Happens in `safeOpenConfigFile`:

Here’s the structure of the `safeOpenConfigFile` function that we discussed earlier:

```go
func safeOpenConfigFile(filename string) (file *os.File, errMsg string) {
    defer func() {
        // If a panic occurs, recover will capture it.
        if r := recover(); r != nil {
            // Handle the panic (recover from it)
            errMsg = fmt.Sprintf("Recovered from panic: %v", r)
        }
    }()
    
    // Attempt to open the file (this could panic)
    file = openConfigFile(filename)
    return file, ""
}
```


### Flow of Execution in `safeOpenConfigFile`

1. **Call to `safeOpenConfigFile("config.txt")`**:

   * The function is called to open the file `"config.txt"`. We assume the file doesn't exist (which will cause a **`panic`**).
   * As soon as **`safeOpenConfigFile`** is invoked, **the `defer` statement** is set up to ensure that the deferred function is run when the function finishes executing.

2. **Defer and `recover`**:

   * The deferred function contains a **`recover`** call.
   * **`defer`** ensures that the deferred function will run **when the current function exits**, even if the function exits due to a **`panic`**.

3. **Execution of `openConfigFile`**:

   * Inside `safeOpenConfigFile`, the function `openConfigFile` is called to open the file `"config.txt"`.
   * **If the file doesn't exist**, `openConfigFile` will **`panic`** with a message like:

     ```bash
     panic: Critical error: cannot open config file 'config.txt': open config.txt: no such file or directory
     ```

4. **Panic is Triggered**:

   * As soon as **`panic`** is called, the normal execution of `openConfigFile` is immediately halted.
   * **The stack starts unwinding**, and Go will attempt to run any **deferred functions** in reverse order (LIFO—Last In, First Out).

5. **Deferred Function Executes**:

   * The deferred function in `safeOpenConfigFile` is executed as part of the stack unwinding process.
   * **`recover`** is called inside the deferred function. This is where **`recover`** catches the **`panic`** and prevents the program from terminating.
   * The **panic message** is captured by `recover`, and we assign that message to `errMsg`.

6. **Function Exit with `errMsg`**:

   * After the deferred function executes, **`safeOpenConfigFile`** exits, and the values `file` and `errMsg` are returned.
   * Since the panic occurred, `file` is not assigned (it remains `nil`), and `errMsg` will contain the recovery message, like:

     ```bash
     Recovered from panic: Critical error: cannot open config file 'config.txt': open config.txt: no such file or directory
     ```

7. **Back in `main()`**:

   * The caller (`main()` in this case) checks if `errMsg` is non-empty. If so, it prints the error message, showing that the panic was recovered.
   * If no panic had occurred, it would proceed to use the `file` object normally.

### Step-by-Step Example with Comments:

```go
package main

import (
	"fmt"
	"os"
)

// openConfigFile tries to open the config file. If it fails, it panics with a message.
func openConfigFile(filename string) *os.File {
	// Try to open the configuration file
	file, err := os.Open(filename)
	if err != nil {
		// If there's an error, panic with a message.
		panic(fmt.Sprintf("Critical error: cannot open config file '%s': %v", filename, err))
	}
	// If successful, return the file pointer (this won't be reached if panic is called).
	return file
}

// safeOpenConfigFile tries to open a config file and recovers from panic if it occurs.
func safeOpenConfigFile(filename string) (file *os.File, errMsg string) {
	// Defer function will be executed when safeOpenConfigFile exits, even if it's due to a panic.
	defer func() {
		// If a panic occurs, recover will capture it.
		if r := recover(); r != nil {
			// Handle the panic (recover from it)
			errMsg = fmt.Sprintf("Recovered from panic: %v", r)
		}
	}()
	
	// Attempt to open the file (this could panic)
	file = openConfigFile(filename)
	return file, ""
}

func main() {
	// Specify the configuration file
	configFile := "config.txt"
	
	// Attempt to open the config file with recovery in place
	file, errMsg := safeOpenConfigFile(configFile)
	if errMsg != "" {
		// Panic was recovered, print the error message
		fmt.Println(errMsg)
	} else {
		// If file opened successfully, proceed as normal
		defer file.Close()
		fmt.Println("Config file opened successfully!")
	}
}
```

### Flow Breakdown When **Panic** Occurs:

1. **`main()` calls `safeOpenConfigFile("config.txt")`**.
2. Inside `safeOpenConfigFile`, the **`defer`** is set up to ensure that **`recover()`** can catch any panic if it happens.
3. **`openConfigFile` is called** inside `safeOpenConfigFile` to attempt opening the file. If the file doesn't exist, **`panic`** is triggered.
4. The program **panics**, but before it completely exits, the deferred function is called, and **`recover()`** handles the panic by capturing it.
5. The **`recover`** function sets **`errMsg`**, which will be returned to `main()`.
6. In **`main()`**, **`errMsg`** is checked. Since there was an error (a panic was recovered), **`errMsg`** will contain the recovery message.
7. **`main()` prints the error message**: "Recovered from panic: Critical error: cannot open config file 'config.txt': open config.txt: no such file or directory".

### What Happens if There Was No Panic?

* If the file exists and is successfully opened, **`openConfigFile`** will return a valid file object.
* The deferred function will still execute, but since no panic occurred, **`recover()`** will return `nil`, and `errMsg` will remain empty.
* The program will then print `"Config file opened successfully!"`.

### Summary:

* **`defer`**: Ensures that certain cleanup actions (like closing files) happen when the function exits, regardless of whether the function exits normally or due to a panic.
* **`recover`**: Catches a panic inside a **deferred function** and allows the program to recover from it, preventing the program from crashing.
* In **`safeOpenConfigFile`**, if **`panic`** occurs, the program won’t crash because **`recover`** is handling it. Instead, the program returns an error message in `errMsg`, and the program flow continues.

This is a simple, yet powerful combination of **`defer`** and **`recover`** that allows you to **manage critical failures** without bringing down the entire program.


---

## 4. When to use each: mental model

* `(T, error)` return pattern:

  * For **expected failures**:

    * Invalid input.
    * Network issues.
    * File-not-found.
  * Caller decides how to handle the error.
* `defer`:

  * For **cleanup logic**:

    * Closing resources.
    * Unlocking mutexes.
    * Logging function exit.
* `panic`:

  * For **truly unrecoverable states**:

    * Invariants broken.
    * Impossible situations (logic bugs).
  * Preferable not to use for normal flow control.
* `recover`:

  * For **boundary layers**:

    * HTTP handler wrappers.
    * Goroutine entry points.
  * Converts panics into logs or safe shutdown behavior.

This trio—`error`, `defer`, `panic/recover`—forms the Go replacement for exception stacks, `try/finally`, and similar constructs in languages like Java.

