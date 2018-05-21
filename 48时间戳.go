/**
* 
* @Description: TODO
* @author：Jerryhax
* @date：2018/5/18 16:45
* @version： V1.0
*/
package main

import (
	"time"
	"fmt"
)

func main() {
	now := time.Now()
	//分别使用带 Unix 或者 UnixNano 的 time.Now来获取从自协调世界时起到现在的秒数或者纳秒数。
	secs:= now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)
	//注意 UnixMillis 是不存在的，所以要得到毫秒数的话，你要自己手动的从纳秒转化一下。
	millis :=nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)
	//也可以将协调世界时起的整数秒或者纳秒转化到相应的时间。
	fmt.Println(time.Unix(secs,0))
	fmt.Println(time.Unix(0,nanos))
}
