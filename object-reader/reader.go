package main

import (
	"io"
	"log"
	"os"
)

func main() {
	objectFileName := os.Args[1]
	file, err := os.Open(objectFileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	file.Read()
}
