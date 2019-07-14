package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * EnclosingMethod_attribute {
 *     u2 attribute_name_index;
 *     u4 attribute_length;
 *     u2 class_index;
 *     u2 method_index;
 * }
 */
type EnclosingMethodAttr struct{
	cp ConstantPool
	class_index	u2
	method_index	u2
}
func (self EnclosingMethodAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.class_index = r.ReadU2()
	self.method_index = r.ReadU2()
}

func (self *EnclosingMethodAttr) ClassIndex() u2{
	return self.class_index
}

func (self *EnclosingMethodAttr) MethodIndex() u2{
	return self.method_index
}

func (self *EnclosingMethodAttr) ClassName() string{
	return self.cp.getClassName(self.class_index)
}

func (self *EnclosingMethodAttr) MethodName() string{
	return self.cp.getUtf8String(self.method_index)
}
