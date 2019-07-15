package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * ModuleMainClass_attribute {
 *     u2 attribute_name_index;
 *     u4 attribute_length;
 *     u2 main_class_index;
 * }
 */
type ModuleMainClassAttr struct{
	cp ConstantPool
	main_class_index	u2
}
func (self *ModuleMainClassAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.main_class_index = r.ReadU2()
}

func (self ModuleMainClassAttr) MainClassName() string{
	return self.cp.getClassName(self.main_class_index)
}
