package classfile

import "../../util"

/**

LineNumberTable_attribute {
    util.U2 attribute_name_index;
    util.U4 attribute_length;
    util.U2 line_number_table_length;
    {   util.U2 start_pc;
        util.U2 line_number;
    } line_number_table[line_number_table_length];
}

*/

type LineNumberTableAttrInfo struct {
	start_pc          util.U2
	line_number            util.U2
}
type LineNumberTableAttr []LineNumberTableAttrInfo

func (self *LineNumberTableAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	attr := make([]LineNumberTableAttrInfo,r.ReadU2())
	for i := range attr{
		attr[i] = LineNumberTableAttrInfo{r.ReadU2(),r.ReadU2()}
	}
	*self = attr
}
