package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * NestHost_attribute {
 *     u2 attribute_name_index;
 *     u4 attribute_length;
 *     u2 host_class_index;
 * }
 */
type NestHostAttr struct{
	cp ConstantPool
	host_class_index	u2
}
func (self NestHostAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.host_class_index = r.ReadU2()
}




func (self *NestHostAttr) HostClassIndex() u2{
	return self.host_class_index
}
