/**
* 
* @Description: panic 意味着有些出乎意料的错误发生。通常我们用它来表示程序正常运行中不应该出现的，或者我们没有处理好的错误。
* @author：Jerryhax
* @date：2018/5/18 14:19
* @version： V1.0
*/
package main

import "os"

func main() {
	panic("a problem")

	_,err :=os.Create("/tmp/file")
	if err != nil{
		panic(err)
	}

}