package classfile


import "../../util"

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * InnerClasses_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U2 number_of_classes;
 *     {   util.U2 inner_class_info_index;
 *         util.U2 outer_class_info_index;
 *         util.U2 inner_name_index;
 *         util.U2 inner_class_access_flags;
 *     } classes[number_of_classes];
 * }
 */

type InnerClassesAttr []InnerClassesInfo

type InnerClassesInfo struct{
	cp ConstantPool
	inner_class_info_index util.U2
	outer_class_info_index util.U2
	inner_name_index util.U2
	inner_class_access_flags util.U2
}
func (self *InnerClassesAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	attr := make([]InnerClassesInfo,r.ReadU2())
	for i,info := range attr {
		info.cp = cf.constant_pool
		info.inner_class_info_index = r.ReadU2()
		info.outer_class_info_index = r.ReadU2()
		info.inner_name_index = r.ReadU2()
		info.inner_class_access_flags = r.ReadU2()
		attr[i] = info // TODO
	}
	*self = attr
}

/**
 * 获取内部类名
 * 匿名类返回空字符串
 */
func (self InnerClassesInfo) InnerClassName() string{
	if (self.inner_class_info_index > 0) {
		return self.cp.getClassName(self.inner_class_info_index)
	}
	return ""
}

/**
 * 获取外部类类名
 * 局部类、匿名类无外部类，返回空字符串
 */
func (self InnerClassesInfo) OutterClassName() string{
	if (self.inner_class_info_index > 0) {
		return self.cp.getClassName(self.outer_class_info_index)
	}
	return ""
}

func (self InnerClassesInfo) InnerName() string{
	return self.cp.getUtf8String(self.inner_name_index)
}

func (self InnerClassesInfo) AccessFlags() util.U2{
	return self.inner_class_access_flags
}
