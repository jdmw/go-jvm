package classfile

import "fmt"

type ConstantPool []ConstantPoolInfo

func parseConstPool(length int,r *BigEndianReader) (ConstantPool,int) {
	var info ConstantPoolInfo
	cp := make(ConstantPool,length)
	accFlagToBeChecked := 0
	flag := 0
	num_of_entries := 1
	for i:=0;i<length;i+=num_of_entries {
		info,flag,num_of_entries = parseConstantPoolInfo(cp, r)
		cp[i] = info
		fmt.Printf("constant_pool[%v] = %v\n",i,info)
		accFlagToBeChecked |= flag
	}
	return cp,accFlagToBeChecked
}

func (self ConstantPool) getConstantPoolInfo(index u2) ConstantPoolInfo {
	if index == 0 || int(index) >= len(self){
		panic("constant pool index error " + string(index))
	}
	return self[index -1]
}

func (self ConstantPool) getUtf8String(index u2 ) string {
	return self.getConstantPoolInfo(index).(*ConstUtf8Info).str
}

func (self ConstantPool) getNameAndType(index u2) (string,string) {
	info := self.getConstantPoolInfo(index).(*ConstNameAndTypeInfo)
	name := self.getUtf8String(info.name_index)
	description := self.getUtf8String(info.descriptor_index)
	return name,description
}

func (self ConstantPool) getClassName(index u2) string {
	if index == 0 {
		return "java.lang.Object"
	}

	classInfo := self.getConstantPoolInfo(index).(*ConstClassInfo)
	return self.getUtf8String(classInfo.name_index)
}

