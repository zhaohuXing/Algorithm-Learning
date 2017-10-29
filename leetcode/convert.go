package main

import "log"

func main() {
	result := convert("PAYPALISHIRING", 3)
	expected := "PAHNAPLSIIGYIR"

	log.Println(result == expected)
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	result := ""
	slen := len(s)
	for i := 0; i < numRows; i++ {
		for j := i; j < slen; j += 2 * (numRows - 1) {
			result += string(s[j])
			if i > 0 && i < numRows-1 {
				index := j + 2*(numRows-i-1)
				if index < slen {
					result += string(s[index])
				}
			}
		}
	}
	return result
}
