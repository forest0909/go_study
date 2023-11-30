package main

import "fmt"

func main() {

	var (
		username   string
		password   string
		repassword string
	)
	for {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&username)
		fmt.Println("请输入密码：")
		fmt.Scanln(&password)
		fmt.Println("请再次输入密码：")
		fmt.Scanln(&repassword)

		// 判断用户名是否为空 判断密码是否为空
		if username == "" || password == "" || repassword == "" {
			fmt.Println("用户名或密码不能为空！")
			continue
		}

		if password == repassword {
			fmt.Println("注册成功！")
			break
		} else {
			fmt.Println("两次密码不一致，请重新输入！")
		}
	}
}
