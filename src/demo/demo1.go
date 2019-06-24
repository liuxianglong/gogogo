package main

import (
	"fmt"
)

func main() {
	a := 1
	a++
	p := &a
	*p++
	fmt.Println(a)
}
