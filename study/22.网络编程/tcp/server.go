package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	//创建 tcp 的监听地址
	tcpAddress, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	listen, err := net.ListenTCP("tcp", tcpAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("服务器启动成功")
	//for {
	//	conn, err := listen.AcceptTCP()
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println("客户端连接成功:" + conn.RemoteAddr().String())
	//	conn.Write([]byte("hello world"))
	//	time.Sleep(time.Second * 2)
	//	conn.Close()
	//}
	for {
		var byteData = make([]byte, 1024)
		conn, err := listen.Accept()
		fmt.Println(conn.RemoteAddr().String() + "连接成功")
		if err != nil {
			fmt.Println(err)
			break
		}
		for {
			n, err := conn.Read(byteData)
			fmt.Println(string(byteData[:n]))
			if err != io.EOF {
				fmt.Println(conn.RemoteAddr().String())
				//break
			}
		}
	}
}
