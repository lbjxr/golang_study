package pkg2

import(
	"fmt"
)

//初始化全局变量，不同包下的变量名可以相同
var Name string = "This is pkg2 name"
var Age int = 2

//init函数
func init()  {
	
	//修改变量值
	Name = "This is pkg2 init name"
	Age = 22

	fmt.Println("name: ", Name)
	fmt.Println("age: ", Age)
}
