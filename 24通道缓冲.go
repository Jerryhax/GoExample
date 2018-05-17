/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 16:41
* @version： V1.0
*/
package main

import "fmt"

func main (){
	//make了一个通道,指定缓存三个值
	messages := make(chan string,3)

	messages <- "buffered"
	messages <- "channnel"
	messages <- "jerryhax"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}