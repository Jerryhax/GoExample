/**
* 
* @Description: 打点器 则是当你想要在固定的时间间隔重复执行
* @author：Jerryhax
* @date：2018/5/17 18:51
* @version： V1.0
*/
package main

import (
	"time"
	"fmt"
)

func main (){
	//打点器和定时器的机制有点相似：一个通道用来发送数据。这里我们在这个通道上使用内置的 range 来迭代值每隔500ms 发送一次的值。
	ticker := time.NewTicker(time.Millisecond*500)
	go func() {
		for t := range ticker.C{
			fmt.Println("Tick at ",t)
		}
	}()

	//打点器可以和定时器一样被停止。一旦一个打点停止了，将不能再从它的通道中接收到值。程序在2501毫秒后停止这个打点器,此时打点5次
	time.Sleep(time.Millisecond*2501)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

