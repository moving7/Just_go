package first

import (
	"math"
	"fmt"
)

func main() {
	consts()
}

func consts() {
	//const filename string =  "abb.log"
	//const a, b = 3, 4
	const (
		filename = "ab.log"
		a, b = 3, 4
		)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}