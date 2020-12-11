/*
* @Author: chunyang.xu
* @Date:   2020-11-16 18:28:05
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-11-23 18:39:51
*/


package main

import "fmt"

func main() {
	var arr1 [3]int
	var arr2 [3]int
	ar1(arr1)
	fmt.Println(arr1)
	ar2(&arr2)
	fmt.Println(arr2)
}

func ar1(arr [3]int) {
	arr[0] = 1
	fmt.Println(arr)
}

func ar2(arr *[3]int) {
	arr[0] = 10
	fmt.Println(arr)
}

