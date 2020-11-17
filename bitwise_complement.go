/*
* @Author: chunyang.xu
* @Date:   2020-11-16 17:59:49
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-11-16 18:00:44
*/

package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("the complement of %b is: %b\n", i, ^i)
	}
}
