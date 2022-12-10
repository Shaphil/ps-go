package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func main() {
	author := "Shaphil"
	course := "Python Overview"

	fmt.Println(converter(author, course))
}

func converter(author string, course string) (a, c string) {
	title := cases.Title(language.English)
	author = strings.ToUpper(author)
	course = title.String(course)

	return author, course
}
