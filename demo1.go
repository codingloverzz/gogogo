package main

import (
	"fmt"
)

func getUserInfo() (string, int) {
	return "zhuwei", 18
}

func main() {

	var (
		age  int    = 18
		name string = "zhuwei"
	)
	const (
		a = 100
		b
		c
	)

	fmt.Printf("name is %v age is %v\n", name, age)

	var username, userAge = getUserInfo()

	fmt.Println(username, userAge)
}
