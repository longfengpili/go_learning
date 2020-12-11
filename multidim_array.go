/*
* @Author: chunyang.xu
* @Date:   2020-12-11 14:19:06
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-11 14:26:30
*/

package main

import "fmt"

const (
	WIDTH = 8
	HEIGHT = 5
)

type pixel int
var screen [WIDTH][HEIGHT]pixel

func main() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
	fmt.Println(screen)
}
