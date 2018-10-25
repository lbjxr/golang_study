package main

import "fmt"
import "time"

func main(){

	//一个基本的switch
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//在一个case语句中，可以用逗号分割多个表达式
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It`s the weekend")
	default:
		fmt.Println("It`s a Weekday")
	}
}
