package main

import "fmt"

func pipe_func_1()  {
	pipe := make(chan int, 10)

	pipe <- 1
	pipe <- 2

	var b1, b2 int 
	b1 = <- pipe
	fmt.Println(b1)
	b2 = <- pipe
	fmt.Println(b2)

	//fmt.Println(pipe_func_2(b1, b2))
}

func pipe_func_2(a1 int, a2 int) (int){

	fmt.Println(a1,a2)

	sum := a1 + a2
	fmt.Println(sum)

	return a1 + a2
}