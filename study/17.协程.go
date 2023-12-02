package main

import (
	"fmt"
	"sync"
	"time"
)

func shopping(name string, wait *sync.WaitGroup) {
	fmt.Printf("%s正在购物\n", name)
	time.Sleep(time.Second * 1)
	fmt.Printf("%s购物完毕\n", name)
	//wait--
	wait.Done()
}

// var wait int
//var wait sync.WaitGroup

func main() {
	var wait sync.WaitGroup
	//wait = 3
	startTime := time.Now()
	// 顺序执行
	//shopping("张三")
	//shopping("李四")
	//shopping("王五")
	//fmt.Println("购物完毕", time.Now().Sub(startTime))

	wait.Add(3)
	// 并发执行 主线程结束 协程也结束
	go shopping("张三", &wait)
	go shopping("李四", &wait)
	go shopping("王五", &wait)

	//time.Sleep(time.Second * 2)

	//for wait > 0 {
	//	time.Sleep(time.Second * 1)
	//}

	//for {
	//	if wait == 0 {
	//		break
	//	}
	//}

	wait.Wait()

	fmt.Println("购物完毕", time.Since(startTime))
}
