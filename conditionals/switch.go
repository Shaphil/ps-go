package main

import (
	"fmt"
	"os"
)

func main() {
	host, _ := os.Hostname()

	switch host {
	case "Xenon":
		fmt.Println("The host is Light Blue")
	case "Radon":
		fmt.Println("The host is Colorless")
	case "Krypton":
		fmt.Println("The host is White")
	default:
		fmt.Println("Unidentified")
	}
}
