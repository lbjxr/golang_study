package main

import(
	"fmt"
	//使用一个别名代替包名
	a "golang_study/day02/example2/add"
)

func main(){

	fmt.Println("Name = ", a.Name)
	fmt.Println("Age = ", a.Age)
}