package main

import "fmt"

func main() {
	// 1. Classic counting for-loop
	fmt.Println("Classic for loop (0 to 4):")
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	// 2. While-style loop (condition-only for)
	fmt.Println("\nWhile-style loop (countdown from 3):")
	j := 3
	for j > 0 {
		fmt.Println("j:", j)
		j--
	}

	// 3. Infinite loop with break
	fmt.Println("\nInfinite loop with break (stop at k == 3):")
	k := 0
	for {
		fmt.Println("k:", k)
		if k == 3 {
			fmt.Println("breaking loop at k == 3")
			break
		}
		k++
	}

	// 4. Loop with continue (skip odd numbers)
	fmt.Println("\nLoop with continue (print only even numbers from 0 to 5):")
	for n := 0; n <= 5; n++ {
		if n%2 != 0 {
			continue
		}
		fmt.Println("even n:", n)
	}
	//A slice in Go is a more flexible, dynamic version of an array
	// 5. Loop over a slice with range
	fmt.Println("\nLoop over slice with range:")
	numbers := []int{10, 20, 30, 40}

	for index, value := range numbers {
		fmt.Println("index:", index, "value:", value)
	}

	// 6. Loop over slice with range, ignoring index
	fmt.Println("\nLoop over slice with range (values only):")
	for _, value := range numbers {
		fmt.Println("value:", value)
	}
}
