package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * Exceptions_attribute {
 *     u2 attribute_name_index;
 *     u4 attribute_length;
 *     u2 number_of_exceptions;
 *     u2 exception_index_table[number_of_exceptions];
 * }
 */
	type ExceptionsAttr struct{
		cp ConstantPool
	}
	func (self *ExceptionsAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
		self.cp = cf.constant_pool
	}

