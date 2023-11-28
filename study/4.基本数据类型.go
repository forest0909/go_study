package main

import "fmt"

func main() {
	//var aget = 127

	//var u8 uint8 = 255
	// uint8: 0 ~ 255
	// 00000000 = 0
	// 11111111 = 255 = 2^8 - 1
	// int8
	// 0 0000000 = 0
	// 0 1111111 = 127\
	// 1 0000001 = -1
	// 1 0000000 = -128

	var a byte = 'a' //ascii 里面的字符
	fmt.Print("a = ", a)

	var b rune = '中'
	fmt.Print("b = ", b)
}
