package main

import "fmt"

func main() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你已经成年了！")
	} else {
		fmt.Println("你还未成年！")
	}
	switch {
	case age > 18:
		fmt.Println("你已经成年了！")
	case age < 18:
		fmt.Println("你还未成年！")
	default:
		fmt.Println("你刚好成年！")
	}
}
