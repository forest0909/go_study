package main

import "fmt"

func main() {
	var userMap map[int]string = map[int]string{
		1: "张三",
		2: "李四",
		4: "李四",
	}
	fmt.Println("userMap = ", userMap)
	fmt.Println("userMap[1] = ", userMap[1])
	fmt.Println("userMap[2] = ", userMap[2])
	fmt.Println("userMap[3] = ", userMap[3])
	fmt.Println("userMap[4] = ", userMap[4])
	value := userMap[4]
	value, ok := userMap[3]
	fmt.Println("value = ", value)
	fmt.Println("ok = ", ok)

	userMap[1] = "王五"
	fmt.Println("userMap = ", userMap)

	delete(userMap, 4)
	fmt.Println("userMap = ", userMap)

	//必须要初始化
	//var aMap map[string]string // 会报错
	var aMap = make(map[string]string)

	aMap["name"] = "张三"
	fmt.Println("aMap = ", aMap)
}
