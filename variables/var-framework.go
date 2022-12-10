package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	name   = "55"
	module int
)

func main() {
	fmt.Println("Type of name", reflect.TypeOf(name))
	fmt.Println("Type of module", reflect.TypeOf(module))

	iModule, err := strconv.Atoi(name)
	if err == nil {
		fmt.Println(iModule)
	} else {
		fmt.Println(err)
	}

	fmt.Println("Memory address of `name`: ", &name)
	var ptr *string = &name
	fmt.Println("Address:", ptr, "Value: ", *ptr)
}
