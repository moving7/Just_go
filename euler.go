package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

func main() {
	euler()
}

func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))

	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi) + 1)

	fmt.Println(
		cmplx.Exp(1i*math.Pi) + 1)
		//cmplx.Pow(math.E, 1i * math.Pi) + 1
		//)

}
