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
	constantvalue_index u2
}

func (self *ConstValueAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.constantvalue_index = r.ReadU2()
}

func (self *ConstValueAttr) ConstantValueIndex() u2{
	return self.constantvalue_index
}



