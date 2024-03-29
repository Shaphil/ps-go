package main

import "fmt"

func main() {
	coursesInProg := []string{
		"Docker & Kubernetes: The Big Picture",
		"Docker Networking",
		"Getting Started with Kubernetes",
		"Kubernetes Deep Dive",
	}

	coursesCompleted := []string{
		"Docker & Kubernetes: The Big Picture",
		"Docker Deep Dive",
	}

	for _, i := range coursesInProg {
		fmt.Println(i)
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("\nError: Clash -- Course:", i, "appears in both lists.")
			}
		}
	}
}
