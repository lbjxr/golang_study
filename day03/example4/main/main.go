package main

import "fmt"

//defer 压栈，defer后面的变量值保持不变
func main(){
	var i int = 0
	defer fmt.Println(i)
	defer fmt.Println("second")

	i = 10
	fmt.Println(i)
}