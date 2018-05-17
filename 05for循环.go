/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 9:43
* @version： V1.0
*/
package main

import "fmt"

func main(){

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for ; ;  {
		fmt.Println("Hi Jerryhax")
		break
	}


}
