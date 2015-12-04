package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	table := map[string]string{"I": "yes", "Am": "you", "A": "are", "Coder": "right"}

	buf := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(buf.Write())
	dec := json.NewDecoder(os.Stdin)
	out := enc.Encode(table)
	fmt.Println("=>", out)

	fmt.Println("<=", dec.Decode(&out))

}
