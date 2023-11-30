package main

import "fmt"

func main() {

	var a int = 10
	var ip *int
	ip = &a
	fmt.Println("a的地址是：", &a)
	fmt.Println("ip的值是：", ip)
	fmt.Println("ip指向的值是：", *ip)

	var y = 10
	add(y)
	fmt.Println("main函数调用结束！y = ", y)
	add1(&y)
	fmt.Println("main函数第二次调用结束！y = ", y)
}

func add(a int) {
	a++
	fmt.Println("add函数调用结束！a = ", a)
}

func add1(a *int) {
	*a++
	fmt.Println("add1函数调用结束！a = ", *a)
}
