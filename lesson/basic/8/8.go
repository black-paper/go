package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"apple": 100, "lemon": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["new"] = 500
	fmt.Println(m)

	v, ok := m["apple"]
	fmt.Println(v, ok)

	m2 := make(map[string]int)
	m2["pc"] = 500
	fmt.Println(m2)
}
