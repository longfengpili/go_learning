/*
* @Author: chunyang.xu
* @Date:   2020-11-03 10:38:28
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-11-03 10:48:23
*/

package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d \n", a)
	}

	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d \n", r)
	}

	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f \n", 100 * rand.Float32())
	}
}