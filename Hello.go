package main

import (
	"fmt"
	"io/ioutil"
)

func add(a, b int) int {
	return a + b
}

func arraytest() {
	a := [...]int{1: 1, 30: 4}
	b := a[1:3]
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(b)
	fmt.Println(a)
}

func main() {
	fmt.Println("hello world")
	fmt.Println(add(1, 3))
	const fileName = "abc.txt"

	if contents, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(contents)
		fmt.Printf("%s", contents)
	}
	arraytest()
}
