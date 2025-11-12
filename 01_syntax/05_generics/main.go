package main

type Number interface {
	int | uint
}

func Sum[T Number](list []T) T {

	var sum T
	for _, v := range list {
		sum += v
	}
	return sum
}

func main() {
	println(Sum([]int{1, 2, 3}))
}
