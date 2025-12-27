package main

import "fmt"

func main() {
	// 1. Basic integer arithmetic
	a := 10
	b := 3

	sum := a + b
	diff := a - b
	prod := a * b
	quotient := a / b  // integer division
	remainder := a % b // modulus (remainder)

	fmt.Println("a:", a, " and value of b:", b)

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

	/*
			same equivalent code in java

			int score = 95;  // Example score
		int bonus = score - 80;  // Calculate bonus

		if (bonus > 0) {
		    System.out.println("bonus (score - 80): " + bonus);
		} else {
		    System.out.println("bonus (score - 80): 0 or negative");
		}


		note in java you cannot create a variable in if() and
		Go does not have a direct ternary operator, like other languages such as C, Java, or JavaScript.
	*/
}
