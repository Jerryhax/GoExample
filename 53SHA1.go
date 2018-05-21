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
	"crypto/sha1"//Go 在多个 crypto/* 包中实现了一系列散列函数。
	"fmt"
)

func main() {
	s := "shaq this string"
	//产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。
	// 1,这里我们从一个新的散列开始。
	h := sha1.New()
	//2,写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(s))
	//3,这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)

	fmt.Println(s)
	//SHA1 值经常以 16 进制输出。使用%x 来将散列结果格式化为 16 进制字符串。
	fmt.Printf("%x\n",bs)
	//可以使用和上面相似的方式来计算其他形式的散列值。例如，计算 MD5 散列，引入 crypto/md5 并使用 md5.New()方法
}
