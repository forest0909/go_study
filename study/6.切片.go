package main

import (
	"fmt"
	"sort"
)

func main() {
	var nameList []string
	nameList = append(nameList, "张三")
	nameList = append(nameList, "李四")
	fmt.Println("nameList[0] = ", nameList[0])

	var nameList1 []string = []string{}
	fmt.Println("nameList1 = ", nameList1)

	var nameList2 []string = make([]string, 0)
	fmt.Println("nameList2 = ", nameList2)

	nameList3 := []string{}
	fmt.Println("nameList3 = ", nameList3)

	nameList4 := make([]string, 0)
	fmt.Println("nameList4 = ", nameList4)

	ageList := make([]int, 3)
	fmt.Println("ageList = ", ageList)

	var ageList1 []int = []int{1, 2, 3, 4, 5}
	slice := ageList1[1:3]
	fmt.Println("slice = ", slice)
	slice = ageList1[1:]
	fmt.Println("slice = ", slice)
	slice = ageList1[:3]
	fmt.Println("slice = ", slice)

	ageList2 := []int{4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7}
	sort.Ints(ageList2) // 升序
	fmt.Println("ageList2 = ", ageList2)
	sort.Sort(sort.Reverse(sort.IntSlice(ageList2))) // 降序
	fmt.Println("ageList2 = ", ageList2)
}
