/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 17:49
* @version： V1.0
*/
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool){
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main(){

	done := make(chan bool,1)
	go worker(done)

	//go worker(done)是协程,并行执行,直到其执行完,done通道才会有值, 语句<- done才能执行,main才能结束
	<- done
}
