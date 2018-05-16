/**
* @author：Jerryhax
* @date：2018/5/16 19:19
* @version： V1.0
* @Description: TODO
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//rand.Intn(n)生成一个0-n之间的随机整数：序列是一样的，按序列往后取
	fmt.Print(rand.Intn(100),",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	fmt.Print(rand.Intn(100),",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	fmt.Print(rand.Intn(100),",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	fmt.Print(rand.Intn(100),",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	fmt.Print(rand.Intn(100),",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	//rand.Float64()返回一个64位浮点数 f，0.0 <= f <= 1.0
	//序列是一样的，按序列往后取
	fmt.Println(rand.Float64())

	//线性偏移到我们想要的范围
	fmt.Print((rand.Float64()*5)+5,",")
	fmt.Print((rand.Float64()*5)+5)
	fmt.Println()
	fmt.Print((rand.Float64()*5)+5,",")
	fmt.Print((rand.Float64()*5)+5)
	fmt.Println()

	//默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。要产生变化的序列，需要给定一个变化的种子。
	//UnixNano()纳秒时间戳，作为变化的种子
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 :=rand.New(s1)


	fmt.Print(r1.Intn(100),",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	//种子相同
	s2 := rand.NewSource(421)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100),",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 :=rand.NewSource(421)
	r3 :=rand.New(s3)
	fmt.Print(r3.Intn(100),",")
	fmt.Print(r3.Intn(100))

}
