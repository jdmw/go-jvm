package classfile

import "../../util"

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * LocalVariableTypeTable_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U2 local_variable_type_table_length;
 *     {   util.U2 start_pc;
 *         util.U2 length;
 *         util.U2 name_index;
 *         util.U2 signature_index;
 *         util.U2 index;
 *     } local_variable_type_table[local_variable_type_table_length];
 * }
 */
type LocalVariableTypeTableAttr []LocalVariableTypeTableInfo

type LocalVariableTypeTableInfo struct{
	cp ConstantPool
	start_pc util.U2
	length	util.U2
	name_index	util.U2
	signature_index	util.U2
	index	util.U2
}

func (self *LocalVariableTypeTableAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	*self = make([]LocalVariableTypeTableInfo,r.ReadU2())
	for i := range *self{
		(*self)[i] = LocalVariableTypeTableInfo{cf.constant_pool,r.ReadU2(),r.ReadU2(),r.ReadU2(),r.ReadU2(),r.ReadU2()}
	}
}

func (self LocalVariableTypeTableInfo) CodeIndices() (util.U2,util.U2){
	return  self.start_pc, self.start_pc + self.length
}

func (self LocalVariableTypeTableInfo) Name() string{
	return self.cp.getUtf8String(self.name_index)
}


func (self LocalVariableTypeTableInfo) Signature() string{
	return self.cp.getUtf8String(self.signature_index)
}


func (self LocalVariableTypeTableInfo) Index() util.U2{
	return self.index
}
