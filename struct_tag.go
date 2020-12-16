/*
* @Author: chunyang.xu
* @Date:   2020-12-16 14:43:10
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-16 14:49:44
*/

package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool "An important answer"
	field2 string "The name of the thing"
	field3 int "How much there are"
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	// fmt.Printf("%v, %d\n", tt, ix)
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v, %v\n", ttType, ixField.Tag)
}
