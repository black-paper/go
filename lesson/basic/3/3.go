package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "Hello shota"
	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(strings.Contains(s, "Shota"))
	fmt.Println(`Test
	                       Shota
Test`)
}
