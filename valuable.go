/*
* @Author: chunyang.xu
* @Date:   2020-12-17 17:03:21
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-17 17:08:17
*/

package main

import "fmt"

type stockPosition struct {
	ticker string
	sharePrice float32
	count float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	var o valuable = stockPosition{"goog", 577.20, 4}
	showValue(o)

	o = car{"BMW", "M3", 66500}
	showValue(o)
}
