/*
* @Author: chunyang.xu
* @Date:   2020-12-16 16:37:41
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-16 16:45:43
*/

package main

import "fmt"

type B struct {
	thing int
}

func (b *B) change() {b.thing = 1}

func (b B) write() string {
	b.thing = 10
	return fmt.Sprint(b)
}

func main() {
	var b1 B
	b1.change()
	fmt.Println(b1.write())
	fmt.Println(b1)

	b2 := new(B)
	b2.change()
	fmt.Println(b2.write())
	fmt.Println(b2)
}
