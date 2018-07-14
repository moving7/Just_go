package main

import "fmt"

func main() {
	var arr = [...]int{-1,0,1,2,3,4,5,7,11,8,9}
	//切片
	fmt.Println("arr[2:8] = ", arr[2:8]) //1-8 < 8
	fmt.Println("arr[:8] = ", arr[:8]) // < 8
	fmt.Println("arr[2:] = ", arr[2:]) //> 2
	fmt.Println("arr[:] = ", arr[:]) //all data
	fmt.Println("\n\narr[:] = ", arr[2:3]) //test
	updateSlice(arr[:])
	fmt.Println(arr[:])
}

func updateSlice(s []int) {
	s[0] = 100
}