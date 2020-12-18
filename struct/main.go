/*
* @Author: chunyang.xu
* @Date:   2020-12-16 11:31:53
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-17 18:30:20
*/

package main
import (
	"fmt"
	"./struct_pack"
)

func main() {
	struct1 := new(struct_test.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16.0
	fmt.Println(struct1)
}
