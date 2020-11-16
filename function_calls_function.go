/*
* @Author: chunyang.xu
* @Date:   2020-11-02 18:12:09
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-11-02 18:23:23
*/


package main

var a string

func main() {
   a = "G"
   print(a)  // G
   f1()  // O G
}

func f1() {
   a := "O"
   print(a)	 // O
   f2()
}

func f2() {
   print(a)	 // G
}
