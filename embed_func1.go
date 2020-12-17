/*
* @Author: chunyang.xu
* @Date:   2020-12-16 17:30:11
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-16 17:36:39
*/

package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	name string
	log *Log
}

func main() {
	c := new(Customer)
	c.name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can !"
	// c = &Customer{"Barak Obama", &Log{"1 - Yes we can !"}}
	c.Log().Add("2 - After me the world will be a better place!")
	fmt.Println(c.Log())
}

func (l *Log) Add (s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}
