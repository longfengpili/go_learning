/*
* @Author: chunyang.xu
* @Date:   2020-12-15 14:38:59
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-15 14:54:00
*/

package main
import "fmt"

type struct1 struct {
	i1 int
	f1 float32
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf(" The int is : %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
    fmt.Printf("The string is: %s\n", ms.str)
    fmt.Println(ms)
    
    ms1 := &struct1{10, 15.5, "Chris"}
    fmt.Printf("%v\n", ms1)

    var ms2 struct1
    ms2 = struct1{10, 15.5, "Chris"}
    fmt.Printf("%v\n", ms2)

    var ms3 *struct1
    ms3 = &struct1{10, 15.5, "Chris"}
    fmt.Printf("%v\n", ms3)
    fmt.Printf("%v\n", *ms3)
}
