/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/18 17:59
* @version： V1.0
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	//当使用 os.Exit 时 defer 将不会 执行，
	defer fmt.Println("!")
	//注意，不像例如 C 语言，Go 不使用在 main 中返回一个整数来指明退出状态。如果你想以非零状态退出，那么你就要使用 os.Exit。
	os.Exit(3)
}
