package main

import "fmt"

func main() {
	fmt.Println(calcFactorial(5, 1))
}

func calcFactorial(n, a int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else if n == 1 {
		return a
	} else {
		return calcFactorial(n-1, n*a)
	}
}
