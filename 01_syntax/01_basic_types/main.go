package main

import (
	"fmt"
)

func main() {
	// var str = "hello你好"
	// println(utf8.RuneCountInString(str))
	// Byte()

	DeferClosure()

}

func Byte() {
	var str = "hello"
	// var char byte = 'a'
	var bs = []byte(str)

	fmt.Println(string(bs))
	println('a' + 1)
}

func DeferClosure() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}
