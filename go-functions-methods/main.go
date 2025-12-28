package main

import "fmt"

/*

Regular Functions: Objects in Go cannot call regular functions directly. You need to pass the object as an argument to the function if you want it to operate on that object. Functions are not inherently tied to types or objects.

Methods: When a function is defined as a method with a receiver, it belongs to the type (like Person). An object of that type can call the method because the method is associated with the object type.
*/

// 1. Struct type reused from previous step
type Person struct {
	Name string
	Age  int
}

// 2. Simple function with one parameter, no return value
func greet(name string) {
	fmt.Println("Hello,", name)
}

// 3. Function with parameters and a single return value
func add(a int, b int) int {
	sum := a + b
	return sum
}

// 4. Function with multiple return values, can return n returns
func divideAndRemainder(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// 5. Method with value receiver (does not modify original)
func (p Person) Info() {
	p.Age = 123
	fmt.Println("Person Info -> Name:", p.Name, "Age:", p.Age)
}

// 6. Method with pointer receiver (can modify original)
func (p *Person) HaveBirthday() {
	p.Age = p.Age + 1
	fmt.Println("Happy birthday,", p.Name, "New age:", p.Age)
}

func main() {
	// Call simple function
	fmt.Println("Calling greet:")
	greet("Go learner")

	// Call function with return value
	fmt.Println("\nCalling add:")
	result := add(10, 20)
	fmt.Println("add(10, 20):", result)

	// Call function with multiple return values
	fmt.Println("\nCalling divideAndRemainder:")
	q, r := divideAndRemainder(17, 5)
	fmt.Println("divideAndRemainder(17, 5) -> quotient:", q, "remainder:", r)

	// Create a Person value
	fmt.Println("\nCreating Person:")
	person := Person{
		Name: "Alice",
		Age:  29,
	}
	fmt.Println("Initial person:", person)

	// Call method with value receiver
	fmt.Println("\nCalling person.Info():")
	person.Info() // doesnot modify the actual object
	fmt.Println("Original person:", person)
	fmt.Println()

	// Modifies the actual object
	// Call method with pointer receiver (modifies Age)
	fmt.Println("\nCalling person.HaveBirthday():")
	person.HaveBirthday()
	fmt.Println("Person after HaveBirthday:", person)

	// Explicit pointer call (same effect)
	fmt.Println("\nCalling HaveBirthday via pointer:")
	personPtr := &person
	personPtr.HaveBirthday()
	fmt.Println("Person after pointer HaveBirthday:", person)
	// fmt.Println("Person after pointer HaveBirthday:", *personPtr)

}

/*


Here, we are defining the method Info() for the Person struct outside the struct itself. Even though the method is defined separately, you can call it on an instance of Person, because Go implicitly knows how to associate the method with the struct using the method's receiver (p Person).

Method Receiver: In Go, a method's receiver (p Person or p *Person) acts as a kind of context for the method, and the method is called on that type (in this case, Person).

Calling the Method: Even though Info is defined outside of Person, you can call it directly on an instance of Person because Go allows you to associate the method with that struct using the receiver.

In Go, this method invocation is not tied to a class in the same way Java methods are tied to a class. Instead, the method is just associated with the type (Person) via the receiver.

*/
