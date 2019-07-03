package main

func main() {
	demo1()
	//结果：
	// 101 102
	// defer a,y= 1 102

	demo2()
	//结果：
	// b
	// a

	//return 和 panic都会终止当前函数，引发延迟调用
	i := demo3()
	println(i)
	//结果
	// xxxx
	// 110
	//可以看出在函数return后触发了defer的闭包，改变了返回值
}

func demo1(){
	x,y:=1,2
	defer func(a int){
		println("defer a,y=",a,y)
	}(x)
	x+=100
	y+=100
	println(x,y)
}

func demo2(){//先进后出
	defer println("a")
	defer println("b")
}

func demo3() (i int){//return后触发defer
	defer func(){
		println("xxxx")
		i+=100
	}()
	return 10
}
