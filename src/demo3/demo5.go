package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := errors.New("exception")
	if err != nil {
		fmt.Println(err)
	}
	// test1()
	
	//注意这块，先行注册defer才会触发
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	// test2()
	//结果
	// test2
	// test1
	// 2019/06/28 14:12:01 i am dead

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	// panic("1")
	// panic("2")

	test3(5,0)
	//结果：说明除数为0的case被catch住了，系统给返回了一个兜底
	// @@x/y= 5 0
	// x/y= 0
}

func test1(){//dead被捕获就不往下走
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	panic("dead")
	println("到不了这")
}

func test2(){
	defer println("test1")
	defer println("test2")
	panic("i am dead")
}

func test3(x,y int){
	z := 0
	func(){
		defer func(){
			if recover() != nil {
				z = 0
			}
		}()
		println("@@x/y=", x,y)
		z = x/y
		println("@@@x/y=", z)
	}()
	println("x/y=", z)
}
