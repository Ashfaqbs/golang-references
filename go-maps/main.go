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
