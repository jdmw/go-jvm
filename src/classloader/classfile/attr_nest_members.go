package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * NestMembers_attribute {
 *     u2 attribute_name_index;
 *     u4 attribute_length;
 *     u2 number_of_classes;
 *     u2 classes[number_of_classes];
 * }
 */
type NestMembersAttr struct{
	cp ConstantPool
	classes []u2
}
func (self NestMembersAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.classes = r.ReadU2s()
}

func (self *NestMembersAttr) ClassesIndices() []u2{
	return self.classes
}
