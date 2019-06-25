package main
import "fmt"

func main() {
	a := [3]int{10,20,30}
	{
		for i, x := range a {
			fmt.Println(i,x)
		}
		//结果0 10 
		// 1 20 
		// 2 30
	}

	{
		for i := range a {
			fmt.Println(i, a[i])
		}
		//结果0 10 
		// 1 20 
		// 2 30
	}

	{
		for i, x := range a {//数组
			if i == 0 {
				a[0] += 100
				a[1] += 100
				a[2] += 100
			}
			fmt.Printf("x: %d, data: %d\n", x, a[i])
			//结果
			// x: 10, data: 110
			// x: 20, data: 120
			// x: 30, data: 130
			//直接从数组a中复制品中取值，故意x的值只会走初始的a数组里面的值
		}
	}
	{
		for i, x := range a[:] {//切片
			if i == 0 {
				a[0] += 100
				a[1] += 100
				a[2] += 100
			}
			fmt.Printf("x: %d, data: %d\n", x, a[i])
		}
		//结果
		// x: 110, data: 210
		// x: 220, data: 220
		// x: 230, data: 230
		//因为切片是直接从底层数组里取值，而此时底层数组有变化，故x也会变化，又因为i==0时，x已被赋值，所以只会有i非0时数据的变化
		
	}

}
