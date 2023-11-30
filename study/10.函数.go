package main

import "fmt"

func main() {
	sum := getSum(1, 2)
	fmt.Println("sum = ", sum)

	sum1, sub1 := getRes(1, 2)
	fmt.Printf("sum1 = %d, sub1 = %d\n", sum1, sub1)

	sum2, sub2 := getRes1(1, 2)
	fmt.Printf("sum2 = %d, sub2 = %d\n", sum2, sub2)
}

func getSum(a int, b int) int {
	res := a + b
	return res
}

func getRes(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

func getRes1(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}
