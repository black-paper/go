package main

import (
	"fmt"
)

func add(x, y int) (int, int) {
	return x + y, x - y
}

func calc(price, count int) (result int) {
	return price * count
}

func main() {
	add, de := add(1, 2)
	fmt.Println(add, de)

	r := calc(100, 2)
	fmt.Println(r)

	f := func() {
		fmt.Println("Inner Func")
	}
	f()
}
