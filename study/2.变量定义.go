package main

import "fmt"

func main() {

	// 定义变量
	var a string
	a = "Hello World!"
	fmt.Println(a)

	// 定义变量并初始化
	var b string = "Hello World!"
	fmt.Println(b)

	// 定义变量并初始化，省略类型
	var c = "Hello World!"
	fmt.Println(c)

	// 定义变量并初始化，省略var
	d := "Hello World!"
	fmt.Println(d)

}
