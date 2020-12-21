/*
* @Author: chunyang.xu
* @Date:   2020-12-21 10:47:00
* @Last Modified by:   chunyang.xu
* @Last Modified time: 2020-12-21 11:40:40
*/

package sort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len() - pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n-1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}


type IntArray []int

func (a IntArray) Len() int {return len(a)}
func (a IntArray) Less(i, j int) bool {return a[i] < a[j]}
func (a IntArray) Swap(i, j int) {a[i], a[j] = a[j], a[i]}

type StringArray []string

func (s StringArray) Len() int {return len(s)}
func (s StringArray) Less(i, j int) bool {return s[i] < s[j]}
func (s StringArray) Swap(i, j int) {s[i], s[j] = s[j], s[i]}

func SortInts(a []int) {Sort(IntArray(a))}
func SortStrings(a []string) {Sort(StringArray(a))}

func IntsAreSorted(a []int) bool {return IsSorted(IntArray(a))}
func StringsAreSorted(a []string) bool {return IsSorted(StringArray(a))}
