package classfile

type ConstantPool []ConstantPoolInfo

func parseConstPool(length int,r *BigEndianReader) (ConstantPool,int) {
	var info ConstantPoolInfo
	cp := make(ConstantPool,length)
	accFlagToBeChecked := 0
	flag := 0
	for i:=0;i<length;i++ {
		info,flag = parseConstantPoolInfo(cp, r)
		cp[i] = info
		accFlagToBeChecked |= flag
	}
	return cp,accFlagToBeChecked
}

func (self ConstantPool) getUtf8String(index u2 ) string {
	return self[index].(*ConstUtf8Info).str
}

func (self ConstantPool) getNameAndType(index u2) (string,string) {
	info := self[index].(*ConstNameAndTypeInfo)
	name := self.getUtf8String(info.name_index)
	description := self.getUtf8String(info.descriptor_index)
	return name,description
}

func (self ConstantPool) getClassName(index u2) string {
	classInfo := self[index].(*ConstClassInfo)
	return self.getUtf8String(classInfo.name_index)
}

