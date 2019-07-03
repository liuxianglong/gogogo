package main
import(
	"fmt"
	"strings"
)
func main() {
	s := "中国人"
	s += "最好的"
	fmt.Println(s)
	//结果：
	// 中国人最好的
	//用这个方式来拼接，每次都要重新分配内存。如果创建超大字符串，此方式性能极差

	{
		s := make([]string, 5)
		for i:=0;i<5;i++ {
			s[i] = "a"
		}
		fmt.Println(strings.Join(s, ""))
		//结果：
		// aaaaa
		//此方式相当于提前分配了内存空间，一次性完成内存分配 
	}
}
