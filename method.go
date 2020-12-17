/*
* @Author: chunyang.xu
* @Date:   2020-12-16 15:31:13
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-16 15:35:57
*/

package main

import "fmt"

type TwoInts struct {
	a, b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is : %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(10))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is : %d\n", two2.AddThem())
	fmt.Printf("Add them to the param: %d\n", two2.AddToParam(10))
}


func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.AddThem() + param
}
