package main

import (
	_"fmt"
)

func main() {
	x := 10
	var p *int = &x
	*p += 20
	var p2 *int = &x
	*p2++
	println(x, *p, *p2)
	//返回值31 31 31
}
