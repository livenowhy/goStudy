package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch1 <- 1

	cc := <- ch1
	fmt.Println(cc)
	/** fatal error: all goroutines are asleep - deadlock! */

	chrw := make(chan int)    // 创建无缓冲 channel
	chor := make(<-chan int)  // 创建只读channel
	chow := make(chan<- int)  // 创建只写channel
}


