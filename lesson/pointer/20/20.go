package main

import "fmt"

type Vertex struct {
	X, Y    int
	private string
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	v.private = "ssss"
	fmt.Println(v.X, v.Y, v.private)

	v2 := Vertex{X: 100}
	fmt.Println(v2)

	var v3 Vertex
	fmt.Println(v3)

	v4 := new(Vertex)
	fmt.Println(v4)

	v5 := &Vertex{}
	fmt.Println(v5)
}
