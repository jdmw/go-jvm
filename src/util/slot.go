package util

import (
	"math"
)


type Slot struct {
	Data U4
	Ref Reference
}
type SlotTable []Slot
type SlotStack struct {
	position U2
	table SlotTable
}


func NewSlotTable(length U2) SlotTable{
	return make([]Slot,length)
}

func (self *SlotTable ) SetInt(index U2,num int32){
	(*self)[index].Data = U4(num)
}
func (self *SlotTable ) GetInt(index U2) int32{
	num := (*self)[index].Data
	ui := *(*int32)(Reference(&num))
	return ui
}

func (self *SlotTable ) SetU4(index U2,num U4){
	(*self)[index].Data = num
}
func (self *SlotTable ) GetU4(index U2) U4{
	return (*self)[index].Data
}



func (self *SlotTable ) SetFloat(index U2,num float32){
	(*self)[index].Data= U4(math.Float32bits(num))
}
func (self *SlotTable ) GetFloat(index U2) float32{
	return math.Float32frombits(uint32((*self)[index].Data))
}

func (self *SlotTable) SetRef(index U2,ref Reference) {
	(*self)[index].Ref = ref
}

func (self *SlotTable) GetRef(index U2) Reference {
	return (*self)[index].Ref
}

func (self *SlotTable ) SetU8(index U2,num U8){
	(*self)[index].Data = (U4)(num)
	(*self)[index+1].Data = (U4)(num>>32)
}
func (self *SlotTable ) GetU8(index U2) U8{
	high := U8((*self)[index+1].Data)
	low := U8((*self)[index].Data)
	return high << 32 +  low
}

func (self *SlotTable ) SetLong(index U2,num int64){
	(*self)[index].Data = (U4)(num)
	(*self)[index+1].Data = (U4)(num>>32)
}
func (self *SlotTable ) GetLong(index U2) int64{
	high := U8((*self)[index+1].Data)
	low := U8((*self)[index].Data)
	u8 := high << 32 +  low
	return *(*int64)(Reference(&u8))
}

func (self *SlotTable ) SetDouble(index U2,num float64){
	self.SetU8(index,U8(math.Float64bits(num)))
}
func (self *SlotTable ) GetDouble(index U2) float64{
	return math.Float64frombits(uint64(self.GetU8(index)))
}

func (self *Slot) GetInt() U4{
	return self.Data
}

func NewSlotStack(length U2) *SlotStack  {
	return &SlotStack{0,NewSlotTable(length)}
}

func (self *SlotStack) PushInt(num int32) {
	if(int(self.position) >= len(self.table)){
		panic("the stack is already full")
	}
	self.table.SetInt(self.position,num )
	self.position++
}

func (self *SlotStack) PushU4(num U4) {
	if(int(self.position) >= len(self.table)){
		panic("the stack is already full")
	}
	self.table.SetU4(self.position,num )
	self.position++
}

func (self *SlotStack) PushU8(num U8) {
	if(int(self.position) + 1 >= len(self.table)){
		panic("the stack don't have enough space")
	}
	self.table.SetU8(self.position,num )
	self.position = self.position+2
}
func (self *SlotStack) PushU4_2(high U4,low U4) {
	if(int(self.position) + 1 >= len(self.table)){
		panic("the stack don't have enough space")
	}
	self.table.SetU4(self.position,high )
	self.table.SetU4(self.position+1,low )
	self.position = self.position+2
}


func (self *SlotStack) PushLong(num int64) {
	if(int(self.position) + 1 >= len(self.table)){
		panic("the stack don't have enough space")
	}
	self.table.SetLong(self.position,num )
	self.position = self.position+2
}

func (self *SlotStack) PushRef(ref Reference) {
	if(int(self.position) >= len(self.table)){
		panic("the stack is already full")
	}
	self.table.SetRef(self.position,ref )
	self.position++
}

func (self *SlotStack) PushFloat(num float32) {
	self.PushU4(U4(math.Float32bits(num)))
}

func (self *SlotStack) PushDouble(num float64) {
	self.PushU8(U8(math.Float64bits(num)))
}

func (self *SlotStack) PopU4() U4 {
	if( self.position == 0){
		panic("stack is empty ")
	}
	self.position--
	result := self.table.GetU4(self.position)
	return result
}
func (self *SlotStack) PopInt() int32 {
	if( self.position == 0){
		panic("stack is empty ")
	}
	self.position--
	result := self.table.GetInt(self.position)
	return result
}
func (self *SlotStack) PopRef() Reference {
	if( self.position == 0){
		panic("stack is empty ")
	}
	self.position--
	result := self.table.GetRef(self.position)
	return result
}
func (self *SlotStack) PopU8() U8{
	if( self.position <= 1){
		panic("the stack operation error,it's not long enough to contain a 64 bits number")
	}
	self.position = self.position-2
	result := self.table.GetU8(self.position)
	return result
}
func (self *SlotStack) PopLong() int64{
	if( self.position <= 1){
		panic("the stack operation error,it's not long enough to contain a 64 bits number")
	}
	self.position = self.position-2
	result := self.table.GetLong(self.position)
	return result
}

func (self *SlotStack) PopFloat() float32 {
	return math.Float32frombits(uint32(self.PopInt()))
}

func (self *SlotStack) PopDouble() float64{
	return math.Float64frombits(uint64(self.PopLong()))
}

func (self *SlotStack) Reset() {
	self.position = 0
}

func SlotTest()  {
	i1 := int32(-1)
	i2 :=  int32(2)
	n1 := U4(99)
	n2 :=  U4(2123)
	f := float32(1.213 )
	u8 := U8(n1) << 32 + U8(n2)
	l := int64(i2 << 32) + int64(i1)
	d := float64(l) + float64(f)
	stack := NewSlotStack(20)
	stack.PushInt(i1)
	stack.PushInt(i2)
	stack.PushU4(n1)
	stack.PushU4(n2)
	stack.PushFloat(f)
	stack.PushLong(l)
	stack.PushDouble(d)
	stack.PushU8(u8)
	println(u8 == stack.PopU8())
	println(d == stack.PopDouble())
	println(l == stack.PopLong())
	println(f == stack.PopFloat())
	println(n2 == stack.PopU4())
	println(n1 == stack.PopU4())
	println(i2 == stack.PopInt())
	println(i1 == stack.PopInt())
}

/**
 note  ; type convert

int -> uint32 -> int
```go
    i := int(-1) // type : int
	iu := uint32(i) // type : uint32
	ui := *(*int)(Reference(&i)) // type : int
	fmt.Println(ui == i) // true
```

uint32 -> float32 : math.Float32frombits

int -> pointer -> int
```go
    i = int(1)
	obj := &i
	ptr := uintptr(Reference(obj))
	ref := int(ptr) // pointer address
	ptr2 := Reference(uintptr(ref))
	intp := (*int)(ptr2)
	fmt.Println(*intp) // print 1
```

 */