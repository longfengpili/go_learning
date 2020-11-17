/*
* @Author: chunyang.xu
* @Date:   2020-11-16 18:01:42
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-11-16 18:06:21
*/


package main

import "fmt"

func main() {
	for i :=1; i <= 100; i++ {
		switch {
		case i % 3 == 0 && i % 5 == 0:
			fmt.Printf("FizzBuzz\n")
		case i % 3 == 0:
			fmt.Printf("Fizz\n")
		case i % 5 == 0:
			fmt.Printf("Buzz\n")
		default:
			fmt.Printf("%d\n", i)
		}
	}
}
