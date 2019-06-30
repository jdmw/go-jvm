package classfile

import (
	"encoding/binary"
)

type u1 byte
type u2 uint16
type u4 uint32
type u8 uint64

type BigEndianReader struct{
	data []byte
}

func (self *BigEndianReader) ReadU1() u1{
	data := self.data[0]
	self.data = self.data[1:]
	return u1(data)
}


func (self *BigEndianReader) ReadU2() u2{
	data := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return u2(data)
}


func (self *BigEndianReader) ReadU4() u4{
	data := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return u4(data)
}


func (self *BigEndianReader) ReadU8() u8{
	data := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return u8(data)
}


func (self *BigEndianReader) ReadU2s() []u2{
	count := self.ReadU2()
	arr := make([]u2,count)
	for i := range arr {
		arr[i] = self.ReadU2()
	}
	return arr
}

func (self *BigEndianReader) ReadU2Array(count u4) []u2{
	arr := make([]u2,count)
	for i := range arr {
		arr[i] = self.ReadU2()
	}
	return arr
}

func (self *BigEndianReader) ReadByteArray(count u4) []byte{
	data := self.data[:count]
	self.data = self.data[count:]
	return data
}


