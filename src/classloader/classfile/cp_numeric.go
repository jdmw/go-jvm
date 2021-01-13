package classfile

import "../../util"


import "math"

/**
CONSTANT_Integer	3	45.3	1.0.2
CONSTANT_Float	4	45.3	1.0.2
CONSTANT_Long	5	45.3	1.0.2
CONSTANT_Double	6	45.3	1.0.2
 */

/**
CONSTANT_Integer_info {
    util.U1 tag;
    util.U4 bytes;
}

CONSTANT_Float_info {
    util.U1 tag;
    util.U4 bytes;
}
*/

type ConstIntegerInfo struct {
	value util.U4
}
type ConstFloatInfo struct {
	value float32
}

func (self *ConstIntegerInfo) readInfo(r *util.BigEndianReader) {
	self.value = r.ReadU4()
}

func (self *ConstIntegerInfo) Value() util.U4 {
	return self.value
}

func (self *ConstFloatInfo) readInfo(r *util.BigEndianReader) {
	self.value = math.Float32frombits(uint32(r.ReadU4()))
}

func (self *ConstFloatInfo) Value() float32 {
	return self.value
}
/*
CONSTANT_Long_info {
    util.U1 tag;
    util.U4 high_bytes;
    util.U4 low_bytes;
}

CONSTANT_Double_info {
    util.U1 tag;
    util.U4 high_bytes;
    util.U4 low_bytes;
}
*/
type ConstLongInfo struct {
	value util.U8
}
type ConstDoubleInfo struct {
	value float64
}

func (self *ConstLongInfo) readInfo(r *util.BigEndianReader) {
	self.value = r.ReadU8()
}

func (self *ConstDoubleInfo) readInfo(r *util.BigEndianReader) {
	self.value = math.Float64frombits(uint64(r.ReadU8()))
}

func (self *ConstLongInfo) Value() util.U8 {
	return self.value
}

func (self *ConstDoubleInfo) Value() float64 {
	return self.value
}

