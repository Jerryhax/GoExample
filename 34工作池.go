/**
* 
* @Description: 使用 Go 协程和通道实现一个工作池 。
* @author：Jerryhax
* @date：2018/5/17 19:06
* @version： V1.0
*/
package main

import (
	"fmt"
	"time"
)
//						从通道jobs读任务				结果写入通道result
func worker2(id int,jobs <- chan int,results chan <- int){
	for j := range jobs{
		fmt.Println("worker2",id,"process job",j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main(){
	//为了使用 worker 工作池并且收集他们的结果，我们需要2 个通道。
	jobs := make(chan int, 100)
	results := make(chan int,100)
	fmt.Println(time.Now())
	//这里启动了 3 个 worker，初始是阻塞的，因为还没有传递任务。
	for k := 1;k <= 3 ;k++  {
		go worker2(k,jobs,results)
	}

	//这里我们发送 9 个 jobs，然后 close 这些通道来表示这些就是所有的任务了。
	for j := 1;j <= 9 ;j++  {
		jobs <- j
	}
	close(jobs)

	//最后，我们收集所有这些任务的返回值。
	for a := 1; a <= 9; a++ {
		<-results

	}
	fmt.Println(time.Now())//显示任务执行开始时刻和结束时刻,结果表明使用了3秒左右,而不是9秒,因为三个worker在并行执行.
}
