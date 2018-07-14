package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "GoGoGo雷欧雷欧雷"
	fmt.Println(len(str))
	for _, b := range []byte(str) {
		fmt.Printf("%X ", b)
	}
	// ch is a rune
	for i, ch := range str {
		fmt.Printf("\n(%d %X) ", i, ch)
	}

	//utf8.RuneCountInString 获取字符数量
	fmt.Println("\ncount:",
	utf8.RuneCountInString(str))

	// []byte 获取字节
	bytes := []byte(str)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ",ch)
	}

	for i, ch := range []rune(str) {
		fmt.Printf("\n(%d %c) ", i, ch)
	}

}
