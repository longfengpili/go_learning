/*
* @Author: chunyang.xu
* @Date:   2020-12-15 15:47:09
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-15 15:50:15
*/

package main
import "fmt"

type number struct {
	f float32
}

type nr number

func main() {
	a := number{5.0}
	b := nr{5.0}
	var c = number(b)

	fmt.Println(a, b, c)
}
