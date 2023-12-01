package main

import "fmt"

func main() {

	c1 := Class{
		Name:      "一班",
		ClassName: "一年级一班",
	}

	s1 := Student{
		Class: c1,
		Name:  "张三",
	}
	s1.Study()
	s1.Info()
	s1.SetName("李四")
	s1.Info()
}

type Student struct {
	Class
	Name string
}

func (s Student) Study() {
	fmt.Println(s.Name, "正在学习")
}

func (s Student) Info() {
	fmt.Printf("名字：%s 班级: %s\n", s.Name, s.Class.Name)
	fmt.Printf("名字：%s 班级: %s\n", s.Name, s.ClassName)
}

func (s *Student) SetName(name string) {
	s.Name = name
}

type Class struct {
	Name      string
	ClassName string
}
