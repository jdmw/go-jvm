package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * SourceDebugExtension_attribute {
 *     u2 attribute_name_index;
 *     u4 attribute_length;
 *     u1 debug_extension[attribute_length];
 * }
 */
type SourceDebugExtensionAttr struct{
	debug_extension	[]byte
}
func (self *SourceDebugExtensionAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	self.debug_extension = r.ReadByteArray(length)
}

func (self *SourceDebugExtensionAttr) DebugExtension()  []byte{
	return self.debug_extension
}
