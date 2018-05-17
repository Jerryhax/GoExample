/**
* 
* @Description:
Go 的通道选择器 让你可以同时等待多个通道操作。Go 协程和通道以及选择器的结合是 Go 的一个强大特性。
* @author：Jerryhax
* @date：2018/5/17 17:04
* @version： V1.0
*/
package main

import (
	"time"
	"fmt"
)

func main (){
	c1 := make(chan string)
	c2 := make(chan string)
	fmt.Println(time.Now())
	//各个通道将在若干时间后接收一个值，这个用来模拟例如并行的 Go 协程中阻塞的 RPC 操作
	go func(){

		time.Sleep(time.Second*1)
		c1 <- "one"
	}()
	go func (){
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	//我们使用 select 关键字来同时等待这两个值，并打印各自接收到的值。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- c1 :
			fmt.Println("received",msg1)
		case msg2 := <- c2 :
			fmt.Println("received",msg2)
		}
	}
	fmt.Println(time.Now())
}


