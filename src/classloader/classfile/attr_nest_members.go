package classfile

import "../../util"

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * NestMembers_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U2 number_of_classes;
 *     util.U2 classes[number_of_classes];
 * }
 */
type NestMembersAttr struct{
	cp ConstantPool
	classes []util.U2
}
func (self NestMembersAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	self.cp = cf.constant_pool
	self.classes = r.ReadU2s()
}

func (self *NestMembersAttr) ClassesIndices() []util.U2{
	return self.classes
}
