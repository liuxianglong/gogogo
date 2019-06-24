package main
import "fmt"

func main() {
	a,b,c := 1,2,3
	x := 1
	switch x {
	case a,b:
		x += 10
		fmt.Println(x)
		fallthrough
	case c :
		x += 10
		fmt.Println("fit")
	default :
		fmt.Println("no fit")
	}
//结果是 11 fit, fallthrough将继续往下执行，不会匹配条件表达式
	{
		switch x:=5; {
		case x > 5 :
			println("a")
		case x >0 && x <= 5:
			println("b")
		default :
			println("c")
		}
		//结果是b
	}
}
