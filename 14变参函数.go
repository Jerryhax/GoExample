/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 13:39
* @version： V1.0
*/
package main

import "fmt"
//这个函数使用任意数目的 int 作为参数。
func sum(nums ...int){
	fmt.Println(nums," ")
	total := 0
	for _, num := range nums {
		total +=  num
	}
	fmt.Println(total)
}

func main(){
	//变参函数使用常规的调用方式，
	sum(1,2)
	sum(1,2,3,4)
	//变参函数使用特殊的调用方式：如果你的 slice 已经有了多个值，想把它们作为变参使用，你要这样调用 func(slice...)。
	nums := []int{1,2,3,4}
	sum(nums...)
}
