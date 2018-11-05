package main

import(
	"golang_study/day01/package_example/calc"
	"fmt"
)

func main(){
	sum := calc.Add(100,200)
	sub := calc.Sub(200,10)

	fmt.Println("sum= ",sum)
	fmt.Println("sub= ", sub)
}