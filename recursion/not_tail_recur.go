package main

import "fmt"

func main() {

	fmt.Println(calcFactorial(5))
}

func calcFactorial(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else if n == 1 {
		return n
	}
	return n * calcFactorial(n-1)
}

// 递归自己调用自己
// 每个方法的创建都对应着 一个栈帧的创建
// 方法的从执行到结束， 对应着栈帧的进出栈
// 这种递归容易产生 栈溢出
