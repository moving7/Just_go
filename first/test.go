package first

import "fmt"

var (
	aa = 1
	bb = 2
	cc = 3
)
// mian

func main() {
	fmt.Println("Hello Word")
	variable()
	varInitialValue()
	varType()
	varShort()
	println(aa, bb ,cc)
}

func varInitialValue() {
	var abc int = 1
	var str string = "123"
	fmt.Println(abc, str)
}

func varType() {
	var abc = 1
	var str = "123"
	fmt.Println(abc, str)
}

func varShort() {
	abc , str := 1123, 123
	fmt.Println(abc, str)
}

func variable() {
	var abc int
	var str string
	// 默认
	/**默认给值 */
	//println 只能打印值
	//printf 可以跟格式 （%d %a）类似C
	fmt.Printf("%d %q\n", abc, str)
	//fmt.Println(abc, str)
}