/*
* @Author: chunyang.xu
* @Date:   2020-11-16 17:40:14
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-11-16 17:43:34
*/


package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println(i)
	}


	i := 0
START:
	fmt.Println(i)
	i++
	if i < 15 {
		goto START
	}

}
