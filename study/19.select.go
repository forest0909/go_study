package main

import (
	"fmt"
	"sync"
	"time"
)

func send(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s正在购物\n", name)
	time.Sleep(time.Second * 1)
	fmt.Printf("%s购物完毕\n", name)

	moneyChan1 <- money
	nameChan1 <- name

	wait.Done()
}

var moneyChan1 = make(chan int)   // 声明并初始化一个长度为0的信道
var nameChan1 = make(chan string) // 声明并初始化一个长度为0的信道
var downChan = make(chan struct{})

func main() {

	var wait sync.WaitGroup
	wait.Add(3)
	// 并发执行 主线程结束 协程也结束
	go send("张三", 2, &wait)
	go send("李四", 3, &wait)
	go send("王五", 5, &wait)

	go func() {
		defer close(moneyChan1)
		defer close(nameChan1)
		defer close(downChan)
		wait.Wait()
		//close(moneyChan)
	}()

	var moneyList []int
	var nameList []string

	var event = func() {
		for {
			select {
			case money := <-moneyChan1:
				moneyList = append(moneyList, money)
			case name := <-nameChan1:
				nameList = append(nameList, name)
			case <-downChan:
				return
			case <-time.After(time.Second * 10):
				fmt.Println("超时")
				return
			}
		}
	}
	event()
	//go func() {
	//	for money := range moneyChan1 {
	//		fmt.Println("收到", money)
	//		moneyList = append(moneyList, money)
	//	}
	//}()
	//for name := range nameChan1 {
	//	nameList = append(nameList, name)
	//}

	fmt.Println("购物完毕")
	fmt.Println(moneyList)
	fmt.Println(nameList)
}
