/*
* @Author: chunyang.xu
* @Date:   2020-12-17 16:43:15
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-17 16:55:52
*/

package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	fmt.Printf("The square has area: %f\n", sq1.Area())

	var areaIntf Shaper
	areaIntf = sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
