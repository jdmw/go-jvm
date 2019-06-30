package classfile

type FieldInfo struct {
	cf 	ClassFile
	access_flags      u2
	name_index        u2
	descriptor_index  u2
	attributes Attributes
}

func parseFieldInfo(cf 	ClassFile,r *BigEndianReader) FieldInfo{
	info := FieldInfo{cf,r.ReadU2(),r.ReadU2(),r.ReadU2(),parseAttributes(cf,r)}
	return info
}

func (self *FieldInfo) Name() string {
	return self.cf.constant_pool.getUtf8String(self.name_index)
}

