package main

import "fmt"

func main() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scanln(&age)

	if age <= 18 {
		fmt.Println("你还未成年！")
	}

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

	var week int
	fmt.Println("请输入星期几：")
	fmt.Scanln(&week)
	switch week {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("输入错误！")
	}
}
