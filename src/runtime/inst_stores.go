package runtime

import (
	"../util"
	"unsafe"
)
/**
 * instrument
 *  ?store : pop var out of Oprand Stack and then store into local variable

Stores
54 (0x36) istore
55 (0x37) lstore
56 (0x38) fstore
57 (0x39) dstore
58 (0x3a) astore
59 (0x3b) istore_0
60 (0x3c) istore_1
61 (0x3d) istore_2
62 (0x3e) istore_3
63 (0x3f) lstore_0
64 (0x40) lstore_1
65 (0x41) lstore_2
66 (0x42) lstore_3
67 (0x43) fstore_0
68 (0x44) fstore_1
69 (0x45) fstore_2
70 (0x46) fstore_3
71 (0x47) dstore_0
72 (0x48) dstore_1
73 (0x49) dstore_2
74 (0x4a) dstore_3
75 (0x4b) astore_0
76 (0x4c) astore_1
77 (0x4d) astore_2
78 (0x4e) astore_3
79 (0x4f) iastore
80 (0x50) lastore
81 (0x51) fastore
82 (0x52) dastore
83 (0x53) aastore
84 (0x54) bastore
85 (0x55) castore
86 (0x56) sastore
 */


type U4STORE struct {
}

/**
	The index is an unsigned byte that must be an index into the local
	variable array of the current frame (§2.6). The value on the top
	of the operand stack must be of type int. It is popped from the
	operand stack, and the value of the local variable at index is set
	to value.
 */
func (self *U4STORE) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := reader.ReadU1()
	num := frame.OprandStack.PopU4()
	frame.LocalVariables.SetU4(util.U2(index),num )
}

type U4STORE_N struct {
	index util.U1
}

/**
The index is an unsigned byte that must be an index into the local
variable array of the current frame (§2.6). The value on the top
of the operand stack must be of type int. It is popped from the
operand stack, and the value of the local variable at index is set
to value.
*/
func (self *U4STORE_N) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := self.index
	num := frame.OprandStack.PopU4()
	frame.LocalVariables.SetU4(util.U2(index),num )
}


type U8STORE struct {
}

/**
The index is an unsigned byte that must be an index into the local
variable array of the current frame (§2.6). The value on the top
of the operand stack must be of type int. It is popped from the
operand stack, and the value of the local variable at index is set
to value.
*/
func (self *U8STORE) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := reader.ReadU1()
	num := frame.OprandStack.PopU8()
	frame.LocalVariables.SetU8(util.U2(index),num )
}

type U8STORE_N struct {
	index util.U1
}

/**
The index is an unsigned byte that must be an index into the local
variable array of the current frame (§2.6). The value on the top
of the operand stack must be of type int. It is popped from the
operand stack, and the value of the local variable at index is set
to value.
*/
func (self *U8STORE_N) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	index := self.index
	num := frame.OprandStack.PopU8()
	frame.LocalVariables.SetU8(util.U2(index),num )
}


type IASTORE struct {
}

func (self *IASTORE) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	value := frame.OprandStack.PopInt()
	index := frame.OprandStack.PopInt()
	arrayRef := frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,arrayRef)){
		return
	}
	array := *(*util.Ints)(unsafe.Pointer(arrayRef))
	if(false == checkArrayLen(frame,len(array),index) ){
		return
	}
	array[int(index)] = value
}


type LASTORE struct {
}

func (self *LASTORE) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	value := frame.OprandStack.PopLong()
	index := frame.OprandStack.PopInt()
	arrayRef := frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,arrayRef)){
		return
	}
	array := *(*util.Longs)(unsafe.Pointer(arrayRef))
	if(false == checkArrayLen(frame,len(array),index) ){
		return
	}
	array[int(index)] = value
}


type FASTORE struct {
}

func (self *FASTORE) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	value := frame.OprandStack.PopFloat()
	index := frame.OprandStack.PopInt()
	arrayRef := frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,arrayRef)){
		return
	}
	array := *(*util.FLoats)(unsafe.Pointer(arrayRef))
	if(false == checkArrayLen(frame,len(array),index) ){
		return
	}
	array[int(index)] = value
}


type DASTORE struct {
}

func (self *DASTORE) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	value := frame.OprandStack.PopDouble()
	index := frame.OprandStack.PopInt()
	arrayRef := frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,arrayRef)){
		return
	}
	array := *(*util.Doubles)(unsafe.Pointer(arrayRef))
	if(false == checkArrayLen(frame,len(array),index) ){
		return
	}
	array[int(index)] = value
}


type AASTORE struct {
}

func (self *AASTORE) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	value := frame.OprandStack.PopRef()
	index := frame.OprandStack.PopInt()
	arrayRef := frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,arrayRef)){
		return
	}
	array := *(*util.References)(unsafe.Pointer(arrayRef))
	if(false == checkArrayLen(frame,len(array),index) ){
		return
	}
	array[int(index)] = value
}


type BASTORE struct {
}

func (self *BASTORE) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	value := byte(frame.OprandStack.PopInt())
	index := frame.OprandStack.PopInt()
	arrayRef := frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,arrayRef)){
		return
	}
	array := *(*util.Bytes)(unsafe.Pointer(arrayRef))
	if(false == checkArrayLen(frame,len(array),index) ){
		return
	}
	array[int(index)] = value
}


type CASTORE struct {
}

func (self *CASTORE) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	value := util.Char(frame.OprandStack.PopU4())
	index := frame.OprandStack.PopInt()
	arrayRef := frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,arrayRef)){
		return
	}
	array := *(*util.Chars)(unsafe.Pointer(arrayRef))
	if(false == checkArrayLen(frame,len(array),index) ){
		return
	}
	array[int(index)] = value
}


type SASTORE struct {
}

func (self *SASTORE) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	value := util.Short(frame.OprandStack.PopInt())
	index := frame.OprandStack.PopInt()
	arrayRef := frame.OprandStack.PopRef()
	if(!checkNullPointer(frame,arrayRef)){
		return
	}
	array := *(*util.Shorts)(unsafe.Pointer(arrayRef))
	if(false == checkArrayLen(frame,len(array),index) ){
		return
	}
	array[int(index)] = value
}




