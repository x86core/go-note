package main

import (
	// "bufio"
	"fmt"
	"io"
	// "strconv"
	"bytes"
	// "os"
)

func main() {
	Read()
}

func Read() {
	buf := bytes.NewBufferString("next time")

	for {
		b, err := buf.ReadByte()

		if err == io.EOF {
			break
		}

		fmt.Println(string(rune(b)))

	}

	for b, err := buf.ReadByte(); err != io.EOF; {
		fmt.Println(string(rune(b)))
	}
}

func Write() {

	// str := []byte("lucky")

	nb := bytes.NewBufferString("day ")

	// nb.Write(str)

	nb.WriteRune('中')

	fmt.Println(nb.String())
}
