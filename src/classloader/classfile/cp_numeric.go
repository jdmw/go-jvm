package classfile

import "math"

/**
CONSTANT_Integer	3	45.3	1.0.2
CONSTANT_Float	4	45.3	1.0.2
CONSTANT_Long	5	45.3	1.0.2
CONSTANT_Double	6	45.3	1.0.2
 */

/**
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}

CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/

type ConstIntegerInfo struct {
	value u4
}
type ConstFloatInfo struct {
	value float32
}

func (self *ConstIntegerInfo) readInfo(r *BigEndianReader) {
	self.value = r.ReadU4()
}

func (self *ConstIntegerInfo) Value() u4 {
	return self.value
}

func (self *ConstFloatInfo) readInfo(r *BigEndianReader) {
	self.value = math.Float32frombits(uint32(r.ReadU4()))
}

func (self *ConstFloatInfo) Value() float32 {
	return self.value
}
/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}

CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstLongInfo struct {
	value u8
}
type ConstDoubleInfo struct {
	value float64
}

func (self *ConstLongInfo) readInfo(r *BigEndianReader) {
	self.value = r.ReadU8()
}

func (self *ConstDoubleInfo) readInfo(r *BigEndianReader) {
	self.value = math.Float64frombits(uint64(r.ReadU8()))
}

func (self *ConstLongInfo) Value() u8 {
	return self.value
}

func (self *ConstDoubleInfo) Value() float64 {
	return self.value
}

