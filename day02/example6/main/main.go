package main

//交换两个数的值
import(
	"fmt"
)
//方法一
func swapNum(a *int, b *int){
	temp := *a
	*a = *b
	*b = temp
	return
}

//方法二

func swapNum2(a int, b int) (int, int){
	return b, a
}

func main(){
	a := 10
	b := 20

	swapNum(&a, &b)

	fmt.Println("a= ", a)
	fmt.Println("b= ", b)

	a, b = swapNum2(a,b)
	fmt.Printf("a= %d,b= %d\n",a,b)
}

