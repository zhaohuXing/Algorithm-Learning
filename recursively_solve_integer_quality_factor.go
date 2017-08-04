package main

func isPrimeNumber(num int) bool {
	return false
}

func maxPrimeNumber(num int) int {
	return 0
}
func F(num int, collection []int) []int {
	if isPrimeNumber(num) {
		return append(collection, num)
	} else {
		max := maxPrimeNumber(num)
		return F(num/max, append(collection, max))
	}
}
func main() {

}
