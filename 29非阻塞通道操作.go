/**
* 
* @Description: 常规的通过通道发送和接收数据是阻塞的。然而，我们可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select。
* @author：Jerryhax
* @date：2018/5/17 17:22
* @version： V1.0
*/
package main

import "fmt"

func main (){
	messages := make(chan string)
	signals := make(chan bool)
	//这里是一个非阻塞接收的例子。如果在 messages 中存在，然后 select 将这个值带入 <-messages case中。
	// 如果不是，就直接到 default 分支中。
	select {
	case msg := <-messages:
		fmt.Println("received message",msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message" ,msg)
	default:
		fmt.Println("no massage sent")
	}

	//我们可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。这里我们试图在 messages和 signals 上同时使用非阻塞的接受操作。
	select {
	case msg := <- messages:
		fmt.Println("received message",msg)
	case sig := <-signals:
		fmt.Println("received signal",sig)
	default:
		fmt.Println("no activity")
	}



}

