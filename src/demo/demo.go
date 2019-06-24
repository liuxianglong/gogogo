package main

import (
	//	"bytes"
	// "container/list"
	//	"encoding/gob"
	"fmt"
	"errors"
	//	"net/url"
)

func main() {
	
	z, err := div(10,0)
	if err != nil {
		panic(err)
	}
	fmt.Println(z)
}

func div(x,y int) (z int,err error){
	if y ==0 {
		err = errors.New("除数不能为0")
		return
	}
	z = x / y
	return
}