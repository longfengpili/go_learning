/*
* @Author: chunyang.xu
* @Date:   2020-12-15 15:27:37
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-15 15:41:48
*/

package main
import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Println(pers1)

	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	upPerson(pers2)
	fmt.Println(*pers2)

	pers3 := &Person{"Chris", "Woodward"}
	upPerson(pers3)
	fmt.Println(pers3)
}
