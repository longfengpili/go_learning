/*
* @Author: chunyang.xu
* @Date:   2020-11-16 18:28:05
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-11-16 18:34:27
*/


package main

import "fmt"

func main() {
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s + "a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}

