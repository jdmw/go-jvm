package classfile

import "../../util"


/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * SourceDebugExtension_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U1 debug_extension[attribute_length];
 * }
 */
type SourceDebugExtensionAttr struct{
	debug_extension	[]byte
}
func (self *SourceDebugExtensionAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	self.debug_extension = r.ReadByteArray(length)
}

func (self *SourceDebugExtensionAttr) DebugExtension()  []byte{
	return self.debug_extension
}
