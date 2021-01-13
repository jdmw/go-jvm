package runtime

import (
	"../util"
	"unsafe"
)
/**
References Control
167 (0xa7) goto
168 (0xa8) jsr
169 (0xa9) ret
170 (0xaa) tableswitch
171 (0xab) lookupswitch
172 (0xac) ireturn
173 (0xad) lreturn
174 (0xae) freturn
175 (0xaf) dreturn
176 (0xb0) areturn
177 (0xb1) return
 */


type GOTO struct {
}

func (self *GOTO) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	reader.SetOffset(util.U4(int16(reader.ReadU2())))
}

type GOTO_W struct {
}

func (self *GOTO_W) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	reader.SetOffset(reader.ReadU4())
}


type JSR struct {
}

func (self *JSR) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	address := int16(reader.ReadU2())
	frame.OprandStack.PushInt(int32(address))
}

type JSR_W struct {
}

func (self *JSR_W) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	address := int32(reader.ReadU4())
	frame.OprandStack.PushInt(address)
}

/**
 * Return from subroutine
 */
type RET struct {
}

func (self *RET) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	address := frame.LocalVariables.GetU4(index(reader,wideMode))
	reader.SetOffset(address) //  write into the Java Virtual Machine's pc register
}



/**
 * Access jump table by index and j
* Format : OP
		<0~3 bytes  padding>
		<default (int32)>
		<low bytes (int32)>
		<high bytes (int32)>
		<jump offsets...>
  Operand Stack
		..., index →
		...
 */
type TABLE_SWITCH struct {
}

func (self *TABLE_SWITCH) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	reminder := int(reader.Offset()) % 4
	if( reminder > 0 ){
		padding := 4 - reminder
		reader.SetOffset(reader.Offset() + util.U4(padding ))
	}
	defaultValueAddress := reader.ReadU4()
	low := reader.ReadInt32()
	high := reader.ReadInt32()

	index := frame.OprandStack.PopInt()
	if index < low || index > high {
		reader.SetOffset(defaultValueAddress)
	}else{
		position := index - low + int32(reader.Offset())
		reader.SetOffset(util.U4(position))
		address := reader.ReadU4()
		reader.SetOffset(address)
	}
}

/**
 * Access jump table by key match and jump
* Format : OP
		<0~3 bytes  padding>
		<default (int32)>
		<npairs (int32)>
		<jmatch-offset pairs.....>
  Operand Stack
		..., key →
		...
*/
type LOOKUP_SWITCH struct {
}

func (self *LOOKUP_SWITCH) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	reminder := int(reader.Offset()) % 4
	if( reminder > 0 ){
		padding := 4 - reminder
		reader.SetOffset(reader.Offset() + util.U4(padding ))
	}
	defaultValueAddress := reader.ReadU4()

	key := frame.OprandStack.PopInt()
	for nparis := reader.ReadInt32() ; nparis > 0 ; nparis-- {
		match := reader.ReadInt32()
		offset := reader.ReadInt32()
		if ( key == match){
			reader.SetOffset(util.U4(offset))
			return
		}
	}
	reader.SetOffset(defaultValueAddress)
}





type RETURN struct {
}

func (self *RETURN) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	frame.thread.stack.Pop()
	frame.SetReturn()
}

type IRETURN struct {
}

func (self *IRETURN) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	value := frame.OprandStack.PopInt()
	frame.thread.stack.Pop()
	frame.thread.stack.Top().OprandStack.PushInt(value)
	frame.SetReturn()
}






type FRETURN struct {
}

func (self *FRETURN) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	value := frame.OprandStack.PopFloat()
	frame.thread.stack.Pop()
	frame.thread.stack.Top().OprandStack.PushFloat(value)
	frame.SetReturn()
}




type DRETURN struct {
}

func (self *DRETURN) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	value := frame.OprandStack.PopDouble()
	frame.thread.stack.Pop()
	frame.thread.stack.Top().OprandStack.PushDouble(value)
	frame.SetReturn()
}




type LRETURN struct {
}

func (self *LRETURN) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	value := frame.OprandStack.PopLong()
	frame.thread.stack.Pop()
	frame.thread.stack.Top().OprandStack.PushLong(value)
	frame.SetReturn()
}




type ARETURN struct {
}

func (self *ARETURN) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	value := frame.OprandStack.PopRef()
	frame.thread.stack.Pop()
	frame.thread.stack.Top().OprandStack.PushRef(value)
	frame.SetReturn()
}


