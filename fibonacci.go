/*
* @Author: chunyang.xu
* @Date:   2020-11-19 17:14:57
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-11-23 18:07:13
*/


package main

import (
	"fmt"
	"time"
)

const LIM = 41

func main() {
	start := time.Now()
	result := 0
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Println(i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
