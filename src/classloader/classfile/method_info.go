package classfile

type MethodInfo struct {
	cp 	ConstantPool
	access_flags      u2
	name_index        u2
	descriptor_index  u2
	attributes Attributes
}

func parseMethodInfo(cf 	ClassFile,r *BigEndianReader) MethodInfo{
	info := MethodInfo{cf.constant_pool,r.ReadU2(),r.ReadU2(),r.ReadU2(),parseAttributes(cf,r.ReadU2(),r)}
	return info
}

func (self *MethodInfo) Name() string {
	return self.cp.getUtf8String(self.name_index)
}

func (self *MethodInfo) Descriptor() string {
	return self.cp.getUtf8String(self.descriptor_index)
}

func (self *MethodInfo) Attributes() Attributes {
	return self.attributes
}
