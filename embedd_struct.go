/*
* @Author: chunyang.xu
* @Date:   2020-12-16 14:56:40
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-16 14:59:33
*/

package main

import "fmt"

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by int
}

func main() {
	b := B{A{1, 2}, 3, 4}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}
