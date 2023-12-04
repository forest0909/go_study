package main

import (
	"fmt"
	"os"
)

func main() {
	bs, err := os.ReadFile("/Users/zzl/CodeSpace/goProject/go_study/study/21.文件操作/test.txt")
	if err != nil {
		fmt.Println("文件读取失败！")
		return
	}
	fmt.Println(string(bs))
}
