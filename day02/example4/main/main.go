package main

import(
	"fmt"
	"time"
)

//定义两个常量

const (
	Man = 1
	Female = 2
)

func main(){

	for {

		//获取当前时间秒数
		second := time.Now().Unix()
		
		if (second % Female == 0){
			fmt.Println("female",second)
		}else {
			fmt.Println("man")
		}

		time.Sleep(1000*time.Millisecond)
	}
}