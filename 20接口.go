/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 14:48
* @version： V1.0
*/
package main

import (
	"math"
	"fmt"
)

type geometry interface{
	area() float64
	perim() float64
}

type rectange struct {
	width,height float64
}

type circle struct {
	radius float64
}
/*要在 Go 中实现一个接口，我们只需要实现接口中的所有方法。
这里我们让 rect 实现了 geometry 接口。*/
func (r rectange) area() float64{
	return r.width * r.height
}
func (r rectange) perim() float64{
	return  2 * r.width + 2 * r.height
}

func (c circle) area() float64{
	return math.Pi * c.radius *  c.radius
}
func (c circle) perim() float64{
	return  2 * math.Pi * c.radius
}
//如果一个变量的是接口类型，那么我们可以调用这个被命名的接口中的方法。
// 这里有一个一通用的 measure 函数，利用这个特性，它可以用在任何 geometry 上。
func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main () {
	r := rectange{width:3,height:4}
	c := circle{radius:5}

	measure(r)
	measure(c)
}
