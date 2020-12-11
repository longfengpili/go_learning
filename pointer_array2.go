/*
* @Author: chunyang.xu
* @Date:   2020-12-11 14:14:22
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-11 14:17:58
*/

package main
import "fmt"

func fp(a *[3]int) {
	a[1] = 5
	fmt.Println(a)
}

func main() {
	var arr [3]int
	for i := 0; i < 3; i++ {
		arr = [3]int{i, i * i, i * i * i}
		fp(&arr)
		fmt.Println(arr)
	}
}
