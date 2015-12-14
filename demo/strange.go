package main

import "fmt"

type MagicError struct{}

func (MagicError) Error() string {
	return "[Magic]"
}

func Generate() *MagicError {
	return nil
}

func Test() error {
	return Generate()
}

func main() {
	if Test() != nil {
		fmt.Println("Hello, Mr. Pike!")
	}
}
