package main

import (
	"fmt"
	// "path"
	"os"
	"path/filepath"
)

func main() {

	str := "../file/fd.txt"

	fmt.Println(filepath.Abs(str))

	filepath.Walk("../file", func(path string, f os.FileInfo, err error) error {
		fmt.Println(path)
		fmt.Println(f)
		fmt.Println(err)
		return nil
	})

}
