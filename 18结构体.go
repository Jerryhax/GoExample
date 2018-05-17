/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 14:17
* @version： V1.0
*/
package main

import "fmt"

type person struct {
	name string
	age int
}

func main (){
	//这个语法创建了一个新的结构体元素。
	fmt.Println(person{"Bob",20})

	//在初始化一个结构体元素时指定字段名字。
	fmt.Println(person{name:"Alice",age:30})
	fmt.Println(person{age:30 ,name:"Alice"})
	//省略的字段将被初始化为零值。
	fmt.Println(person{name:"Fred"})
	//& 前缀生成一个结构体指针。
	fmt.Println(&person{name:"Ann",age:40})

	//使用点来访问结构体字段。
	s := person{name :"Sean",age:50}
	s.age = 17
	fmt.Println(s.name,"\t",s.age)

	//也可以对结构体指针使用. - 指针会被自动解引用
	sp := &s
	fmt.Println(sp.age)
	//结构体是可变的。
	sp.age = 51
	fmt.Println(sp.age)



}
