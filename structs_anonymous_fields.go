/*
* @Author: chunyang.xu
* @Date:   2020-12-16 14:51:32
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-16 14:55:06
*/

package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int  // anonymous field
	innerS  // anonymous field
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Println(outer)

	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println(outer2)
}
