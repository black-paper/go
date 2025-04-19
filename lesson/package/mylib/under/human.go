package under

import "fmt"

// struct名の先頭が小文字だと、private
type person struct {
	Name string
	age  int
}

type Person struct {
	Name string
	Age  int
}

func Hello() {
	fmt.Println("Hello")
}
