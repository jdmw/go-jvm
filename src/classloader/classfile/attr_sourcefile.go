package classfile

import "../../util"


/**
SourceFile_attribute {
    util.U2 attribute_name_index;
    util.U4 attribute_length;
    util.U2 sourcefile_index;
}
 */

type SourceFileAttr struct {
	sourcefile string
}

func (self *SourceFileAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	self.sourcefile = cf.constant_pool.getUtf8String(r.ReadU2())
}

func (self *SourceFileAttr) SourceFile() string {
	return self.sourcefile
}
