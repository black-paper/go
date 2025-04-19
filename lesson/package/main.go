package main

import (
	"fmt"

	"lesson/package/mylib"
	"lesson/package/mylib/under"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))
	under.Hello()

	human := under.Person{Name: "shota", Age: 99}
	fmt.Println(human)
}
