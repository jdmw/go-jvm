package runtime

import (
	"../util"
	"../classloader"
	// "../classloader/classfile"
)

const (
	EXP_NULL_POINT = 2  // NullPointerException
	EXP_ARRAY_INDEX_OUT_OF_BOUNDS = 3  // ArrayIndexOutOfBoundsException
)
type Stack struct {
	maxSize uint
	size uint
	top *StackFrame
}

type StackFrame struct {
	thread *Thread
	method *classloader.Method
	instance *Object
	stackedPc util.U4
	errorCode util.U1
	//reader *util.BigEndianReader

	LocalVariables util.SlotTable
	OprandStack *util.SlotStack
	lower *StackFrame
}

func NewStack(maxSize uint) *Stack {
	return &Stack{maxSize:maxSize}
}

func (self *StackFrame) SetReturn(){
	self.errorCode = 1
}
func (self *StackFrame) SetException(exception int){
	self.errorCode = util.U1(exception)
}

func NewStackFrame(maxLocals, maxStack util.U2) *StackFrame {
	return &StackFrame{
		LocalVariables : util.NewSlotTable(maxLocals),
		OprandStack: util.NewSlotStack(maxStack),
		errorCode:0,
		//reader:util.NewBigEndianReader(method.Code),
	}
}

func NewStackFrameByMethod(thread *Thread,method *classloader.Method) *StackFrame {
	if method.IsNative() {
		panic("can't load native method in jvm thread ")
	}
	method.Load()
	frame := NewStackFrame(method.MaxLocals,method.MaxStack)
	frame.method = method
	frame.thread = thread
	return frame
}

func NewStackFrameByInstanceMethod(thread *Thread,instance *Object,method *classloader.Method) *StackFrame {
	frame := NewStackFrameByMethod(thread,method)
	frame.instance = instance
	return frame
}

func (self *Stack) Push(frame *StackFrame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self.top != nil {
		frame.lower = self.top
	}
	self.top = frame
	self.size++
}
func (self *Stack) Pop() *StackFrame {
	if self.top == nil {
		panic("jvm stack is empty")
	}
	t := self.top
	self.top = t.lower
	//t.lower.reader = nil
	t.lower.instance = nil
	t.lower.method = nil
	t.lower = nil
	self.size --
	return t
}
func (self *Stack) Top() *StackFrame {
	if self.top == nil {
		panic("jvm stack is empty")
	}
	return self.top
}
