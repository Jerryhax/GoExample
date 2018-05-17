/**
* 
* @Description:
关闭 一个通道意味着不能再向这个通道发送值了。这个特性可以用来给这个通道的接收方传达工作已经完成的信息。
* @author：Jerryhax
* @date：2018/5/17 17:38
* @version： V1.0
*/
package main

import "fmt"

func main (){
	jobs := make(chan int,5)
	done := make(chan bool)


	go func() {
		for{
			j,more := <- jobs
			if more {
				fmt.Println("received job",j)

			}else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job",j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done


}
