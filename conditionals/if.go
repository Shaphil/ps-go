package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Day()

	if day > 25 {
		fmt.Println("There isn't much time left")
	} else if day == 15 {
		fmt.Println("Half the days are gone")
	} else if day > 15 && day <= 25 {
		fmt.Printf("Hurry now, only %d days left\n", 30-day)
	} else {
		fmt.Println("You still have time")
	}
}
