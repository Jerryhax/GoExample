/**
* 
* @Description: 标准库的 strings 包提供了很多有用的字符串相关的函数。
* @author：Jerryhax
* @date：2018/5/18 15:36
* @version： V1.0
*/
package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {

	//注意他们都包中的函数，不是字符串对象自身的方法，这意味着我们需要考虑在调用时传递字符作为第一个参数进行传递。
	p("Contains:	",s.Contains("test","es"))
	p("Count:	",s.Count("test","t"))
	p("hasPrefix:	",s.HasPrefix("test","te"))
	p("HasSuffix:	",s.HasSuffix("test","st"))
	p("Index:	",s.Index("test","e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
