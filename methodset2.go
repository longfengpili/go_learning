/*
* @Author: chunyang.xu
* @Date:   2020-12-18 10:37:19
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-18 11:30:35
*/

package main

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i < end; i++ {
		a.Append(i)
	}
}

type lener interface {
	Len() int
}

func LongEnough(l lener) bool {
	return l.Len() * 10 > 42
}

func main() {
	var lst List
	// CountInto(lst, 1, 10)  # error
	if LongEnough(lst) {
		fmt.Printf("- lst is not enough \n")
	} else {
		fmt.Printf("- lst is : %v\n", lst)
	}

	plst := new(List)
	CountInto(plst, 1, 10)
	fmt.Println(plst)
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough \n")
	}
}

// 在 lst 上调用 CountInto 时会导致一个编译器错误，因为 CountInto 需要一个 Appender，而它的方法 Append 只定义在指针上。 在 lst 上调用 LongEnough 是可以的，因为 Len 定义在值上。
// 在 plst 上调用 CountInto 是可以的，因为 CountInto 需要一个 Appender，并且它的方法 Append 定义在指针上。 在 plst 上调用 LongEnough 也是可以的，因为指针会被自动解引用。
