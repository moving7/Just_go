package main

import (
	"fmt"
)

func main() {
	//定义
	var arr1 [5]int
	arr2 := [6]int{1, 3, 5, 7, 9, 8}
	arr3 := [...]int {2, 4, 6, 8, 10}

	var erarray [4][5]int




	fmt.Println(arr1, arr2, arr3)


	fmt.Println(erarray)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	// range 用来遍历数组和切片的时候返回索引和元素值
	// 如果我们不要关心索引可以使用一个下划线(_)来忽略这个返回值
	for key,value := range arr3 {
		fmt.Println(key,value)
	}
}
