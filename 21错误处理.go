/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 15:18
* @version： V1.0
*/
package main

import (
	"errors"
	"fmt"
)
//按照惯例，错误通常是最后一个返回值并且是 error 类型，一个内建的接口。
func  f1(arg int)(int,error){
	if arg == 42 {
		//errors.New 构造一个使用给定的错误信息的基本error 值
		return -1,errors.New("can't work with 42")
	}
	return arg +  3,nil
}


//通过实现 Error 方法来自定义 error 类型是可以的。这里使用自定义错误类型来表示上面的参数错误。
type argError struct {
	arg int
	prob string
}
func (e *argError) Error()string{
	return fmt.Sprintf("%d - %s",e.arg,e.prob)
}
func f2(arg int)(int,error){
	if arg == 42{
		return -1,&argError{arg,"can't work with it"}

	}
	//返回错误值为 nil 代表没有错误
	return  arg + 3,nil
}

func main (){
	for _,i := range []int{7,42}{
		if  r,e := f1(i);e != nil{
			fmt.Println("f1 failed:",e)
		} else{
			fmt.Println("f1 worked:",r)
		}
	}
	for _, i := range []int{7,42}{
		if r,e := f2(i);e != nil{
			fmt.Println("f2 failed:",e)
		}else{
			fmt.Println("f2 worked:",r)
		}
	}

	_,e := f2(42)
	if ae, ok :=e.(*argError); ok{
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}


}
