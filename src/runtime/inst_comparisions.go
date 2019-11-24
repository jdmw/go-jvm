package runtime

import (
	"../util"
	"math"
)
/**
Comparisons
148 (0x94) lcmp
149 (0x95) fcmpl
150 (0x96) fcmpg
151 (0x97) dcmpl
152 (0x98) dcmpg
153 (0x99) ifeq
154 (0x9a) ifne
155 (0x9b) iflt
156 (0x9c) ifge
157 (0x9d) ifgt
158 (0x9e) ifle
159 (0x9f) if_icmpeq
160 (0xa0) if_icmpne
161 (0xa1) if_icmplt
162 (0xa2) if_icmpge
163 (0xa3) if_icmpgt
164 (0xa4) if_icmple
165 (0xa5) if_acmpeq
166 (0xa6) if_acmpne
 */

/**
 * Operation Compare long
 *    Operand Stack
 *..., value1, value2 →
 *..., result
 * If value1 is greater than value2, the int value 1 is pushed onto the operand stack.
 * If value1 is equal to value2, the int value 0 is pushed onto the operand stack.
 * If value1 is less than value2, the int value -1 is pushed onto the operand stack.
 */

type LCMP struct {
}

func (self *LCMP) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value2 := frame.OprandStack.PopLong()
	value1 := frame.OprandStack.PopLong()
	result := compareInt64(value1-value2)
	frame.OprandStack.PushInt(result)
}

type FCMP_OP struct {
	valueWhenNaN int32
}

func (self *FCMP_OP) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	var result int32
	value2 := frame.OprandStack.PopFloat()
	value1 := frame.OprandStack.PopFloat()
	if ( math.IsNaN(float64(value1)) || math.IsNaN(float64(value2))) {
		result = self.valueWhenNaN
	}else{
		result = compareFloat(float64(value1-value2))
	}
	frame.OprandStack.PushInt(result)
}

type DCMP_OP struct {
	valueWhenNaN int32
}

func (self *DCMP_OP) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	var result int32
	value2 := frame.OprandStack.PopDouble()
	value1 := frame.OprandStack.PopDouble()
	if ( math.IsNaN(float64(value1)) || math.IsNaN(float64(value2))) {
		result = self.valueWhenNaN
	}else{
		result = compareFloat(value1-value2)
	}
	frame.OprandStack.PushInt(result)
}



type IF_COND struct {
	comp _IComp
}

func (self *IF_COND) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value := frame.OprandStack.PopInt()
	if (self.comp.compare(value)) {
		offset := reader.ReadU2()
		reader.SetOffset(util.U4(offset))
	}
}

type IF_ICMP_COMD struct {
	comp _IComp
}

func (self *IF_ICMP_COMD) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value2 := frame.OprandStack.PopInt()
	value1 := frame.OprandStack.PopInt()
	if (self.comp.compare(value1 - value2)) {
		offset := reader.ReadU2()
		reader.SetOffset(util.U4(offset))
	}
}

type IF_ACMP_COMD struct {
	eq bool
}

func (self *IF_ACMP_COMD) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value2 := frame.OprandStack.PopRef()
	value1 :=  frame.OprandStack.PopRef()
	if ((value1 == value2) == self.eq ) {
		offset := reader.ReadU2()
		reader.SetOffset(util.U4(offset))
	}
}


type IFNULL struct {
	isNull bool
}

func (self *IFNULL) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value := frame.OprandStack.PopRef()
	if ((value == nil) == self.isNull ) {
		offset := reader.ReadU2()
		reader.SetOffset(util.U4(offset))
	}
}




func compareInt64( sub int64) int32{
	var res int32
	if sub != 0 {
		if sub > 0 {
			sub = 1
		}else{
			sub = -1
		}
	}else{
		res = 0
	}
	return  res
}

func compareFloat(sub float64 ) int32{
	var res int32
	if sub != 0 {
		if sub > 0 {
			sub = 1
		}else{
			sub = -1
		}
	}else{
		res = 0
	}
	return  res
}



type _IComp interface  {
	compare(value int32) bool ;
}

type eq struct {}
func (self eq) compare(value int32 ) bool  {
	return value == 0
}


type ne struct {}
func (self ne) compare(value int32 ) bool  {
	return value != 0
}


type lt struct {}
func (self lt) compare(value int32 ) bool  {
	return value < 0
}


type le struct {}
func (self le) compare(value int32 ) bool  {
	return value <= 0
}


type gt struct {}
func (self gt) compare(value int32 ) bool  {
	return value > 0
}


type ge struct {}
func (self ge) compare(value int32 ) bool  {
	return value >= 0
}
