package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * MethodParameters_attribute {
 *     u2 attribute_name_index;
 *     u4 attribute_length;
 *     u1 parameters_count;
 *     {   u2 name_index;
 *         u2 access_flags;
 *     } parameters[parameters_count];
 * }
 */
type MethodParametersAttr []MethodParametersInfo
type MethodParametersInfo struct{
	cp ConstantPool
	name_index	u2
	access_flags	u2
}
func (self *MethodParametersAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	infos := make([]MethodParametersInfo,r.ReadU1())
	for i := range infos {
		infos[i] = MethodParametersInfo{cf.constant_pool,r.ReadU2(),r.ReadU2()}
	}
}

func (self *MethodParametersInfo) Parameter() (u2,string){
	name := ""
	if self.name_index > 0 {
		name = self.cp.getUtf8String(self.name_index)
	}
	return self.access_flags,name
}

func (self MethodParametersInfo) AccessFlags() u2{
	return self.access_flags
}
func (self MethodParametersInfo) NameIndex() u2{
	return self.name_index
}
