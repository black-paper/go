package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/workspaces/go/lesson/statement/16/16.go")
	if err != nil {
		log.Fatalln("Error!", err)
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err2 := file.Read(data)
	if err2 != nil {
		log.Fatalln("Error2!", err2)
	}
	fmt.Println(count, string(data))
}
