package main

import (
	"fmt"
	"net"
	"reflect"
)

func main() {

	ip, err := net.LookupIP("www.baidu.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ip list=>", ip)
	fmt.Println("data type: ", reflect.TypeOf(ip))
	fmt.Println(ip[0])
}
