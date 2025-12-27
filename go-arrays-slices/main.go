package main

import "fmt"

func main() {
	// 1. Fixed-size array
	var arr [4]int
	fmt.Println("arr (zero values):", arr) // will have default values in this arr or slice
	fmt.Println("len(arr):", len(arr))

	// 2. Array with initialization
	scores := [4]int{10, 20, 30, 40}
	fmt.Println("\nscores:", scores)
	fmt.Println("len(scores):", len(scores))

	// 3. Indexing and updating array elements
	fmt.Println("\nAccess elements of scores:")
	fmt.Println("scores[0]:", scores[0])
	fmt.Println("scores[1]:", scores[1])

	scores[2] = 35
	fmt.Println("scores after update (scores[2] = 35):", scores)

	// 4. Slice from array (view over part of the array)
	sub := scores[1:3]
	fmt.Println("\nsub slice (scores[1:3]):", sub)
	fmt.Println("len(sub):", len(sub))

	// 5. Slice literal (direct slice creation, backing array created automatically)
	nums := []int{100, 200, 300}
	fmt.Println("\nnums slice:", nums)
	fmt.Println("len(nums):", len(nums))

	// 6. Append to slice (dynamic growth)
	nums = append(nums, 400)
	fmt.Println("nums after append 400:", nums)

	nums = append(nums, 500, 600)
	fmt.Println("nums after append 500, 600:", nums)

	// 7. Make slice with make (length and capacity)
	data := make([]int, 3) // length 3, capacity 3
	fmt.Println("\ndata slice created with make:", data)
	fmt.Println("len(data):", len(data))

	data[0] = 7
	data[1] = 14
	data[2] = 21
	fmt.Println("data after assignments:", data)

	// 8. Append beyond initial length
	data = append(data, 28)
	fmt.Println("data after append 28:", data)
	fmt.Println("len(data):", len(data))

	// 9. Loop over slice with range
	fmt.Println("\nLoop over nums slice with range:")
	for index, value := range nums {
		fmt.Println("index:", index, "value:", value)
	}
}
