package classfile

import "../../util"

/**
Code_attribute {
    util.U2 attribute_name_index;
    util.U4 attribute_length;
    util.U2 max_stack;
    util.U2 max_locals;
    util.U4 code_length;
    util.U1 code[code_length];
    util.U2 exception_table_length;
    {   util.U2 start_pc;
        util.U2 end_pc;
        util.U2 handler_pc;
        util.U2 catch_type;
    } exception_table[exception_table_length];
    util.U2 attributes_count;
    attribute_info attributes[attributes_count];
}
 */

type CodeAttr struct {
	cf ClassFile
	maxStack util.U2
	maxLocals util.U2
	//util.U4 code_length util.U4
	code []byte
	exceptionTable []ExceptionTableInfo
	attributes Attributes
}

func (self *CodeAttr) Attributes() Attributes {
	return self.attributes
}

func (self *CodeAttr) ExceptionTable() []ExceptionTableInfo {
	return self.exceptionTable
}

func (self *CodeAttr) Code() []byte {
	return self.code
}

func (self *CodeAttr) MaxLocals() util.U2 {
	return self.maxLocals
}

func (self *CodeAttr) MaxStack() util.U2 {
	return self.maxStack
}



type ExceptionTableInfo struct {
	start_pc    util.U2
	end_pc      util.U2
	handler_pc  util.U2
	catch_type  util.U2
}

func (self *CodeAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
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

