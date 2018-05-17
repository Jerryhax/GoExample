/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 10:00
* @version： V1.0
*/
package main

import "fmt"

func main (){

	//go语言可以不用小括号，没有三目运算
	if 7%2 == 0{
		fmt.Println("7 is even")
	}else{
		fmt.Println("7 is odd")
	}

	if 8%4 == 0{
		fmt.Println("8 is divisible by 4")
	}
	//在条件语句之前可以有一个语句；任何在这里声明的变量都可以在所有的条件分支中使用。
	if num := 9;num < 0{
		fmt.Println(num,"is negative")
	}else if num < 10{
		fmt.Println(num,"has 1 digit")
	}else{
		fmt.Println(num,"has multiple digits")
	}

}

