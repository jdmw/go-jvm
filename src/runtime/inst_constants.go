package runtime

import "../util"

/**
 *
 *   push constant value into oprand stack
		Constants
		00 (0x00) nop
		01 (0x01) aconst_null
		02 (0x02) iconst_m1
		03 (0x03) iconst_0
		04 (0x04) iconst_1
		05 (0x05) iconst_2
		06 (0x06) iconst_3
		07 (0x07) iconst_4
		08 (0x08) iconst_5
		09 (0x09) lconst_0
		10 (0x0a) lconst_1
		11 (0x0b) fconst_0
		12 (0x0c) fconst_1
		13 (0x0d) fconst_2
		14 (0x0e) dconst_0
		15 (0x0f) dconst_1
		16 (0x10) bipush
		17 (0x11) sipush
		18 (0x12) ldc
		19 (0x13) ldc_w
		20 (0x14) ldc2_w

 */


type ACONST_NULL struct {
}

func (self *ACONST_NULL) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	frame.OprandStack.PushRef(nil )
}

type ICONST_N struct {
	num int32
}

func (self *ICONST_N) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	frame.OprandStack.PushInt(self.num )
}

type FCONST_N struct {
	num float32
}

func (self *FCONST_N) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	frame.OprandStack.PushFloat(self.num )
}

type DCONST_N struct {
	num float64
}

func (self *DCONST_N) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	frame.OprandStack.PushDouble(self.num )
}

type LCONST_N struct {
	num util.U8
}

func (self *LCONST_N) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	frame.OprandStack.PushU8(self.num )
}


/**
 * Push byte
 */
type BIPUSH struct {
}

func (self *BIPUSH) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	smallInt := int8(reader.ReadU1()) // sign-extended to int byte
	frame.OprandStack.PushInt(int32(smallInt))
}

/**
 * Push signed-short
 */
type SIPUSH struct {
}

func (self *SIPUSH) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	num :=  int16 (reader.ReadU2()) // sign-extended to int byte
	frame.OprandStack.PushU4(util.U4(num))
}



type LDC struct {
}

func (self *LDC) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	index := reader.ReadU1() // sign-extended to int byte
	frame.class.Constantpool.PushU4Num(util.U2(index),frame.OprandStack)
}

type LDC_W struct {
}

func (self *LDC_W) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	index := reader.ReadU2() // sign-extended to int byte
	frame.class.Constantpool.PushU4Num(index,frame.OprandStack)
}


type LDC2_W struct {
}

func (self *LDC2_W) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)  {
	index := reader.ReadU2() // sign-extended to int byte
	frame.class.Constantpool.PushU8Num(index,frame.OprandStack)
}







