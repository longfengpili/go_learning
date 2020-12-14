/*
* @Author: chunyang.xu
* @Date:   2020-12-14 10:40:44
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-14 10:58:47
*/

package main
import "fmt"

func bubbleSort(l []int) (nl []int) {
	nl = make([]int, len(l))
	copy(nl, l)
	// nl = append(nl, l...)
	for i := 0; i < len(nl)-1; i++ {
		for j := 0; j < len(nl)-1; j++ {
			if nl[j] > nl[j+1] {
				nl[j], nl[j+1] = nl[j+1], nl[j]
			}
		}
	}
	return
}

func main() {
	l := []int{3, 2, 1, 6, 5}
	nl := bubbleSort(l)
	fmt.Println(nl)
}
