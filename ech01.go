package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	//依赖于字符串默认初始化零值机制  var s string
	for i := 1; i < len(os.Args); i++ {
		// 短变量声明，只能函数内部使用 i := 1
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
