package main

import (
	"log"
)

func main() {
	log.Println("logging")

	log.Fatalln("Fatal!")
	log.Println("after fatal")
}
