/**
* 
* @Description: Defer 被用来确保一个函数调用在程序执行结束前执行。同样用来执行一些清理工作。
* @author：Jerryhax
* @date：2018/5/18 14:44
* @version： V1.0
*/
package main

import (
	"os"
	"fmt"
)

func main() {
	f := createFile("D:/WorkSpace/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f,err := os.Create(p)
	if err != nil{
		panic(err)
	}
	return f

}

func writeFile(f *os.File)  {
	fmt.Println("writing")
	fmt.Fprintln(f,"hi jerryhax")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

