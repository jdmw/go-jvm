package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * InnerClasses_attribute {
 *     u2 attribute_name_index;
 *     u4 attribute_length;
 *     u2 number_of_classes;
 *     {   u2 inner_class_info_index;
 *         u2 outer_class_info_index;
 *         u2 inner_name_index;
 *         u2 inner_class_access_flags;
 *     } classes[number_of_classes];
 * }
 */

type InnerClassesAttr []InnerClassesInfo

type InnerClassesInfo struct{
	cp ConstantPool
	inner_class_info_index u2
	outer_class_info_index u2
	inner_name_index u2
	inner_class_access_flags u2
}
func (self *InnerClassesAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	attr := make([]InnerClassesInfo,r.ReadU2())
	for _,info := range attr {
		info.cp = cf.constant_pool
		info.inner_class_info_index = r.ReadU2()
		info.outer_class_info_index = r.ReadU2()
		info.inner_name_index = r.ReadU2()
		info.inner_class_access_flags = r.ReadU2()
	}
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

func (self InnerClassesInfo) AccessFlags() u2{
	return self.inner_class_access_flags
}
