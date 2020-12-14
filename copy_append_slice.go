/*
* @Author: chunyang.xu
* @Date:   2020-12-14 10:12:48
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-14 10:19:21
*/

package main
import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slFrom := array[1:3]
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	array[2] = 10
	fmt.Println(array)
	fmt.Println(slTo)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
