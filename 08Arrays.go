/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 10:29
* @version： V1.0
*/
package main

import "fmt"

func main (){
	//数组默认是零值的，对于 int 数组来说也就是 0。
	var a [5]int
	fmt.Println("emp:",a)

	a[4] = 100
	fmt.Println("set;",a)
	fmt.Println("get:",a[4])
	//内置函数 len 返回数组的长度
	fmt.Println("len:",len(a))
	//这个语法在一行内初始化一个数组
	b := [5]int {1,2,3,4,5}
	fmt.Println("dcl:",b)


	//数组的存储类型是单一的，但是你可以组合这些数据来构造多维的数据结构。
	var twoD [2][3]int
	for  i := 0;i < 2; i++{
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	//在使用 fmt.Println 来打印数组的时候，会使用[v1 v2 v3 ...] 的格式显示
	fmt .Println("2d: ", twoD)



}
