/**
* 
* @Description: for 和 range为基本的数据结构提供了迭代的功能。我们也可以使用这个语法来遍历从通道中取得的值。
* @author：Jerryhax
* @date：2018/5/17 18:19
* @version： V1.0
*/
package main

import "fmt"

func main(){
	queue := make(chan string,2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem:= range queue {
		fmt.Println(elem)
	}
}
