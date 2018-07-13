package first

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func main() {
	fmt.Println(
		converToBin(5), //101
		converToBin(10), //1101
		converToBin(10219012),
		converToBin(0),
	)

	readFile("start/abc.log")
}

func converToBin (n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//读取文件
func readFile (filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic (err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}