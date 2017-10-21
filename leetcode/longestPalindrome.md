# 5. Longest Palindromic Substring 

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

___Example:___

``` 
Input: "babad"

Output: "bab"

Note: "aba" is also a valid answer.
```

___Example:___

```
Input: "cbbd"

Output: "bb"
```

解题思路：
- 暴力遍历
- 分治
- XXX算法

### 暴力遍历
当遍历到下标为 i 的字符时，就以该字符为中心，像两端扩展, 暴力的遍历所有的字符。 

以 "cbbd" 为例:
- 遍历到 "c" 时, 其左右两端不不相等，则继续遍历。
- 遍历到 "b" 时，发下下一个也是 "b"，这时我们应该判断下，如果发下其右边与其相等，则继续遍历(这是个剪枝的操作)。
- 遍历到了第二个 "b" 时，发现左右不相等，则继续遍历。
- 遍历到 "d" 时，发现其左右不相等，由于遍历完了，遍历程序终止。
- 每次遍历都会比教下，看看发现的回文长度是否大于现在的回文长度，所以这边离下来，结果是 "bb"


实现代码如下：

``` go
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
```
