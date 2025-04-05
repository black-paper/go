package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	l := []string{"Python", "Go", "TypeScript"}
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
