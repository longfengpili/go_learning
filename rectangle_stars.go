/*
* @Author: chunyang.xu
* @Date:   2020-11-16 18:15:36
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-11-16 18:22:47
*/


package main

import "fmt"

func main() {
	w, h := 20, 10
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
