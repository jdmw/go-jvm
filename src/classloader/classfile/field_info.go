package classfile

type FieldInfo struct {
	cf 	ClassFile
	access_flags      u2
	name_index        u2
	descriptor_index  u2
	attributes Attributes
}

func parseFieldInfo(cf 	ClassFile,r *BigEndianReader) FieldInfo{
	info := FieldInfo{cf,r.ReadU2(),r.ReadU2(),r.ReadU2(),parseAttributes(cf,r.ReadU2(),r)}
	return info
}

func (self *FieldInfo) AccessFlag() uint16 {
	return uint16(self.access_flags)
}
func (self *FieldInfo) Name() string {
	return self.cf.constant_pool.getUtf8String(self.name_index)
}

func (self *FieldInfo) Descriptor() string {
	return self.cf.constant_pool.getUtf8String(self.descriptor_index)
}

func (self *FieldInfo) Attributes() Attributes {
	return self.attributes
}
