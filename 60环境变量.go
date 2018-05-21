/**
*
* @Description: 环境变量是一个在为 Unix 程序传递配置信息的普遍方式
* @author：Jerryhax
* @date：2018/5/21 21:30
* @version： V1.0
*/

package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	//使用 os.Setenv 来设置一个键值对。使用 os.Getenv获取一个键对应的值。如果键不存在，将会返回一个空字符串。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	//使用 os.Environ 来列出所有环境变量键值队。这个函数会返回一个 KEY=value 形式的字符串切片。你可以使用strings.Split 来得到键和值。这里我们打印所有的键。
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
