package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 如果不关闭会阻塞在这里面等第三个值,崩溃 ？
	for elem := range queue {
		fmt.Println(elem)
	}

}
