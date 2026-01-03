### Go — HTTP + JSON Basics (Step 13: mini service with in-memory data)

---

## 1. New folder and module

Windows example:

```bash
mkdir C:\code\go-http-json
cd C:\code\go-http-json
```

Initialize module:

```bash
go mod init example.com/go-http-json
```

---

## 2. `main.go`: simple HTTP server with JSON

Create `main.go` with the following content:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Person represents a simple data model for JSON input/output.
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// In-memory storage for Person records.
var (
	people []Person
	nextID = 1
)

// statusHandler returns a simple JSON status.
func statusHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{
		"status": "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		log.Println("error encoding status response:", err)
	}
}

// listPeopleHandler returns all people as JSON.
func listPeopleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println("error encoding people list:", err)
	}
}

// createPersonRequest represents the expected JSON body for creating a person.
type createPersonRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// createPersonHandler reads JSON body, creates a new Person, and returns it.
func createPersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req createPersonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Age <= 0 {
		http.Error(w, "name and age must be provided and valid", http.StatusBadRequest)
		return
	}

	person := Person{
		ID:   nextID,
		Name: req.Name,
		Age:  req.Age,
	}
	nextID++
	people = append(people, person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		log.Println("error encoding created person:", err)
	}
}

func main() {
	// Preload some in-memory data.
	people = append(people, Person{ID: nextID, Name: "Alice", Age: 30})
	nextID++
	people = append(people, Person{ID: nextID, Name: "Bob", Age: 25})
	nextID++

	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listPeopleHandler(w, r)
		case http.MethodPost:
			createPersonHandler(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := 8080
	fmt.Println("Starting server on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("server error:", err)
	}
}
```

Run the server:

```bash
go run main.go
```

Server listens on port `8080`.

Example HTTP calls (from another terminal):

* Status:

  ```bash
  curl http://localhost:8080/status
  ```

* List people (GET):

  ```bash
  curl http://localhost:8080/people
  ```

* Create person (POST):

  ```bash
  curl -X POST http://localhost:8080/people \
    -H "Content-Type: application/json" \
    -d '{"name": "Charlie", "age": 35}'
  ```

---

## 3. Explanation of every part

### Package and imports

```go
package main
```

* Declares executable package.

```go
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
```

* `net/http`:

  * Standard HTTP server and client library.
* `encoding/json`:

  * JSON encoding and decoding.
* `fmt`:

  * Basic formatted strings and printing.
* `log`:

  * Logging errors and server issues.

---

### Data model and storage

```go
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

* Struct represents an entity to send/receive via JSON.
* Struct tags such as `` `json:"id"` ``:

  * Define JSON field names.
  * `ID` → JSON key `"id"`, `Name` → `"name"`, `Age` → `"age"`.

```go
var (
	people []Person
	nextID = 1
)
```

* `people`:

  * In-memory slice of `Person`, acts as a fake database.
* `nextID`:

  * Simple counter to assign unique IDs when creating new records.

---

### `/status` handler

```go
func statusHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{
		"status": "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		log.Println("error encoding status response:", err)
	}
}
```

* `w http.ResponseWriter`:

  * Used to write response headers and body.
* `r *http.Request`:

  * Incoming request data (method, headers, body, URL).
* `status`:

  * Map representing JSON object `{"status":"ok"}`.
* `w.Header().Set("Content-Type", "application/json")`:

  * Sets HTTP `Content-Type` header so clients know JSON is being returned.
* `w.WriteHeader(http.StatusOK)`:

  * Sets HTTP status code `200`.
* `json.NewEncoder(w).Encode(status)`:

  * Serializes `status` to JSON and writes directly to response body.
* Error handling:

  * If encoding fails, error is logged.

---

### List people handler (`GET /people`)

```go
func listPeopleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println("error encoding people list:", err)
	}
}
```

* Method check:

  * Guards to allow only GET requests here.
* Response:

  * Returns the entire `people` slice as JSON array.
* `http.Error`:

  * Writes an error message and status code to the client if method is not allowed.

---

### Request struct for creating a person

```go
type createPersonRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

* Represents expected JSON request body for POST `/people`.
* Only `Name` and `Age` are provided by client; `ID` is generated server-side.

---

### Create person handler (`POST /people`)

```go
func createPersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req createPersonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
```

* Method check:

  * Allows only POST.
* `json.NewDecoder(r.Body).Decode(&req)`:

  * Reads JSON from request body and fills `req`.
* If decode fails, responds with `400 Bad Request`.

Validation:

```go
	if req.Name == "" || req.Age <= 0 {
		http.Error(w, "name and age must be provided and valid", http.StatusBadRequest)
		return
	}
```

* Simple sanity check on input fields.

Record creation:

```go
	person := Person{
		ID:   nextID,
		Name: req.Name,
		Age:  req.Age,
	}
	nextID++
	people = append(people, person)
```

* Builds a new `Person`:

  * `ID` generated by incrementing `nextID`.
  * `Name` and `Age` from request.
* Appends new person to `people` slice.

Response:

```go
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		log.Println("error encoding created person:", err)
	}
}
```

* Sets `Content-Type` to JSON.
* `http.StatusCreated` (`201`) indicates a resource was created.
* Encodes the created person back to the client as confirmation.

---

### `main` function: routing and server

Preloading data:

```go
people = append(people, Person{ID: nextID, Name: "Alice", Age: 30})
nextID++
people = append(people, Person{ID: nextID, Name: "Bob", Age: 25})
nextID++
```

* Adds two initial people to in-memory slice.

Handler registration:

```go
http.HandleFunc("/status", statusHandler)
http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listPeopleHandler(w, r)
	case http.MethodPost:
		createPersonHandler(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
})
```

* `http.HandleFunc(path, handlerFunc)`:

  * Registers handlers for specific paths.
* `/status`:

  * Uses `statusHandler`.
* `/people`:

  * Uses an inline function that switches on HTTP method:

    * `GET` → list people.
    * `POST` → create person.
    * Others → `405 Method Not Allowed`.

Server startup:

```go
port := 8080
fmt.Println("Starting server on port", port)
err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
if err != nil {
	log.Fatal("server error:", err)
}
```

* `http.ListenAndServe`:

  * Starts HTTP server on the given address.
  * Second argument `nil`:

    * Uses the default `http.DefaultServeMux` routing table, which was configured using `http.HandleFunc`.
* On error (for example, port already in use):

  * `log.Fatal` prints error and exits program.

---

## 4. What this step completes conceptually

* Basic HTTP server with Go’s standard library (`net/http`).
* JSON serialization and deserialization with `encoding/json`.
* Use of structs with JSON tags for clear API contracts.
* In-memory slice as a simple store for GET and POST operations.
* HTTP status codes and content-type headers.
* Separation between:

  * Data model (`Person`)
  * Request model (`createPersonRequest`)
  * Handlers (`statusHandler`, `listPeopleHandler`, `createPersonHandler`)
