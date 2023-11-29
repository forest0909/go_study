package main

import "fmt"

func main() {
	var nameList [3]string = [3]string{"张三", "李四", "王五"}
	fmt.Println("nameList = ", nameList)
	// 索引
	fmt.Println("nameList[0] = ", nameList[0])
	fmt.Println("nameList[1] = ", nameList[1])
	fmt.Println("nameList[2] = ", nameList[2])
	fmt.Println("nameList[2] = ", nameList[len(nameList)-1])

	nameList[0] = "赵六"
	fmt.Println("nameList = ", nameList)
}
