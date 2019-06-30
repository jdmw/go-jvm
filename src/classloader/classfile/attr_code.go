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
	self.maxStack = r.ReadU2()
	self.maxLocals = r.ReadU2()

	code_length := r.ReadU4()
	self.code = r.ReadByteArray(code_length)

	self.exceptionTable = make([]ExceptionTableInfo,r.ReadU2())
	for i := range self.exceptionTable {
		self.exceptionTable[i] = ExceptionTableInfo{
			r.ReadU2(),
			r.ReadU2(),
			r.ReadU2(),
			r.ReadU2(),
		}
	}

	self.attributes = parseAttributes(cf ,r.ReadU2(),r)
}

