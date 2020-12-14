/*
* @Author: chunyang.xu
* @Date:   2020-11-16 18:28:05
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-11 17:14:29
*/


package main

import "fmt"



func main() {
	items := [...]int{10, 20, 30, 40, 50}
	fmt.Println(items)
	for i, item := range items {
		items[i] = 2 * item
	}
	fmt.Println(items)
}

