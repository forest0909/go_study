package main

import "fmt"

// 批量定义全局变量
var (
	aa = 3
	bb = 4
	cc = 5
)

// 常量不可修改
const name = "go"

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

	fmt.Println("s1=", aa, "s2=", bb, "s3=", cc)

	//name = "golang"
	fmt.Println("name=", name)
}
