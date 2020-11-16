/*
* @Author: chunyang.xu
* @Date:   2020-11-09 18:43:44
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-11-09 18:48:19
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	// fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	// fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("%d\n", strings.Index(str, "n"))
	fmt.Printf("%d\n", strings.LastIndex(str, "n"))
}
