package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	if result := by2(10); result == "ok" {
		fmt.Println("Wow")
	}
}
