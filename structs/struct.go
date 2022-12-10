package main

import "fmt"

func main() {
	type courseMeta struct {
		author string
		level  string
		rating float64
	}

	//var pythonBasics courseMeta
	//gettingStartedWithGo := new(courseMeta)

	gettingStartedWithK8s := courseMeta{
		author: "Nigel Poulton",
		level:  "Intermediate",
		rating: 5,
	}

	fmt.Println(gettingStartedWithK8s.author)
	fmt.Println(gettingStartedWithK8s.level)
	fmt.Println(gettingStartedWithK8s.rating)

	gettingStartedWithK8s.rating = 1.2
	fmt.Println(gettingStartedWithK8s.rating)
}
