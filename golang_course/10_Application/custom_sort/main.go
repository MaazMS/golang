package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}
type ByAge []person

func (a ByAge) Len() int           { return len(a) } // first p1 return
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []person

func (b ByName) Len() int           { return len(b) } // first p1 return
func (b ByName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByName) Less(i, j int) bool { return b[i].Name < b[j].Name }
func main() {

	p1 := person{"Maaz", 22}
	p2 := person{"Amir", 25}
	p3 := person{"Shaikh", 35}
	p4 := person{"irfan", 5}

	//fmt.Println(p1.Name, p1.Age)
	//fmt.Println(p2.Name, p2.Age)
	//fmt.Println(p3.Name, p3.Age)
	//fmt.Println(p4.Name, p4.Age)

	p := []person{p1, p2, p3, p4}
	sort.Sort(ByAge(p))
	fmt.Println(p)

	sort.Sort(ByName(p))
	fmt.Println(p)

}
