package main

import "fmt"

func main() {
	// 一但定义 zero value for slice is nil
	// go 没有 null; sli == nil
	var sli []int
	for i := 0; i < 10; i++ {
		printSli(sli)
		sli = append(sli, 2 * i + 1)
	}
	fmt.Println(sli)


	sli2 := make([]int, 16)

	sli3 := make([]int, 20, 32)

	printSli(sli2)
	printSli(sli3)

	copy(sli2, sli)
	printSli(sli2)

	//delete

	sli2 = append(sli2[:3], sli2[4:]...)
	printSli(sli2)

	front := sli2[0]
	sli2 = sli2[1:]
	tail := sli2[len(sli2) - 1]
	sli2 = sli2[:len(sli2) - 1]
	fmt.Println(front, tail)
	printSli(sli2)
}

func printSli(sli []int){
	fmt.Printf("%v, len = %d, cap = %d \n",
		sli, len(sli), cap(sli))
}
