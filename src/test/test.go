package main

import (
	"../runtime"
	"fmt"
)

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


func main(){
	//util.SlotTest()
	runtime.InstLoadTest()
/*	i := int(-1)
	iu := uint32(i)
	ui := *(*int)(unsafe.Pointer(&i))
	fmt.Println(unsafe.Pointer(&iu) == unsafe.Pointer( &i))
	fmt.Println(ui == i)
	fmt.Println(ui)
	obj := &i
	ptr := uintptr(unsafe.Pointer(obj))
	ref := int(ptr)
	ptr2 := unsafe.Pointer(uintptr(ref))
	intp := (*int)(ptr2)
	fmt.Println(*intp)*/
	//arr := infos{}
	//arr.parse()
	//fmt.Printf("%v\n",arr)
}
