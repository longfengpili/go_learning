/*
* @Author: chunyang.xu
* @Date:   2020-12-11 14:09:11
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-11 14:12:48
*/

package main
import "fmt"

func main() {
	// var arrAge = [5]int{18, 20, 15, 22, 16}
	// var arrLazy = [...]int{5, 6, 7, 8, 22}
	// var arrLazy = []int{5, 6, 7, 8, 22}	//注：初始化得到的实际上是切片slice
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}	//注：初始化得到的实际上是切片slice
	
	for i:=0; i < len(arrKeyValue); i++ {
		fmt.Printf("%d is %s\n", i, arrKeyValue[i])
	}

	for i, j := range arrKeyValue {
		fmt.Printf("%d is %s\n", i, j)
	}
}
