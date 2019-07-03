package main

import("fmt")

func a(){}
func b(){}
func test(x,y int,s string, _ bool) *int{ 
   return nil
} 

func test2(x *int) {
	fmt.Printf("pointer:%p, target: %v\n", &x, x)
}

func test3(s string, a ...int) {
	fmt.Printf("%T, %v\n", a, a)
}

func test4(a ...int) {
	for i:=range a{
		a[i] += 100
	}
}

func main() {
	println(a==nil)	//结果：false 函数只能判断是否是nil，不支持其他比较操作
	//println(a==b)	//不支持
	test(1, 2, "abc", true)

	a := 100
	p := &a
	fmt.Printf("pointer:%p, target: %v\n", &p, p)
	test2(p)

	test3("abc", 1,2,3,4)
	//结果：[]int, [1 2 3 4]

	arr := [3]int{10,20,30}
	test4(arr[:]...)
	fmt.Printf("%T, %v\n", arr, arr)
	//结果 [3]int, [110 120 130]
	slice := []int{10,20,30}
	test4(slice...)
	fmt.Printf("%T, %v\n", slice, slice)
	//结果 []int, [110 120 130]
}
