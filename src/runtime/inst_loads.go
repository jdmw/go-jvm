package runtime

import (
	"../util"
	"fmt"
	"unsafe"
)
/**
 * instrument
 *  ?load :  load  var from local variable and then push into Operand Stack
 *  ?store : pop var out of Oprand Stack and then store into local variable

Loads
21 (0x15) iload
22 (0x16) lload
23 (0x17) fload
24 (0x18) dload
25 (0x19) aload
26 (0x1a) iload_0
27 (0x1b) iload_1
28 (0x1c) iload_2
29 (0x1d) iload_3
30 (0x1e) lload_0
31 (0x1f) lload_1
32 (0x20) lload_2
33 (0x21) lload_3
34 (0x22) fload_0
35 (0x23) fload_1
36 (0x24) fload_2
37 (0x25) fload_3
38 (0x26) dload_0
39 (0x27) dload_1
40 (0x28) dload_2
41 (0x29) dload_3
42 (0x2a) aload_0
43 (0x2b) aload_1
44 (0x2c) aload_2
45 (0x2d) aload_3
46 (0x2e) iaload
47 (0x2f) laload
48 (0x30) faload
49 (0x31) daload
50 (0x32) aaload
51 (0x33) baload
52 (0x34) caload
53 (0x35) saload
 */

//iload_0  ~ iload_3

type U4LOAD struct {
}

func (self *U4LOAD) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := reader.ReadU1()
	num := frame.LocalVariables.GetU4(util.U2(index))
	frame.OprandStack.PushU4(num )
}

type U4LOAD_N struct {
	num util.U2
}

func (self *U4LOAD_N) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	num := frame.LocalVariables.GetU4(self.num)
	frame.OprandStack.PushU4(num )
}


type U8LOAD struct {
}

func (self *U8LOAD) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := reader.ReadU1()
	num := frame.LocalVariables.GetU8(util.U2(index))
	frame.OprandStack.PushU8(num )
}

type U8LOAD_N struct {
	num util.U2
}

func (self *U8LOAD_N) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	num := frame.LocalVariables.GetU8(self.num)
	frame.OprandStack.PushU8(num )
}

type ALOAD struct {
}

func (self *ALOAD) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := reader.ReadU1()
	ref := frame.LocalVariables.GetRef(util.U2(index))
	frame.OprandStack.PushRef(ref )
}

type ALOAD_N struct {
	num util.U2
}

func (self *ALOAD_N) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	num := frame.LocalVariables.GetRef(self.num)
	frame.OprandStack.PushRef(num )
}

type IALOAD struct {
}


func (self *IALOAD) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := frame.OprandStack.PopInt()
	ref :=  frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,ref)) {
		return
	}
	arr := *((*util.Ints)(unsafe.Pointer(ref)))
	if( false == checkArrayLen(frame,len(arr),index)){
		return
	}
	frame.OprandStack.PushInt(arr[index] )
}

type LALOAD struct {
}


func (self *LALOAD) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := frame.OprandStack.PopInt()
	ref :=  frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,ref)) {
		return
	}
	arr := *((*[]int64)(unsafe.Pointer(ref)))
	if( false == checkArrayLen(frame,len(arr),index)){
		return
	}
	frame.OprandStack.PushLong(arr[index] )
}

type FALOAD struct {
}


func (self *FALOAD) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := frame.OprandStack.PopInt()
	ref :=  frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,ref)) {
		return
	}
	arr := *((*util.FLoats)(unsafe.Pointer(ref)))
	if( false == checkArrayLen(frame,len(arr),index)){
		return
	}
	frame.OprandStack.PushFloat(arr[index] )
}

type DALOAD struct {
}


func (self *DALOAD) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := frame.OprandStack.PopInt()
	ref :=  frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,ref)) {
		return
	}
	arr := *((*util.Doubles)(unsafe.Pointer(ref)))
	if( false == checkArrayLen(frame,len(arr),index)){
		return
	}
	frame.OprandStack.PushDouble(arr[index] )
}

type AALOAD struct {
}

func (self *AALOAD) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := frame.OprandStack.PopInt()
	ref :=  frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,ref)) {
		return
	}
	arr := *((*util.References)(unsafe.Pointer(ref))) // byte array or boolean array
	if( false == checkArrayLen(frame,len(arr),index)){
		return
	}
	frame.OprandStack.PushRef(arr[index])
}

type BALOAD struct {
}

func (self *BALOAD) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := frame.OprandStack.PopInt()
	ref :=  frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,ref)) {
		return
	}
	arr := *((*[]byte)(unsafe.Pointer(ref))) // byte array or boolean array
	if( false == checkArrayLen(frame,len(arr),index)){
		return
	}
	value := arr[index]
	frame.OprandStack.PushInt(int32(int8(value)))
}


type CALOAD struct {
}

func (self *CALOAD) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := frame.OprandStack.PopInt()
	ref :=  frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,ref)) {
		return
	}
	arr := *((*util.Chars)(unsafe.Pointer(ref))) // byte array or boolean array
	if( false == checkArrayLen(frame,len(arr),index)){
		return
	}
	frame.OprandStack.PushU4(util.U4(arr[index]))
}

type SALOAD struct {
}


func (self *SALOAD) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := frame.OprandStack.PopInt()
	ref :=  frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,ref)) {
		return
	}
	arr := *((*util.Shorts)(unsafe.Pointer(ref))) // byte array or boolean array
	if( false == checkArrayLen(frame,len(arr),index)){
		return
	}
	frame.OprandStack.PushInt(int32(int16(arr[index])))
}

func InstLoadTest(){
	reader := util.NewBigEndianReader(make([]byte,10))

	// test iaload
	frame := NewStackFrame(10,10)
	iarr := util.Ints{-2,-1,0,1,2,3,4}
	frame.OprandStack.PushRef(util.Reference(unsafe.Pointer(&iarr)))
	frame.OprandStack.PushInt(1)
	iaload.execute(reader,frame)
	rst := int(frame.OprandStack.PopInt())
	fmt.Println(int(iarr[1]) == rst  )

}

func checkNullPointer(frame *StackFrame,ref util.Reference )  bool {
	if( ref == nil){
		frame.SetException(EXP_NULL_POINT)
		return false
	}
	return true
}

func checkArrayLen(frame *StackFrame, arrlen int, index int32) bool{
	if( index <0 || index >= int32(arrlen)){
		frame.SetException(EXP_ARRAY_INDEX_OUT_OF_BOUNDS)
		return false
	}
	return true
}
