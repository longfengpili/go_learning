/*
* @Author: chunyang.xu
* @Date:   2020-12-18 10:37:19
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-18 10:43:28
*/

package main

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	l = append(l, val)
}

type Appender struct {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i < end; i++ {
		a.Append(i)
	}
}

type lener interface {
	len() int
}

func LongEnough(l lener) bool {
	return l.Len() * 10 > 42
}

func main() {
	var lst List
	if LongEnough(lst) {
		fmt.Printf("- lst is not enough \n")
	}

	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(lst) {
		fmt.Printf("- plst is long enough \n")
	}
}
