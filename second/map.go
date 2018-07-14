package main

import (
	"fmt"
)

func main() {
	m := map[string]string {
		"name": "小明",
		"sex": "1",
		"age": "20",

	}

	fmt.Println(m)

	m2 := make(map[int]string) // m2 == empty map

	fmt.Println(m2)

	var m3 map[int]string // m3 == nil

	fmt.Println(m, m2 ,m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	courseName := m["name"]
	fmt.Println(courseName)

	//find name
	find_text, err := m["name"]
	if (err) {
		fmt.Println("找到了 ", find_text)
	} else {
		fmt.Println("没有找到")
	}

	find_text, err = m["name"]
	delete(m, "name")
	fmt.Println(m)
}
