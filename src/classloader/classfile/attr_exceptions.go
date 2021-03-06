package classfile


import "../../util"
/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * Exceptions_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U2 number_of_exceptions;
 *     util.U2 exception_index_table[number_of_exceptions];
 * }
 */

type ExceptionsAttr struct{
	cp ConstantPool
	exception_index_table []util.U2
}

func (self *ExceptionsAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	self.cp = cf.constant_pool
	self.exception_index_table = r.ReadU2s()
}

func (self *ExceptionsAttr) GetExceptions() []string{
	excpetions := make([]string,len(self.exception_index_table))
	for i,index := range self.exception_index_table {
		excpetions[i] = self.cp.getUtf8String(index)
	}
	return excpetions
}

