# 6. ZigZag Conversion

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```
P   A   H   N
A P L S I I G
Y   I   R
```

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

```
string convert(string text, int nRows);
```

convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

解题思路：本质是一个找规律的数学题，首行和最后一行每个元素之间的坐标相差 `2 * (numRows - 1)`, 其它的在其基础上计算 `j + 2*(numRows - i - 1)`。

#### 代码实现 

``` go
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
```
