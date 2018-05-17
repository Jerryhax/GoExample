/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 9:14
* @version： V1.0
*/
package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main(){
	fmt.Println(s)
	//const 语句可以出现在任何 var 语句可以出现的地方
	const n = 500000000

	//常数表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)

	//数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如一次显示的类型转化。
	//这里的 math.Sin函数需要一个 float64 的参数。
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}