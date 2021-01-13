package classfile

import "../../util"

/*
ConstantValue_attribute {
util.U2 attribute_name_index;
util.U4 attribute_length;
util.U2 constantvalue_index;
}

*/
type ConstantValueAttr struct {
	cp ConstantPool
	constantvalue_index util.U2
}

func (self *ConstantValueAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	self.cp = cf.constant_pool
	self.constantvalue_index = r.ReadU2()
}

func (self *ConstantValueAttr) GetConstantPoolInfo() ConstantPoolInfo{
	return self.cp.getConstantPoolInfo(self.constantvalue_index)
}



