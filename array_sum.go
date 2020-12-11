/*
* @Author: chunyang.xu
* @Date:   2020-12-11 14:53:32
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-11 15:03:03
*/

package main
import "fmt"

func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array)
	fmt.Printf("The sum of then array is : %f", x)
}


func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}
