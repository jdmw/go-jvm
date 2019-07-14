package main

import "fmt"

type info struct {
	num int
}
type infos []info

var i int

func (self *infos) parse() {
	fo := []info{info{1 },info{2}}
	for i,e := range fo {
		fmt.Println(&e == &fo[i])
		e.num += 1
		fo[i] = e
	}
	*self = fo
	fmt.Printf("%v\n",fo)
}


func main1(){
	arr := infos{}
	arr.parse()
	fmt.Printf("%v\n",arr)
}
