package first

import "fmt"

func main() {
	enums()
}

func enums () {
	const (
		php = iota
		python
		java
		android
	)
	fmt.Println(php, java, python)

	const (
		b = 1 << (10 *  iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
