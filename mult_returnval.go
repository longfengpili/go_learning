/*
* @Author: chunyang.xu
* @Date:   2020-11-17 17:56:03
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-11-17 18:07:39
*/


package main

import (
	"fmt"
)

func SumProductDiff(i, j int) (int, int, int) {
	return i + j, i * j, i - j
}

func SumProductDiffN(i, j int) (s int, p int, d int) {
	s, p, d = i+j, i*j, i-j
	return
}

func main() {
	sum, prod, diff := SumProductDiff(3, 4)
	fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)
	sum, prod, diff = SumProductDiffN(3, 4)
	fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)
}
