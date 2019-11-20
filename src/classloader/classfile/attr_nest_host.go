package classfile

import "../../util"

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * NestHost_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U2 host_class_index;
 * }
 */
type NestHostAttr struct{
	cp ConstantPool
	host_class_index	util.U2
}
func (self NestHostAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	self.cp = cf.constant_pool
	self.host_class_index = r.ReadU2()
}




func (self *NestHostAttr) HostClassIndex() util.U2{
	return self.host_class_index
}
