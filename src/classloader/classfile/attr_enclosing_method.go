package classfile

import "../../util"

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * EnclosingMethod_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U2 class_index;
 *     util.U2 method_index;
 * }
 */
type EnclosingMethodAttr struct{
	cp ConstantPool
	class_index	util.U2
	method_index	util.U2
}
func (self EnclosingMethodAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	self.cp = cf.constant_pool
	self.class_index = r.ReadU2()
	self.method_index = r.ReadU2()
}

func (self *EnclosingMethodAttr) ClassIndex() util.U2{
	return self.class_index
}

func (self *EnclosingMethodAttr) MethodIndex() util.U2{
	return self.method_index
}

func (self *EnclosingMethodAttr) ClassName() string{
	return self.cp.getClassName(self.class_index)
}

func (self *EnclosingMethodAttr) MethodName() string{
	return self.cp.getUtf8String(self.method_index)
}
