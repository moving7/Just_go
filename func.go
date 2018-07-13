package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func pow (a, b int) int {
	return int (math.Pow(float64(a), float64(b)))
}

func sum (num ...int) int {
	s := 0
	for i := range num {
		s += num[i]
	}
	return s
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	//fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(func (a int, b int) int {
		return int (math.Pow(
			float64(a), float64(b)))
	}, 3, 4))

	//
	//if result, err := eval(3, 4, "/");
	//err != nil {
	//	fmt.Println("Error: ", err)
	//} else {
	//	fmt.Println(result)
	//}


	//fmt.Println(eval(9, 3, "/"))
	//fmt.Println(3*7, "12\n2")
	//fmt.Println(div(10, 3))
	//a,b := div(20, 6)
	//fmt.Println(a, b)
	//
	//q, r := divv (31, 5)
	//fmt.Println(q, r)
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		q, _ := divv (a, b)
		return q, nil
	default:
		//中断执行
		//panic("unsupported operation: " + op)
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

// 求余
func div (a, b int) (int, int) {
	return a / b, a % b
}

func divv (a, b int) (q, r int) {
	//q = a / b
	//r = a % b
	return a / b, a % b
}

func apply (op func(int, int) int, a,b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println("Calling function %s with args" +
		"(%d, %d) \n", opName, a, b)
	return op(a, b)
}