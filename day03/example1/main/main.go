package main

import (
	"strings"
	"fmt"
)

//strings中的HasPrefix和HasSuffix函数的使用

func urlProcess(url string) string{
	result := strings.HasPrefix(url, "http://")

	if !result {
		url = fmt.Sprintf("http://%s", url)
	}

	return url
}

func pathProcess(path string) string{
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/",path)
	}

	return path
}

func main(){

	var (
		url string
		path string
	)

	fmt.Scanf("%s%s", &url,&path)
	url = urlProcess(url)
	path = pathProcess(path)

	fmt.Print(url)
	fmt.Println(path)
}