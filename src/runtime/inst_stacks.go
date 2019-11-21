package runtime

/**
Stack
87 (0x57) pop
88 (0x58) pop2
89 (0x59) dup
90 (0x5a) dup_x1
91 (0x5b) dup_x2
92 (0x5c) dup2
93 (0x5d) dup2_x1
94 (0x5e) dup2_x2
95 (0x5f) swap

 */

import (
	"../util"
)


type POP struct {
}

func (self *POP) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	frame.OprandStack.PopU4()
}


type POP2 struct {
}

func (self *POP2) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	frame.OprandStack.PopU8()
}

/**
 * Duplicate the top operand stack val
 */
type DUP struct {
}

func (self *DUP) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	num := frame.OprandStack.PopU4()
	frame.OprandStack.PushU4(num)
	frame.OprandStack.PushU4(num)
}


type DUP_X1 struct {
}

func (self *DUP_X1) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	value1 := frame.OprandStack.PopU4()
	value2 := frame.OprandStack.PopU4()
	frame.OprandStack.PushU4(value1)
	frame.OprandStack.PushU4(value2)
	frame.OprandStack.PushU4(value1)
}


type DUP_X2 struct {
}

func (self *DUP_X2) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	value1 := frame.OprandStack.PopU4()
	value2 := frame.OprandStack.PopU4()
	value3 := frame.OprandStack.PopU4()
	frame.OprandStack.PushU4(value1)
	frame.OprandStack.PushU4(value3)
	frame.OprandStack.PushU4(value2)
	frame.OprandStack.PushU4(value1)
}

type DUP2 struct {
}

func (self *DUP2) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	num := frame.OprandStack.PopU8()
	frame.OprandStack.PushU8(num)
	frame.OprandStack.PushU8(num)
}


type DUP2_X1 struct {
}

func (self *DUP2_X1) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	value1 := frame.OprandStack.PopU4()
	value2 := frame.OprandStack.PopU4()
	value3 := frame.OprandStack.PopU4()
	frame.OprandStack.PushU4(value2)
	frame.OprandStack.PushU4(value1)
	frame.OprandStack.PushU4(value3)
	frame.OprandStack.PushU4(value2)
	frame.OprandStack.PushU4(value1)
}


type DUP2_X2 struct {
}

func (self *DUP2_X2) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	value1 := frame.OprandStack.PopU8()
	value2 := frame.OprandStack.PopU8()
	frame.OprandStack.PushU8(value1)
	frame.OprandStack.PushU8(value2)
	frame.OprandStack.PushU8(value1)
}


type SWAP struct {
}

func (self *SWAP) execute(reader *util.BigEndianReader,frame *StackFrame)  {
	value1 := frame.OprandStack.PopU4()
	value2 := frame.OprandStack.PopU4()
	frame.OprandStack.PushU4(value1)
	frame.OprandStack.PushU4(value2)
}



