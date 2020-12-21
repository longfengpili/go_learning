/*
* @Author: chunyang.xu
* @Date:   2020-12-15 15:27:37
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-18 11:55:28
*/

package main
import (
	"fmt"
	// "strings"
)

type Intlist []int

func (l Intlist) Len() int {
	return len(l)
}

func main() {
	data := []int{1, 4, 3, 2}
	il := Intlist(data)
	fmt.Println(il)
}
