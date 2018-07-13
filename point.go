package main

import "fmt"

func main() {
	c, d := 3, 4
	//swap(&c, &d)
	c, d = swaps (c, d)
	fmt.Println(c, d)
	//fmt.Println(e, f)
}
// 指针
func swap (a, b *int) {
	*b, *a = *a, *b
}

func swaps (a, b int) (int, int) {
	return b, a
}
