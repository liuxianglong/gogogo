package main

// import("fmt")


func main() {
	func(s string){
		println(s)
	}("hello world")

	add := func(x,y int) int {
		return x+y
	}
	println(add(1,2))

	test(func(){
		println("test")
	})

	test2 := test2()
	println(test2(2,3))
}

func test(f func()) {
	f()
}

func test2() func(int,int)int{
	return func(x, y int) int{
		return x+y
	}
}