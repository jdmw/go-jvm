package classfile

/**
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
 */

type CodeAttr struct {
	cf ClassFile
	maxStack u2
	maxLocals u2
	//u4 code_length u4
	code []byte
	exceptionTable []ExceptionTableInfo
	attributes Attributes
}

type ExceptionTableInfo struct {
	start_pc    u2
	end_pc      u2
	handler_pc  u2
	catch_type  u2
}

func (self *CodeAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	self.cf = cf
	attribute_name_index := r.ReadU2()
	attribute_length := r.ReadU4()
	self.maxStack = r.ReadU2()
	self.maxLocals = r.ReadU2()

	code_length := r.ReadU4()
	self.code = r.ReadByteArray(code_length)

	exception_table_length := r.ReadU2()


	self.attribute_name_index = r.ReadU2()
	r.ReadU4()
	self.constantvalue_index = r.ReadU2()
}

func (self *CodeAttr) AttributeName() string{
	return self.cp.getUtf8String(self.attribute_name_index)
}

