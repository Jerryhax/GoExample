/**
* 
* @Description: 超时 对于一个连接外部资源，或者其它一些需要花费执行时间的操作的程序而言是很重要的。得益于通道和 select，在 Go中实现超时操作是简洁而优雅的。
* @author：Jerryhax
* @date：2018/5/17 18:02
* @version： V1.0
*/
package main

import (
	"time"
	"fmt"
)

func main(){

	c1 := make(chan string ,1)
	//两秒后出结果
	go func() {
		time.Sleep(time.Second*2)
		c1 <- "result 1"
	}()

	select {
	case res:= <-c1:
		fmt.Println(res)
		//一秒后出结果
	case <-time.After(time.Second*1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string,1)
	go func() {
		//两秒
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	//三秒
	case <- time.After(time.Second * 3):
		fmt.Println("timeout 2")

	}
//所以选择结果是timeo1和result2

}
