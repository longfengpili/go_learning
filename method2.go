/*
* @Author: chunyang.xu
* @Date:   2020-12-16 15:59:51
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-16 16:01:46
*/

package main

import "fmt"

type InVector []int

func (v InVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func main() {
	fmt.Println(InVector{1, 2, 3}.Sum())
}
