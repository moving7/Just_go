package main

import "fmt"

func main() {
	var arr = [...]int{0,1,2,3,4,5,6,7,8,9}
	res1 := arr[1:6] //2,3,4,5
	res2 := res1[3:8] //5

	res3 := res2[1:5]

	fmt.Println("res1 = ", res1, "\nres2 = ", res2, "\nres3 = ", res3)

	//fmt.Println(arr[9])
	fmt.Printf("res1 = %v, len(res1) = %d, cap(res1) = %d\n",
		res1, len(res1), cap(res1))

}
