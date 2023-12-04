package main

import "os"

func main() {
	file, err := os.OpenFile("study/21.文件操作/test2.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte("hello world\n"))
}
