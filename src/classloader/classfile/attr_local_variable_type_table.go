package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * LocalVariableTypeTable_attribute {
 *     u2 attribute_name_index;
 *     u4 attribute_length;
 *     u2 local_variable_type_table_length;
 *     {   u2 start_pc;
 *         u2 length;
 *         u2 name_index;
 *         u2 signature_index;
 *         u2 index;
 *     } local_variable_type_table[local_variable_type_table_length];
 * }
 */
type LocalVariableTypeTableAttr []LocalVariableTypeTableInfo

type LocalVariableTypeTableInfo struct{
	cp ConstantPool
	start_pc u2
	length	u2
	name_index	u2
	signature_index	u2
	index	u2
}

func (self *LocalVariableTypeTableAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	*self = make([]LocalVariableTypeTableInfo,r.ReadU2())
	for i := range *self{
		(*self)[i] = LocalVariableTypeTableInfo{cf.constant_pool,r.ReadU2(),r.ReadU2(),r.ReadU2(),r.ReadU2(),r.ReadU2()}
	}
}

func (self LocalVariableTypeTableInfo) CodeIndices() (u2,u2){
	return  self.start_pc, self.start_pc + self.length
}

func (self LocalVariableTypeTableInfo) Name() string{
	return self.cp.getUtf8String(self.name_index)
}


func (self LocalVariableTypeTableInfo) Signature() string{
	return self.cp.getUtf8String(self.signature_index)
}


func (self LocalVariableTypeTableInfo) Index() u2{
	return self.index
}
