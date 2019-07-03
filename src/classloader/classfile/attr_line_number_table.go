package classfile

/**

LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}


*/
type LineNumberTableAttrInfo struct {
	start_pc          u2
	line_number            u2
}
type LineNumberTableAttr []LineNumberTableAttrInfo

func (self *LineNumberTableAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	attr := make([]LineNumberTableAttrInfo,r.ReadU2())
	for i := range attr{
		attr[i] = LineNumberTableAttrInfo{r.ReadU2(),r.ReadU2()}
	}
}
