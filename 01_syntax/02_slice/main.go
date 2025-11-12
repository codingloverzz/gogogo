package main

import "fmt"

func main() {
	SubSlice()

}

func SubSlice() {

	//只要切片或子切片没有发生扩容，他们就共享内容
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := s1[3:5]
	s2[0] = 999

	fmt.Println(s1, s2)
}
