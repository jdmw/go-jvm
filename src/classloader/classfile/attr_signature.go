package classfile

import "../../util"


/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * Signature_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U2 signature_index;
 * }
 */

type SignatureAttr struct{
	cp ConstantPool
	signature_index	util.U2
}
func (self SignatureAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	self.cp = cf.constant_pool
	self.signature_index = r.ReadU2()
}

func (self *SignatureAttr) Signature() string{
	return self.cp.getUtf8String(self.signature_index)
}
