package runtime

import (
	"../util"
	"math"
	"unsafe"
)

/**
Math
96 (0x60) iadd
97 (0x61) ladd
98 (0x62) fadd
99 (0x63) dadd
100 (0x64) isub
101 (0x65) lsub
102 (0x66) fsub
103 (0x67) dsub
104 (0x68) imul
105 (0x69) lmul
106 (0x6a) fmul
107 (0x6b) dmul
108 (0x6c) idiv
109 (0x6d) ldiv
110 (0x6e) fdiv
111 (0x6f) ddiv
112 (0x70) irem
113 (0x71) lrem
114 (0x72) frem
115 (0x73) drem
116 (0x74) ineg
117 (0x75) lneg
118 (0x76) fneg
119 (0x77) dneg
120 (0x78) ishl
121 (0x79) lshl
122 (0x7a) ishr
123 (0x7b) lshr
124 (0x7c) iushr
125 (0x7d) lushr
126 (0x7e) iand
127 (0x7f) land
128 (0x80) ior
129 (0x81) lor
130 (0x82) ixor
131 (0x83) lxor
132 (0x84) iinc

 */

/**************************************
	Addiction
*************************************/
// integer : int1 + int2
type IADD struct {
}

func (self *IADD) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopInt()
	result := value1 + value2
	frame.OprandStack.PushInt(result)
}

// long : long1 + long2
type LADD struct {
}

func (self *LADD) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopLong()
	value2 := frame.OprandStack.PopLong()
	result := value1 + value2
	frame.OprandStack.PushLong(result)
}

// float : value1 + value2
type FADD struct {
}

func (self *FADD) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopFloat()
	value2 := frame.OprandStack.PopFloat()
	result := value1 + value2
	frame.OprandStack.PushFloat(result)
}

// double : value1 + value2
type DADD struct {
}

func (self *DADD) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopDouble()
	value2 := frame.OprandStack.PopDouble()
	result := value1 + value2
	frame.OprandStack.PushDouble(result)
}



/**************************************
	Subtract
*************************************/



// integer : value2 - value1
type ISUB struct {
}

func (self *ISUB) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopInt()
	result := value2 - value1
	frame.OprandStack.PushInt(result)
}

// long : value2 - value1
type LSUB struct {
}

func (self *LSUB) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopLong()
	value2 := frame.OprandStack.PopLong()
	result := value2 - value1
	frame.OprandStack.PushLong(result)
}

// float : value2 - value1
type FSUB struct {
}

func (self *FSUB) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopFloat()
	value2 := frame.OprandStack.PopFloat()
	result := value2 - value1
	frame.OprandStack.PushFloat(result)
}

// double : value2 - value1
type DSUB struct {
}

func (self *DSUB) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopDouble()
	value2 := frame.OprandStack.PopDouble()
	result := value2 - value1
	frame.OprandStack.PushDouble(result)
}



/**************************************
 	Multiply
 *************************************/


// integer : value2 * value1
type IMUL struct {
}

func (self *IMUL) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopInt()
	result := value2 * value1
	frame.OprandStack.PushInt(result)
}

// long : value2 * value1
type LMUL struct {
}

func (self *LMUL) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopLong()
	value2 := frame.OprandStack.PopLong()
	result := value2 * value1
	frame.OprandStack.PushLong(result)
}

// float : value2 * value1
type FMUL struct {
}

func (self *FMUL) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopFloat()
	value2 := frame.OprandStack.PopFloat()
	result := value2 * value1
	frame.OprandStack.PushFloat(result)
}

// double : value2 * value1
type DMUL struct {
}

func (self *DMUL) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopDouble()
	value2 := frame.OprandStack.PopDouble()
	result := value2 * value1
	frame.OprandStack.PushDouble(result)
}



/**************************************
	Divide
*************************************/

// integer : value2 / value1
type IDIV struct {
}

func (self *IDIV) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopInt()
	result := value2 / value1
	frame.OprandStack.PushInt(result)
}

// long : value2 / value1
type LDIV struct {
}

func (self *LDIV) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopLong()
	value2 := frame.OprandStack.PopLong()
	result := value2 / value1
	frame.OprandStack.PushLong(result)
}

// float : value2 / value1
type FDIV struct {
}

func (self *FDIV) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopFloat()
	value2 := frame.OprandStack.PopFloat()
	result := value2 / value1
	frame.OprandStack.PushFloat(result)
}

// double : value2 / value1
type DDIV struct {
}

func (self *DDIV) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopDouble()
	value2 := frame.OprandStack.PopDouble()
	result := value2 / value1
	frame.OprandStack.PushDouble(result)
}




/**************************************
	Reminder
*************************************/
// integer : value2 % value1
type IREM struct {
}

func (self *IREM) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopInt()
	result := value2 % value1
	frame.OprandStack.PushInt(result)
}

// long : value2 % value1
type LREM struct {
}

func (self *LREM) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopLong()
	value2 := frame.OprandStack.PopLong()
	result := value2 % value1
	frame.OprandStack.PushLong(result)
}

// float : value2 % value1
type FREM struct {
}

func (self *FREM) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopFloat()
	value2 := frame.OprandStack.PopFloat()
	result := float32(math.Mod(float64(value2) ,float64( value1)) )
	frame.OprandStack.PushFloat(result)
}

// double : value2 % value1
type DREM struct {
}

func (self *DREM) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopDouble()
	value2 := frame.OprandStack.PopDouble()
	result := math.Mod(float64(value2) ,float64( value1))
	frame.OprandStack.PushDouble(result)
}




/**************************************
	Negate
*************************************/
// integer : 0 - value1
type INEG struct {
}

func (self *INEG) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	result := 0 - value1
	frame.OprandStack.PushInt(result)
}

// long : 0 - value1
type LNEG struct {
}

func (self *LNEG) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopLong()
	result := 0 - value1
	frame.OprandStack.PushLong(result)
}

// float : 0 - value1
type FNEG struct {
}

func (self *FNEG) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopFloat()
	result := 0 - value1
	frame.OprandStack.PushFloat(result)
}

// double : 0 - value1
type DNEG struct {
}

func (self *DNEG) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopDouble()
	result := 0 - value1
	frame.OprandStack.PushDouble(result)
}



/**************************************
	Shift left
*************************************/
// integer : value2 << (value1 & 0b11111)
type ISHL struct {
}

func (self *ISHL) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopInt()
	result := value2 << ( uint32(value1) & uint32(0x1F))
	frame.OprandStack.PushInt(result)
}

// long : value2 << (value1 & 0b11_1111)
type LSHL struct {
}

func (self *LSHL) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopLong()
	result := value2 << ( uint32(value1) & uint32(0x3F))
	frame.OprandStack.PushLong(result)
}


/**************************************
	Shift right ,with sign extension
*************************************/
// integer : value2 >> (value1 & 0b11111)
type ISHR struct {
}

func (self *ISHR) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopInt()
	result := value2 >> ( uint32(value1) & uint32(0x1F))
	frame.OprandStack.PushInt(result)
}

// long : value2 >> (value1 & 0b11_1111)
type LSHR struct {
}

func (self *LSHR) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopLong()
	result := value2 >> ( uint32(value1) & uint32(0x3F))
	frame.OprandStack.PushLong(result)
}


/**************************************
	Shift right ,with zero extension
*************************************/
// integer : value2 >> (value1 & 0b11111) ,with zero extension
type IUSHR struct {
}

func (self *IUSHR) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopInt()
	ui := *(*uint32)(unsafe.Pointer(&value2))
	result := ui >> ( uint32(value1) & uint32(0x1F))
	frame.OprandStack.PushInt(int32(result))
}

// long : value2 >> (value1 & 0b11_1111),with zero extension
type LUSHR struct {
}

func (self *LUSHR) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopLong()
	ul := *(*uint32)(unsafe.Pointer(&value2))
	result := ul >> ( uint32(value1) & uint32(0x3F))
	frame.OprandStack.PushLong(int64(result))
}



/**************************************
	bitwise : AND , OR ,XOR
*************************************/
// integer : value1 & value2
type IAND struct {
}

func (self *IAND) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopInt()
	result := value1 & value2
	frame.OprandStack.PushInt(result)
}

// long : value1 & value2
type LAND struct {
}

func (self *LAND) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopLong()
	value2 := frame.OprandStack.PopLong()
	result := value1 & value2
	frame.OprandStack.PushLong(result)
}


// integer : value1 | value2
type IOR struct {
}

func (self *IOR) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopInt()
	result := value1 | value2
	frame.OprandStack.PushInt(result)
}

// long : value1 | value2
type LOR struct {
}

func (self *LOR) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool) {
	value1 := frame.OprandStack.PopLong()
	value2 := frame.OprandStack.PopLong()
	result := value1 | value2
	frame.OprandStack.PushLong(result)
}


// integer : value1 ^ value2
type IXOR struct {
}

func (self *IXOR) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopInt()
	value2 := frame.OprandStack.PopInt()
	result := value1 ^ value2
	frame.OprandStack.PushInt(result)
}

// long : value1 ^ value2
type LXOR struct {
}

func (self *LXOR) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	value1 := frame.OprandStack.PopLong()
	value2 := frame.OprandStack.PopLong()
	result := value1 ^ value2
	frame.OprandStack.PushLong(result)
}


/**************************************
	Increment local variable by constant
*************************************/
type IINC struct {
}

func (self *IINC) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	indexValue := index(reader,wideMode)
	constValue := index(reader,wideMode)
	value1 := frame.LocalVariables.GetInt(indexValue)
	value1 += int32(constValue)
	frame.LocalVariables.SetInt(indexValue,value1)
}
