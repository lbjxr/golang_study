package main

//import "time"

//测试helloworld
func main() {
	
	//helloworld()

	//测试gorouter 轻量级线程

	/*
	for j := 0; j < 100; j++ {
		go for_func(j,j + 1)
	}

	time.Sleep(10*time.Second)
	*/

	//pipe的使用

	go pipe_func_1()
}