package main

import (
	"bufio"
	"os"
)

func main() {
	//bs, err := os.ReadFile("study/21.文件操作/test.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(bs))

	//file, err := os.Open("study/21.文件操作/test.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//for {
	//	bs := make([]byte, 1024)
	//	n, err := file.Read(bs)
	//	if err == io.EOF {
	//		break
	//	}
	//	//fmt.Println(string(bs), n)
	//	fmt.Println(string(bs[:n]))
	//}

	file, err := os.Open("study/21.文件操作/test.txt")
	if err != nil {
		panic(err)
	}
	//buf := bufio.NewReader(file)
	//
	//for {
	//	line, _, err := buf.ReadLine()
	//	if err != nil {
	//		break
	//	}
	//	fmt.Println(string(line))
	//}

	buf := bufio.NewScanner(file)
	buf.Split(bufio.ScanWords)
	var index int
	for buf.Scan() {
		index++
		line := buf.Text()
		println(line, index)
	}
}
