/*
* @Author: chunyang.xu
* @Date:   2020-11-16 17:22:00
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-11-16 17:31:51
*/


package main

import (
	"fmt"
)


func main() {
	var m int64
	m = 6
	switch {
	case m >=1 && m <= 3:
		fmt.Printf("spring: %d", m)
	case m > 3 && m <= 6:
		fmt.Printf("summer: %d", m)
	case m > 6 && m <= 9:
		fmt.Printf("autumn: %d", m)
	default:
		fmt.Printf("winter: %d", m)
	}
}
