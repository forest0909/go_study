package main

import "fmt"

func main() {
	// for循环
	i := 1
	for {
		fmt.Print(" i = ", i)
		i++
		if i == 10 {
			break
		}
	}
	fmt.Println("")
	// for循环
	j := 1
	for j <= 10 {
		fmt.Print(" j = ", j)
		j++
	}
	fmt.Println("")
	// for循环
	for k := 1; k <= 10; k++ {
		fmt.Print(" k = ", k)
	}
	fmt.Println("")
	// for循环 打印九九乘法表
	for m := 1; m <= 9; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%d * %d = %d ", n, m, m*n)
		}
		fmt.Println("")
	}

}
