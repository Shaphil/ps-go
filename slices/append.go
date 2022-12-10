package main

import "fmt"

func main() {
	nums := make([]int, 1, 4)
	fmt.Printf("Length of the slice is %d and capacity is %d\n", len(nums), cap(nums))

	for i := 1; i < 17; i++ {
		nums = append(nums, i)
		fmt.Printf("Length of the slice is %d and capacity is %d\n", len(nums), cap(nums))
	}
}
