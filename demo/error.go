package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	ret, err := Sqrt(-4)

	if err != nil {
		fmt.Printf("out:%f\n", ret)
	} else {
		fmt.Printf("err:%v\n", err)
	}

	r, e := Sqrt(4)

	if e != nil {
		fmt.Printf("err:%v\n", e)
	} else {
		fmt.Printf("out:%f\n", r)
	}

}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("不能为负")
	}

	return math.Sqrt(f), nil
}
