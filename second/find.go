package main

import "fmt"

func main() {
	//lastOccurred[x]
	/**
	 * lastOccureed 查找指定内容最后出现的位置
	 * 1. 从来没出现过 或者小于start 无需操作
	 * 2. >= start 更新start
	 */

	 fmt.Println(
	 	lengthOfNonRepeatingSubStr("abcabcdeaabcd"))
	 fmt.Println(
	 	lengthOfNonRepeatingSubStr("adskfladsflkwfiowe"))
	 fmt.Println(
	 	lengthOfNonRepeatingSubStr("测试测试请问考虑数量的咖啡机"))
	 fmt.Println(
	 	lengthOfNonRepeatingSubStr("重复文字重复"))

}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccerred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccerred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccerred[ch] = i
	}

	return maxLength
}
