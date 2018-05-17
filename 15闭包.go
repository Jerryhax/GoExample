/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 13:51
* @version： V1.0
*/
package main

import "fmt"

func intSeq() func()int {

	i := 0
	return func() int {
		i += 1
		return  i
	}
}

func main (){
	//我们调用 intSeq 函数，将返回值（也是一个函数）赋给nextInt。
	// 这个函数的值包含了自己的值 i，这样在每次调用 nextInt 时都会更新 i 的值。
	nextInt :=intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()

	fmt.Println(newInts())



}
