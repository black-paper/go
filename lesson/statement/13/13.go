package main

import (
	"fmt"
	"time"
)

func main() {
	os := "Windows"

	switch os {
	case "Windows":
		fmt.Println("Windows")
	case "mac":
		fmt.Println("mac")
	default:
		fmt.Println("default")
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 10:
		fmt.Println("10時未満")
	case t.Hour() <= 10:
		fmt.Println("10時以降")
	}
}
