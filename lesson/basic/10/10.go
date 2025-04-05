package main

import (
	"fmt"
)

func foo(params ...int) {
	fmt.Println(params)
}

func main() {
	foo(1, 2, 3, 4, 5)
	foo(1, 2, 3)
}
