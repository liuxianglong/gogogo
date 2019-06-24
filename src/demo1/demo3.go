package main

type demo struct{
	x int
	y string
}
func main() {
	var d1 demo = demo{1,"a"}
	println(d1.x)
	d2 := demo{2, "b"}
	println(d2.y)
}
