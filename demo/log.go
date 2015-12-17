package main

import (
	"log"
	"os"
)

func main() {

	file, _ := os.OpenFile("./make.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	defer file.Close()
	var logger = log.New(file, "--", 0) // writer

	logger.Println("good afeter")
	logger.Println("good afeter")
}
