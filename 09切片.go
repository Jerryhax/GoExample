/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/17 10:45
* @version： V1.0
*/
package main

import (
	"fmt"
)

func main (){
	//make (类型，长度)
	s := make([]string,3)
	fmt.Println("empty:",s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:",s)
	fmt.Println("get:",s[2])

	fmt.Println("len",len(s))
	//append追加元素并返回新的slice，此处注意接收新生成的slice
	s = append(s,"d")
	s = append(s,"e","f")
	fmt.Println("apd:",s)
	//Slice 也可以被 copy。这里我们创建一个空的和 s 有相同长度的 slice c，并且将 s 复制给 c。
	c :=  make([]string, len(s))
	copy(c,s)
	fmt.Println("copy:",c)

	//Slice 支持通过 slice[low:high] 语法进行“切片”操作。包含前边不包含后面
	// 例如，这里得到一个包含元素 s[2], s[3],s[4] 的 slice。包含是s[2],不包含s[5]

	l := s[2:5]
	fmt.Println("sl1",l)
	//从开始到5
	l = s[:5]
	fmt.Println("sl2",l)
	//从2到最后
	l = s[2:]
	fmt.Println("sl3",l)

	t := []string{"g","h","i"}
	fmt.Println("dcl:",t)

	//Slice 可以组成多维数据结构。内部的 slice 长度可以不同，这和多维数组不同。
	twoDimension := make([][]int,3)
	for i := 0 ;i < 3;i++  {
		innerLen := i + 1
		twoDimension[i] = make([]int,innerLen)
		for j :=0;j < innerLen ;j++  {
			twoDimension[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoDimension)



}

