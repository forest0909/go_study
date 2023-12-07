package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//for {
	//	var byteData = make([]byte, 1024)
	//	n, err := conn.Read(byteData)
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(byteData[:n]))
	//}

	for {
		//var byteData = make([]byte, 1024)
		var text string
		fmt.Println("请输入内容：")
		fmt.Scanln(&text)
		if text == "q" {
			break
		} else {
			n, err := conn.Write([]byte(text))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("发送成功，共发送了", n, "个字节")
		}
	}
}
