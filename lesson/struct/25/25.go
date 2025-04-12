package main

import "fmt"

// interface{}...どんなインターフェースでも受け付ける
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "aaa")
	default:
		fmt.Println("default!")
	}
}

func main() {
	do(10)
	do("string")
	do(true)
}
