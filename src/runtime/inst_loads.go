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
	if( ref == nil){
		frame.SetException(EXP_NULL_POINT)
	}
	arr := *((*[]int32)(unsafe.Pointer(ref)))
	if( index <0 || index >= int32(len(arr))){
		frame.SetException(EXP_ARRAY_INDEX_OUT_OF_BOUNDS)
	}
	frame.OprandStack.PushInt(arr[index] )
}

type BALOAD struct {
}

func (self *BALOAD) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := frame.OprandStack.PopInt()
	ref :=  frame.OprandStack.PopRef()
	if( ref == nil){
		frame.SetException(EXP_NULL_POINT)
	}
	arr := *((*[]byte)(unsafe.Pointer(ref))) // byte array or boolean array
	if( index <0 || index >= int32(len(arr))){
		frame.SetException(EXP_ARRAY_INDEX_OUT_OF_BOUNDS)
	}
	value := arr[index]
	frame.OprandStack.PushInt(int8(value))
}



func InstLoadTest(){
	reader := util.NewBigEndianReader(make([]byte,10))

	// test iaload
	frame := NewStackFrame(10,10)
	iarr := []int{-2,-1,0,1,2,3,4}
	frame.OprandStack.PushRef(unsafe.Pointer(&iarr))
	frame.OprandStack.PushInt(1)
	iaload.execute(reader,frame)
	rst := int(frame.OprandStack.PopInt())
	fmt.Println(iarr[1] == rst  )



}
