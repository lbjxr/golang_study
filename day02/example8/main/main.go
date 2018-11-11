package main

import (
	"fmt"
)

func main(){
	var str = "Hello World\n"

	var str1 = `
	反引号可以
	换行输入
	会保存格式
	`

	fmt.Println(str)
	fmt.Println(str1)
}