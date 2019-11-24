package runtime

import (
	"../util"
	"unsafe"
)
/**
Type Conversion Instructions
133 (0x85) i2l
134 (0x86) i2f
135 (0x87) i2d
136 (0x88) l2i
137 (0x89) l2f
138 (0x8a) l2d
139 (0x8b) f2i
140 (0x8c) f2l
141 (0x8d) f2d
142 (0x8e) d2i
143 (0x8f) d2l
144 (0x90) d2f
145 (0x91) i2b
146 (0x92) i2c
147 (0x93) i2s
 */



type I2L struct {
}

func (self *I2L) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	n := frame.OprandStack.PopInt()
	l := int64(n)
	frame.OprandStack.PushLong(l)
}


type I2F struct {
}

func (self *I2F) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	n := frame.OprandStack.PopInt()
	f := float32(n)
	frame.OprandStack.PushFloat(f)
}

type I2D struct {
}

func (self *I2D) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	n := frame.OprandStack.PopInt()
	d := float64(n)
	frame.OprandStack.PushDouble(d)
}


type L2I struct {
}

func (self *L2I) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	n := frame.OprandStack.PopLong()
	l := int32(n)
	frame.OprandStack.PushInt(l)
}


type L2F struct {
}

func (self *L2F) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	n := frame.OprandStack.PopLong()
	f := float32(n)
	frame.OprandStack.PushFloat(f)
}

type L2D struct {
}

func (self *L2D) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	n := frame.OprandStack.PopLong()
	d := float64(n)
	frame.OprandStack.PushDouble(d)
}



type F2I struct {
}

func (self *F2I) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	v1 := frame.OprandStack.PopFloat()
	v2 := int32(v1)
	frame.OprandStack.PushInt(v2)
}


type F2L struct {
}

func (self *F2L) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	v1 := frame.OprandStack.PopFloat()
	v2 := int64(v1)
	frame.OprandStack.PushLong(v2)
}

type F2D struct {
}

func (self *F2D) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	f := frame.OprandStack.PopFloat()
	d := float64(f)
	frame.OprandStack.PushDouble(d)
}

type D2I struct {
}

func (self *D2I) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	v1 := frame.OprandStack.PopDouble()
	v2 := int32(v1)
	frame.OprandStack.PushInt(v2)
}


type D2L struct {
}

func (self *D2L) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	v1 := frame.OprandStack.PopDouble()
	v2 := int64(v1)
	frame.OprandStack.PushLong(v2)
}

type D2F struct {
}

func (self *D2F) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	v1 := frame.OprandStack.PopDouble()
	v2 := float32(v1)
	frame.OprandStack.PushFloat(v2)
}


// integer -> byte
type I2B struct {
}

func (self *I2B) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	n := frame.OprandStack.PopInt()
	b := util.JByte(n)
	frame.OprandStack.PushInt(int32(b))
}

type I2C struct {
}

func (self *I2C) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	n := int16(frame.OprandStack.PopInt())
	c := *(*util.Char)(unsafe.Pointer(&n))
	frame.OprandStack.PushU4(util.U4(c))
}

type I2S struct {
}

func (self *I2S) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
	n := frame.OprandStack.PopInt()
	l := util.Short(n)
	frame.OprandStack.PushInt(int32(l))
}

