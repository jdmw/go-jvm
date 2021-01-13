package classfile


import "../../util"

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * MethodParameters_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U1 parameters_count;
 *     {   util.U2 name_index;
 *         util.U2 access_flags;
 *     } parameters[parameters_count];
 * }
 */
type MethodParametersAttr []MethodParametersInfo
type MethodParametersInfo struct{
	cp ConstantPool
	name_index	util.U2
	access_flags	util.U2
}
func (self *MethodParametersAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	infos := make([]MethodParametersInfo,r.ReadU1())
	for i := range infos {
		infos[i] = MethodParametersInfo{cf.constant_pool,r.ReadU2(),r.ReadU2()}
	}
}

func (self *MethodParametersInfo) Parameter() (util.U2,string){
	name := ""
	if self.name_index > 0 {
		name = self.cp.getUtf8String(self.name_index)
	}
	return self.access_flags,name
}

func (self MethodParametersInfo) AccessFlags() util.U2{
	return self.access_flags
}
func (self MethodParametersInfo) NameIndex() util.U2{
	return self.name_index
}
