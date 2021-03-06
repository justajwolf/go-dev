package main

import (
	"fmt"
	"strings"
)

/*
题目344. 反转字符串(https://leetcode-cn.com/problems/reverse-string/)
	描述:
		编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
		不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
		你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

	示例1：
		输入：["h","e","l","l","o"]
		输出：["o","l","l","e","h"]
	示例2：
		输入：["H","a","n","n","a","h"]
		输出：["h","a","n","n","a","H"]
*/
func main() {
	reverseString([]byte(strings.Join([]string{"h", "e", "l", "l", "o"}, "")))
	reverseString2([]byte(strings.Join([]string{"H", "a", "n", "n", "a", "h"}, "")))
}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	fmt.Println(string(s))
}

func reverseString2(s []byte) {
	for i, limit := 0, len(s)/2; i < limit; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(string(s))
}
