package main

import (
	"fmt"
)

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	fmt.Println(b)

	var c []int = []int{100, 200, 300}
	c = append(c, 400)
	fmt.Println(c)
	fmt.Println(c[1:3])
	fmt.Println(c[:3])
	fmt.Println(c[3:])

	var board = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(board)
}
