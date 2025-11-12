package main

import "fmt"

func main() {

	BaseMap()

	m := map[string]string{
		"name": "zhuwei",
		"age":  "20",
	}

	keys := getKeys(m)
	fmt.Println("看看keys", keys, len(keys))
}

func BaseMap() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	fmt.Println(m)
}

func getKeys(m map[string]string) (keys []string) {

	keys = make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}
	return keys

}
