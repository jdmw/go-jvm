package classfile

/*
ConstantValue_attribute {
u2 attribute_name_index;
u4 attribute_length;
u2 constantvalue_index;
}

*/
type ConstantValueAttr struct {
	cp ConstantPool
	constantvalue_index u2
}

func (self *ConstantValueAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.constantvalue_index = r.ReadU2()
}

func (self *ConstantValueAttr) GetConstantPoolInfo() ConstantPoolInfo{
	return self.cp.getConstantPoolInfo(self.constantvalue_index)
}



