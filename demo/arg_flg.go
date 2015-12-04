package main

import (
	"flag"
	"fmt"
	// "os"
)

func main() {
	// os.Args // args

	host := flag.String("host", "localhost", " host to connect")
	port := flag.Int("port", 8080, " connect port")

	flag.Parse()

	fmt.Println("host:", *host)

	fmt.Println("port", *port)
}
