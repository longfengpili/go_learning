/*
* @Author: chunyang.xu
* @Date:   2020-11-16 17:44:13
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-11-16 17:59:25
*/


package main

import (
	"fmt"
)

func main() {
	// 1 - use 2 nested for loops
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
	// 2 -  use only one for loop and string concatenation
	str := "G"
	for i := 1; i <= 25; i++ {
		println(str)
		str += "G"
	}
}
