package classfile

import "../../util"


type MethodInfo struct {
	cp 	ConstantPool
	access_flags      util.U2
	name_index        util.U2
	descriptor_index  util.U2
	attributes Attributes
}

func parseMethodInfo(cf 	ClassFile,r *util.BigEndianReader) MethodInfo{
	info := MethodInfo{cf.constant_pool,r.ReadU2(),r.ReadU2(),r.ReadU2(),parseAttributes(cf,r.ReadU2(),r)}
	return info
}


func (self *MethodInfo) AccessFlag() uint16 {
	return uint16(self.access_flags)
}

func (self *MethodInfo) Name() string {
	return self.cp.getUtf8String(self.name_index)
}

func (self *MethodInfo) Descriptor() string {
	return self.cp.getUtf8String(self.descriptor_index)
}

func (self *MethodInfo) Attributes() *Attributes {
	return &self.attributes
}
