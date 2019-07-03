package main
import(
	"fmt"
)
func main() {
	s := "中国人"
	for i:=0; i < len(s); i++{//byte
		fmt.Printf("%d:[%c]\n", i, s[i])
	}
	//因为一个中文字符对应三个字节
	// 0:[ä]
	// 1:[¸]
	// 2:[­]
	// 3:[å]
	// 4:[]
	// 5:[½]
	// 6:[ä]
	// 7:[º]
	// 8:[º]

	for k,v := range s {
		fmt.Printf("%d:[%c]\n", k, v)
	}
	//结果
	// 0:[中]
	// 3:[国]
	// 6:[人]
}
