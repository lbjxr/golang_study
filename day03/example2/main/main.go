package main

import "fmt"

//函数的相关说明

//声明一个函数类型
type add_func func(int, int) int

func add(a, b int) int{
	return a + b
}

//第一个参数是函数类型
func operator(op add_func, a int, b int) int {
	return op(a,b)
}

func main(){
	//定一个函数变量，并将add函数的地址付给c
	c := add

	fmt.Println(c)
	sum := operator(c, 100, 200)
	fmt.Println(sum)
}