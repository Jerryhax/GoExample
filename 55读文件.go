/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/18 18:02
* @version： V1.0
*/
package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"io"
	"bufio"
)
//读取文件需要经常进行错误检查，这个帮助方法可以精简下面的错误检查过程。
func check(e error) {
	if e != nil{
		panic(e)
	}
}
func main() {
	//也许大部分基本的文件读取任务是将文件内容读取到内存中
	dat,err := ioutil.ReadFile("D:/tmp/dat1")
	check(err)
	fmt.Print(string(dat))

	//你经常会想对于一个文件是怎么读并且读取到哪一部分进行更多的控制。
	// 对于这个任务，从使用 os.Open打开一个文件获取一个 os.File 值开始。
	f, err := os.Open("D:/tmp/dat1")
	check(err)

	//从文件开始位置读取一些字节。这里最多读取 5 个字节，并且这也是我们实际读取的字节数。
	b1 :=make([]byte,5)
	n1,err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes:	%s\n",n1,string(b1))

	//你也可以 Seek 到一个文件中已知的位置并从这个位置开始进行读取。
	o2,err := f.Seek(6,0)
	check(err)
	b2 := make([]byte,2)
	n2,err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3,err :=f.Seek(6,0)
	check(err)
	b3 :=make([]byte,2)
	n3,err := io.ReadAtLeast(f,b3,2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_,err = f.Seek(0,0)
	check(err)

	r4 := bufio.NewReader(f)
	b4,err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()


}
