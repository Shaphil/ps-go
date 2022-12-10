package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USERNAME")
	course := "Go Fundamentals"

	fmt.Println("Hi", name, "your current course is:", course)
	updateCourse(&course)

	fmt.Println("Your current course is:", course)
}

func updateCourse(course *string) string {
	*course = "Python getting started"
	fmt.Println("Updated course to:", *course)
	return *course
}
