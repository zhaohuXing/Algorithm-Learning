package main

import (
	"log"
)

func main() {
	str := "bb"
	log.Println(longestPalindrome(str))
}

func longestPalindrome(s string) string {
	slen := len(s)
	if slen < 2 {
		return s
	}

	maxLen, maxLeft := 1, 0
	for i := 0; (i < slen) && (slen-i > (maxLen / 2)); {
		left, right := i, i
		for right < slen-1 && s[right] == s[right+1] {
			right++
		}
		i = right + 1

		for right < slen-1 && left > 0 && s[right+1] == s[left-1] {
			right++
			left--
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
			maxLeft = left
		}
	}
	return s[maxLeft : maxLeft+maxLen]
}
