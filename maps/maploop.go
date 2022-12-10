package main

import "fmt"

func main() {
	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
	}

	for key, val := range testMap {
		fmt.Printf("Key is: %v, Value is: %v\n", key, val)
	}

	testMap["A"] = 100
	fmt.Println(testMap)

	delete(testMap, "C")
	fmt.Println(testMap)
}
