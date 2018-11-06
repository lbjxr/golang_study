package pkg1

import(
	"fmt"
	//引用pkg2包，但不使用
	_ "golang_study/day02/example3/pkg2"
)

//全局变量，初始化值
var Name string = "This is pkg1 Name"
var Age int = 1

//init函数初始化
func init(){

	//对变量赋值
	Name = "This is pkg1 init Name"
	Age = 11

	fmt.Println("name: " , Name)
	fmt.Println("age: ", Age)
}