package classfile

/**
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
 */

type SourceFileAttr struct {
	sourcefile string
}

func (self *SourceFileAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	self.sourcefile = cf.constant_pool.getUtf8String(r.ReadU2())
}

func (self *SourceFileAttr) SourceFile() string {
	return self.sourcefile
}
