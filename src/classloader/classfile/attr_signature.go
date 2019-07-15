package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * Signature_attribute {
 *     u2 attribute_name_index;
 *     u4 attribute_length;
 *     u2 signature_index;
 * }
 */

type SignatureAttr struct{
	cp ConstantPool
	signature_index	u2
}
func (self SignatureAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.signature_index = r.ReadU2()
}

func (self *SignatureAttr) Signature() string{
	return self.cp.getUtf8String(self.signature_index)
}
