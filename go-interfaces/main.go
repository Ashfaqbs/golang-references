package main

import "fmt"

// 1. Interface definition (behavior contract)
type Describer interface {
	Describe() string
}

// 2. Concrete type 1
type Person struct {
	Name string
	Age  int
}

// Person implements Describer with a value receiver method
func (p Person) Describe() string {
	return fmt.Sprintf("Person(Name=%s, Age=%d)", p.Name, p.Age)
}

// 3. Concrete type 2
type Product struct {
	Name  string
	Price float64
}

// Product implements Describer with a value receiver method
func (p Product) Describe() string {
	return fmt.Sprintf("Product(Name=%s, Price=%.2f)", p.Name, p.Price)
}

// 4. Function that works with any Describer
func PrintDescription(d Describer) {
	fmt.Println("Description:", d.Describe())
}

func main() {
	// 5. Create concrete values
	person := Person{
		Name: "Alice",
		Age:  30,
	}
	product := Product{
		Name:  "Laptop",
		Price: 1299.99,
	}

	fmt.Println("Concrete values:")
	fmt.Println("person:", person)
	fmt.Println("product:", product)

	// 6. Assign concrete values to interface variable
	var d Describer

	d = person
	fmt.Println("\nDescriber holding Person:")
	PrintDescription(d)

	d = product
	fmt.Println("\nDescriber holding Product:")
	PrintDescription(d)

	// 7. Slice of Describer (polymorphic collection)
	items := []Describer{
		person,
		product,
		Product{Name: "Phone", Price: 699.50},
		Person{Name: "Bob", Age: 25},
	}

	fmt.Println("\nLoop over slice of Describer:")
	for index, item := range items {
		fmt.Println("index:", index)
		PrintDescription(item)
	}

	// 8. Zero value (nil) of an interface
	var empty Describer
	fmt.Println("\nZero value of interface (empty):", empty)
	fmt.Println("Is empty == nil?", empty == nil)
}
