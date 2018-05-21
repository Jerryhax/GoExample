/**
*
* @Description:
SHA1 散列经常用生成二进制文件或者文本块的短标识。
* @author：Jerryhax
* @date：2018/5/18 17:37
* @version： V1.0
*/

package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
)

func checkerr(e error) {
	if e != nil{
		panic(e)
	}
}

func main()  {
	//开始，这里是展示如写入一个字符串（或者只是一些字节）到一个文件
	d1 := []byte("\thello\t\n\tjerryhax\n")
	err := ioutil.WriteFile("D:/tmp/dat1",d1,0644)
	checkerr(err)
	//对于更细粒度的写入，先打开一个文件。
	f, err := os.Create("D:/tmp/dat2")
	checkerr(err)
	//打开文件后，习惯立即使用 defer 调用文件的 Close操作。
	defer f.Close()
	//你可以写入你想写入的字节切片
	d2 := []byte{115,111,109,101,10}
	n2,err := f.Write(d2)
	checkerr(err)
	fmt.Printf("wrote %d bytes\n",n2)
	//WriteString 也可以用来写
	n3,err := f.WriteString("Writes\n")
	fmt.Printf("wrote %d bytes \n",n3)

	//调用 Sync 来将缓冲区的信息写入磁盘
	f.Sync()

	//bufio 提供了和我们前面看到的带缓冲的读取器一样的带缓冲的写入器。
	w:= bufio.NewWriter(f)
	n4,err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes \n",n4)
	//使用 Flush 来确保所有缓存的操作已写入底层写入器。
	w.Flush()





}