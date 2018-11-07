package main
//值传递和引用传递
import (
	"fmt"
)

func test(b *int){
	*b = 10
}

func main(){
	a := 5
	b := make(chan int,2)

	fmt.Println("a= ",a)
	fmt.Println("b= ",b)

	test(&a)
	fmt.Println("test 后的 a= ", a)
}