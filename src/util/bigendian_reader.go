package util

import (
	"encoding/binary"
)

type U1 byte
type U2 uint16
type U4 uint32
type U8 uint64

type BigEndianReader struct{
	offset U4
	data []byte
}

func NewBigEndianReader(data []byte) *BigEndianReader{
	return &BigEndianReader{0,data }
}

func (self *BigEndianReader) Offset() U4 {
	return self.offset
}

func (self *BigEndianReader) ReadU1() U1{
	data := self.data[self.offset]
	self.offset++
	return U1(data)
}


func (self *BigEndianReader) ReadU2() U2{
	data := binary.BigEndian.Uint16(self.data[self.offset:])
	self.offset += 2
	return U2(data)
}


func (self *BigEndianReader) ReadU4() U4{
	data := binary.BigEndian.Uint32(self.data[self.offset:])
	self.offset += 4
	return U4(data)
}


func (self *BigEndianReader) ReadU8() U8{
	data := binary.BigEndian.Uint64(self.data[self.offset:])
	self.offset += 8
	return U8(data)
}

func (self *BigEndianReader) ReadU2s() []U2{
	count := self.ReadU2()
	arr := make([]U2,count)
	for i := range arr {
		arr[i] = self.ReadU2()
	}
	return arr
}

func (self *BigEndianReader) ReadU2Array(count U4) []U2{
	arr := make([]U2,count)
	for i := range arr {
		arr[i] = self.ReadU2()
	}
	return arr
}

func (self *BigEndianReader) ReadByteArray(count U4) []byte{
	data := self.data[self.offset:self.offset+count]
	self.offset += count
	return data
}


func (self *BigEndianReader) MaxLength() U4{
	return U4(len(self.data))
}

func (self *BigEndianReader) HasNext() bool{
	return self.offset + 1 < U4(len(self.data))
}