package first

import (
	"io/ioutil"
	"fmt"
)

func main() {
	const filename = "start/abc.log"
	contents, err := ioutil.ReadFile(filename)
	//if contents, err := ioutil.ReadFile(filename); err != nil {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", contents)
	}
}
