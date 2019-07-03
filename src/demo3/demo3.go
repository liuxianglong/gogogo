package main

// import("fmt")


func main() {
	t := func (x string) func() {
		return func(){
			println(x)
		}
	}
	
	x := "hello"
	t2 := func(){
		println(x)
	}

	f := t("abcde")
	f()
	//结果： abcde

	t2()
	//结果： hello

	a := [3]int{1,2,3}
	for i,v := range a{
		z := func(){
			println(i,v)
		}
		z()
	}
	//结果
	// 	0 1
	//  1 2
	//  2 3

	for _,f := range test() {
		f()
	}
	//结果	
	// 	0xc420014060 2
	//  0xc420014060 2
	//原因：因为append仅仅是把匿名函数放入列表，并未执行，main执行的时候读取的是环境变量i最后的值，即2

	for _,f := range test2() {
		f()
	}
	//结果	
	// 0xc420014068 1
	// 0xc420014068 1
	//原因：因为append仅仅是把匿名函数放入列表，并未执行，main执行的时候读取的是环境变量z最后的值，即1

	for _,f := range test3() {
		f()
	}
	//结果
	// 	0xc420014070 0
	//  0xc420014078 1
	//原因：因为我们用了z:=i，相当于每次都是一个全新的环境变量

	t1,t2 := test4(10)
	t1()
	t2()
	//结果
	// 110
	// 110
	// 原因：环境变量变化了，依赖环境变量的值也相应变化

	zz := 1
	x1 := func(){
		println(zz)
	}
	x1()
	//结果： 1
	zz = 2
	x1()
	//结果： 2
}
func test() []func(){
	var s []func()
	for i:=0; i < 2; i++ {
		s = append(s, func(){
			println(&i, i)
		})
	}
	return s
}

func test2() []func(){
	var s []func()
	var z int
	for i:=0; i < 2; i++ {
		z = i
		s = append(s, func(){
			println(&z, z)
		})
	}
	return s
}

func test3() []func(){
	var s []func()

	for i:=0; i < 2; i++ {
		z := i
		s = append(s, func(){
			println(&z, z)
		})
	}
	return s
}

func test4(x int) (func(),func()) {
	return func() {
		x+=100
		println(x)
	},func(){
		println(x)
	}
}
