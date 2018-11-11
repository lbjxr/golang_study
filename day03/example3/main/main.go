package main

import "fmt"

//可变参数的练习
func add(arg...int) int{
	var sum = arg[0]
	for i := 1; i < len(arg); i++ {
		sum += arg[i]
	}

	return sum
}

func main(){
	sum := add(10,10,33)

	fmt.Println(sum)
}

