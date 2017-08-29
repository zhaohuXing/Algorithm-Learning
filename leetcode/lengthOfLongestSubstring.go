package main

import "fmt"

func main() {

	n := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(n)
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	tbl := make(map[byte]int, n)
	var left, curr, longest int

	for right := 0; right < n; right++ {
		if k, ok := tbl[s[right]]; ok && k >= left {
			left = k + 1
		}

		tbl[s[right]] = right
		curr = right - left + 1

		if curr > longest {
			longest = curr
		}
	}
	return longest
}
