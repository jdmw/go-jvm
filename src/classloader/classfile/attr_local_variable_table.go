package classfile


import "../../util"

/**

LocalVariableTable_attribute {
    util.U2 attribute_name_index;
    util.U4 attribute_length;
    util.U2 local_variable_table_length;
    {   util.U2 start_pc;
        util.U2 length;
        util.U2 name_index;
        util.U2 descriptor_index;
        util.U2 index;
    } local_variable_table[local_variable_table_length];
}
 */

type LocalVariableTableAttrInfo struct {
	start_pc          util.U2
	length            util.U2
	name_index        util.U2
	descriptor_index  util.U2
	index             util.U2
}
type LocalVariableTableAttr []LocalVariableTableAttrInfo

func (self *LocalVariableTableAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	*self = make([]LocalVariableTableAttrInfo,r.ReadU2())
	for i := range *self{
		(*self)[i] = LocalVariableTableAttrInfo{r.ReadU2(),r.ReadU2(),r.ReadU2(),r.ReadU2(),r.ReadU2(),}
	}
	//self.sourcefile = cf.constant_pool.getUtf8String(r.ReadU2())
}

