package main

import "fmt"

func main() {
	//F1()

	//F2()
	Compose()
}

func F1() {
	u1 := &User{
		name: "zhuwei",
		age:  18,
	}

	println(u1)
	fmt.Println(u1)
}

func F2() {
	u := User{
		name: "zhuwei",
		age:  18,
	}
	u.ChangeName("zhangxiao")
	u.ChangeAge(20)

	fmt.Println(u)
}
