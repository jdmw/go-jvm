package main

import "fmt"

type info struct {
	num int
}
type infos []info

var i int

func (self *infos) parse() {
	fo := info{i }
	i++
	*self =  append(*self,fo)
}
func main(){
	arr := infos{}
	arr.parse()
	arr.parse()
	arr.parse()
	arr.parse()
	fmt.Printf("%v\n",arr)
}
