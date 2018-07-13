package main

import (
	"math"
	"fmt"
)

func main() {
	three()
}

func three () {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
