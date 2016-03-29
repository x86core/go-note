package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func main() {

	type T struct {
		XMLName xml.Name `json:"-" xml:"node"` // node name
		Name    string   `xml:"name"`
		Age     int      `xml:"age"`
	}

	var t T

	t.Name = "a  name "
	t.Age = 27

	x, _ := xml.Marshal(t)

	fmt.Println(string(x))
	j, _ := json.Marshal(t)

	fmt.Println(string(j))

}
