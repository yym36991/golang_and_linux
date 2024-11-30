package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 6 {
			fmt.Println("Skipping the rest of the loop when i = 3")
			goto skipLoop // 跳转到标签 skipLoop
		}
		fmt.Println(i) // 打印当前的 i
	}

skipLoop:
	fmt.Println("Exited the loop")
}
// 写了一个读文件的脚本。