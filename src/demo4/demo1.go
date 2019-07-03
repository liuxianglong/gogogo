package main
import(
	"fmt"
)
func main() {
	s := "我是"
	fmt.Printf("%s\n", s)
	fmt.Printf("%x, len: %d\n", s, len(s))
	// 结果
	// 我是
	// e68891e698af, len: 6

	//使用"`"定义不做转义处理的原始字符串，支持跨行
	s2 := `line\r\n
	line 2`
	println(s2)

	s3 := "abcdefg"
	println(s3[:3])
	println(s3[1:4])
	println(s3[2:])
}
