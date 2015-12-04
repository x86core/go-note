package main

import (
	"fmt"
	// "math/cmplx"
)

func main() {

	v := 6 + 3i //(3+3i)*(3+3i)= 3*3+2*(3*3i)+3i*3i=9+18i+9*i^2 = 0+18i

	fmt.Println(v * v)
}
