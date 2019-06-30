package classfile

/*
ConstantValue_attribute {
u2 attribute_name_index;
u4 attribute_length;
u2 constantvalue_index;
}

*/
type ConstValueAttr struct {
	cp ConstantPool
	attribute_name_index u2
	constantvalue_index u2
}

func (self *ConstValueAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.attribute_name_index = r.ReadU2()
	r.ReadU4()
	self.constantvalue_index = r.ReadU2()
}

func (self *ConstValueAttr) AttributeName() string{
	return self.cp.getUtf8String(self.attribute_name_index)
}

func (self *ConstValueAttr) ConstantValueIndex() u2{
	return self.constantvalue_index
}



