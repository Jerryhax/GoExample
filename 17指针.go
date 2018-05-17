/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 14:01
* @version： V1.0
*/
package main

import "fmt"

func zeroval(ival int){
	ival = 0
}

//*i 取指针指向的值
func zeroptr(iptr *int){
	*iptr = 0
}

func main (){
	i := 1
	fmt.Println("initial:",i)


	zeroval(i)
	fmt.Println("zeroval:",i)

	//&i取变量i 的地址  其类型是 *int
	zeroptr(&i)
	fmt.Println("zeroptr:",i)

	fmt.Println("pointer:",&i)



}

