/*
* @Author: chunyang.xu
* @Date:   2020-12-16 11:19:48
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-16 11:29:00
*/

package main
import "fmt"

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	// ok
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1
	fmt.Printf("%v\n", y)

	// not ok
	// z := make(Bar)
	// (*y).thingOne = "hello"
	// (*y).thingTwo = 1
	
	// ok
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"
	fmt.Println(x)

	// not ok
	// u := new(Foo)
	// (*u)["x"] = "goodbye"
	// (*u)["y"] = "world"
}

// 试图 make() 一个结构体变量，会引发一个编译错误，这还不是太糟糕，但是 new() 一个 map 并试图向其填充数据，将会引发运行时错误！ 因为 new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存。所以在使用 map 时要特别谨慎。
