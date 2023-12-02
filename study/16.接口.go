package main

import "fmt"

type SingInterface interface {
	Sing()
	GetName() string
}

type Chicken struct {
	Name string
}

func (c Chicken) Sing() {
	println(c.Name, "正在唱歌")
}

func (c Chicken) GetName() string {
	return c.Name
}

type Cat struct {
	Name string
}

func (c Cat) Sing() {
	println(c.Name, "正在唱歌")
}

func (c Cat) GetName() string {
	return c.Name
}

func Sing(c SingInterface) {
	ch, ok := c.(Chicken)
	fmt.Println(ch, ok)

	switch anmial := c.(type) {
	case Chicken:
		fmt.Println(anmial.Name, "是一只小鸡")
	case Cat:
		fmt.Println(anmial.Name, "是一只小猫")
	default:
		fmt.Println("不知道是什么动物")
	}
	c.Sing()
	fmt.Println("正在唱歌的是：", c.GetName())
}

// 空接口 任何接口都实现了空接口
type EmptyInterface interface {
}

func Print(val EmptyInterface) {
	fmt.Println(val)
}

func Print1(val interface{}) {
	fmt.Println(val)
}

func Print2(val any) {
	fmt.Println(val)
}

func main() {

	ch := Chicken{
		Name: "小鸡",
	}
	ca := Cat{
		Name: "小猫",
	}
	Sing(ch)
	Sing(ca)
	Print(ch)
	Print(ca)
	Print1(ch)
	Print1(ca)
	Print2(ch)
	Print2(ca)
}
