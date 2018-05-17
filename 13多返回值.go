/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 12:34
* @version： V1.0
*/
package main

import "fmt"

func vals()(int,int){
	return  3,7
}

func main (){

	a,b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)



}