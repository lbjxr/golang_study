package main

import(
	"golang_study/day01/goroute_example/goroute"
	"fmt"
)

func main(){
	//声明一个管道
	var pipe chan int
	//为管道分配空间
	pipe = make(chan int, 1)
	go goroute.Add(100,200,pipe)
	//将管道的值取出]]
	result := <- pipe
	fmt.Println("resule= ", result)
}