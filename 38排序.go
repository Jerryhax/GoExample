/**
* 
* @Description:
Go 的 sort 包实现了内置和用户自定义数据类型的排序功能。我们首先关注内置数据类型的排序。
* @author：Jerryhax
* @date：2018/5/18 13:53
* @version： V1.0
*/
package main

import (
	"sort"
	"fmt"
)
/*排序是原地更新的，所以他会改变给定的序列并且不返回一个新值。*/
func main() {
	strs := []string{"ca","af","ab","am"}
	sort.Strings(strs)
	fmt.Println("Strings:	",strs)


	ints :=[]int{90,1,42,8,2,0,56,23}
	sort.Ints(ints)
	fmt.Println("ints:	",ints)
	//使用 sort 来检查一个序列是不是已经是排好序的。
	s :=sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ",s)


}
