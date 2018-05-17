/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 10:14
* @version： V1.0
*/
package main

import (
	"fmt"
	"time"
)

func main (){

	i :=2
	fmt.Print("write ",i," as ")
	switch  i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday(){
	//在一个 case 语句中，你可以使用逗号来分隔多个表达式
	case time .Saturday,time.Sunday:
		fmt.Println("it is the weekend")
	default:
		fmt.Println("it is a weekday")
	}

	//不带表达式的 switch 是实现 if/else 逻辑的另一种方式
	t :=time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it is after noon")

	}



}
