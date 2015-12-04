package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	Tstring := "need to encrypt"

	Md5inst := md5.New()

	Md5inst.Write([]byte(Tstring))
	Result := Md5inst.Sum([]byte(""))

	fmt.Printf("%x\n\n", Result)

	Sha1inst := sha1.New()
	Sha1inst.Write([]byte(Tstring))
	Result = Sha1inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)
}
