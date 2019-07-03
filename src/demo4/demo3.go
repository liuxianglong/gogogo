package main
import(
	"fmt"
)
func main() {
	s := "中国人"
	rs := []rune(s)
	rs[1] = '大'
	fmt.Println(string(rs))
}
