package classfile

/**

LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
 */

type LocalVariableTableAttrInfo struct {
	start_pc          u2
	length            u2
	name_index        u2
	descriptor_index  u2
	index             u2
}
type LocalVariableTableAttr []LocalVariableTableAttrInfo

func (self *LocalVariableTableAttrInfo) parse(cf ClassFile,length u4,r *BigEndianReader) {
	attr := make([]LocalVariableTableAttrInfo,r.ReadU2())
	for i := range attr{
		attr[i] = LocalVariableTableAttrInfo{r.ReadU2(),r.ReadU2(),r.ReadU2(),r.ReadU2(),r.ReadU2(),}
	}
	//self.sourcefile = cf.constant_pool.getUtf8String(r.ReadU2())
}

