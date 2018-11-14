package main

import (
	"strings"
	"fmt"
)

//闭包概念
func Adder() func(int) int {
	var x int
	return func (d int) int{
		x += d
		return x
	}
}

//闭包测试二。
func makeSuffix(suffix string) func(string) string{
	return func(name string) string{
		if strings.HasSuffix(name, suffix) == false{
			return name + suffix
		}
		return name
	}
}

func main() {
	p := Adder()
	fmt.Println(p(1))
	fmt.Println(p(100))
	fmt.Println(p(1000))

	//闭包测试二
	f1 := makeSuffix(".jpg")
	f2 := makeSuffix(".bmp")

	fmt.Println(f1("123"))
	fmt.Println(f2("456"))
}