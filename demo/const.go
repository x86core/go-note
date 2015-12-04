package main

import (
	"fmt"
)

//常量表达式，必须是内置函数
//运行时不会被修改

const (
	b  = "go on"
	bl = len(b)
)

// 遇到下个const关键字时，重置iota
const (
	ONE = iota //0
	TWO
	THREE

	FOUR
	FIVE = "FIVE"

	SIX = iota // 恢复序  6-1
)

func main() {
	// b = "let do" //常量是不允许重新赋值的
	fmt.Println(b)
	fmt.Println(bl)

	fmt.Println(ONE, THREE, FIVE, SIX)
}
