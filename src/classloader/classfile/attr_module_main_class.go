package classfile

import "../../util"

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * ModuleMainClass_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U2 main_class_index;
 * }
 */
type ModuleMainClassAttr struct{
	cp ConstantPool
	main_class_index	util.U2
}
func (self *ModuleMainClassAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	self.cp = cf.constant_pool
	self.main_class_index = r.ReadU2()
}

func (self ModuleMainClassAttr) MainClassName() string{
	return self.cp.getClassName(self.main_class_index)
}
