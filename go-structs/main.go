package main

import "fmt"

// 1. Define a struct type
type Person struct {
	Name   string
	Age    int
	Active bool
	skills [3]int
}

func main() {
	// 2. Zero value of a struct

	var p1 Person
	fmt.Println("p1 (zero value):", p1)
	fmt.Println("p1 fields -> Name:", p1.Name, "Age:", p1.Age, "Active:", p1.Active, " Skills: ", p1.skills)
	// p1.Name = "sample"
	// p1.skills = [3]int{1, 2, 3}

	// 3. Struct literal with field names
	p2 := Person{
		Name:   "Alice",
		Age:    30,
		Active: true,
	}
	fmt.Println("\np2 (with field names):", p2)

	// 4. Update struct fields
	p2.Age = 31
	p2.Active = false
	fmt.Println("p2 after updates:", p2)

	// 5. Struct literal without field names (positional)
	p3 := Person{"Bob", 25, true, [3]int{}}
	fmt.Println("\np3 (positional literal):", p3)

	// 6. Pointer to struct and modification via pointer

	/*
			In Go (and other languages that have pointers like C/C++), a pointer is essentially a variable that stores the memory address of another variable, instead of the actual value. So when you create a pointer to a variable, you’re just creating a reference to that variable’s memory location.
			A pointer is a variable that holds the address of another variable. Instead of storing the actual value (like an integer, string, or struct), it stores where that value is located in memory.
		Since both p4 and p4Ptr are referring to the same location in memory, modifying the values via the pointer (p4Ptr) automatically modifies the original variable (p4), because they are the same object in memory.
	*/
	p4 := Person{
		Name:   "Charlie",
		Age:    40,
		Active: true,
	}
	fmt.Println("\np4 before pointer modification:", p4)

	p4Ptr := &p4 //making p4Ptr a pointer to the memory address of p4, so both p4 and p4Ptr refer to the same memory location, using &.
	// & (Address-of operator): It gives you the memory address of a variable, essentially turning a variable into a pointer.
	p4Ptr.Age = 41 //When we change p4Ptr.Age = 41, it directly changes p4 because they are the same thing. This is the pointer behavior.
	p4Ptr.Active = false

	fmt.Println("p4 after pointer modification:", p4)
	fmt.Println("p4Ptr points to:", *p4Ptr)
	// * (Dereference operator): It allows you to access the value stored at the memory address pointed to by the pointer.
	fmt.Println()

	p5 := Person{
		Name:   "Name ",
		Age:    20,
		Active: true,
	}

	var pointerPerson Person = p5 // copy of the p5 data in pointerPerson variable so any change in pointerPerson wont affect the p5 var

	pointerPerson.Name = "Name Test"
	fmt.Println("p5 post modification : ", p5)
	fmt.Println("printing pointer person: ", pointerPerson)
	/*

		Above in Java terms :

						class Person {
				    String name;
				    int age;
				    boolean active;

				    // Constructor
		           // toString()
				}

				public class Main {
				    public static void main(String[] args) {
				        // Create a Person object
				        Person p4 = new Person("Charlie", 40, true);
				        System.out.println("\np4 before reference modification: " + p4);

				        // Reference to the Person object (similar to a pointer in Go)
				        Person p4Ref = p4; // This is a reference to p4, pointing to the same object in memory. it's just another reference pointing to the same object in memory. So, both p4Ref and p4 refer to the same instance of the Person object.

				        // Modify the object's fields through the reference
				        p4Ref.age = 41;
				        p4Ref.active = false;

				        // Output the object after modification
				        System.out.println("p4 after reference modification: " + p4);
				        System.out.println("p4Ref points to: " + p4Ref);
				    }
				}




	*/

	// 7. Slice of structs
	people := []Person{
		p2,
		p3,
		p4,
	}
	fmt.Println("\npeople slice:", people)

	// 8. Loop over slice of structs with range
	fmt.Println("\nLoop over people slice:")
	for index, person := range people {
		fmt.Println("index:", index, "Name:", person.Name, "Age:", person.Age, "Active:", person.Active)
	}
}
