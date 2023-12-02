package main

import (
	"fmt"
	"sync"
	"time"
)

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s正在购物\n", name)
	time.Sleep(time.Second * 1)
	fmt.Printf("%s购物完毕\n", name)

	moneyChan <- money

	wait.Done()
}

var moneyChan = make(chan int) // 声明并初始化一个长度为0的信道

func main() {

	var wait sync.WaitGroup
	wait.Add(3)
	// 并发执行 主线程结束 协程也结束
	go pay("张三", 2, &wait)
	go pay("李四", 3, &wait)
	go pay("王五", 5, &wait)

	go func() {
		defer close(moneyChan)
		wait.Wait()
		//close(moneyChan)
	}()

	var moneyList []int
	for money := range moneyChan {
		fmt.Println("收到", money)
		moneyList = append(moneyList, money)
	}
	//for {
	//	money, ok := <-moneyChan
	//	fmt.Println("收到", money, ok)
	//	if !ok {
	//		break
	//	}
	//}
	//wait.Wait()

	fmt.Println("购物完毕")
	fmt.Println(moneyList)
}
