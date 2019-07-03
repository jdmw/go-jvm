package classfile

type MethodInfo struct {
	cf 	ClassFile
	access_flags      u2
	name_index        u2
	descriptor_index  u2
	attributes Attributes
}

func parseMethodInfo(cf 	ClassFile,r *BigEndianReader) MethodInfo{
	info := MethodInfo{cf,r.ReadU2(),r.ReadU2(),r.ReadU2(),parseAttributes(cf,r.ReadU2(),r)}
	return info
}

func (self *MethodInfo) Name() string {
	return self.cf.constant_pool.getUtf8String(self.name_index)
}

