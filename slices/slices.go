package main

import "fmt"

func main() {
	courses := make([]string, 5, 10)
	fmt.Printf("Length of the slice is %d and capacity is %d\n", len(courses), cap(courses))

	var sl []int
	fmt.Println(sl)
}
