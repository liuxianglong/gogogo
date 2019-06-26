package main

func main() {
	for i:=0; i < 3; i++ {
		println(i)
		if i> 0 {
			goto exit
		}
	}
	exit:
	println("exit")
	//结果
	// 0
	// 1
	// exit	

	{
		outer:
		for x := 0; x < 5; x++{
			for y:= 0; y < 10; y++{
				if y > 2{
					println()
					continue outer
				}

				if x > 2 {
					break outer
				}

				print(x, ":", y, " ")
			}
		}
		//结果
		// 0:0 0:1 0:2
		// 1:0 1:1 1:2
		// 2:0 2:1 2:2
	}
}




