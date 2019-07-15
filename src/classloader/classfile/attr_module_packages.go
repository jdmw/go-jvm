package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * ModulePackages_attribute {
 *     u2 attribute_name_index;
 *     u4 attribute_length;
 *     u2 package_count;
 *     u2 package_index[package_count];
 * }
 */
type ModulePackagesAttr struct{
	cp ConstantPool
	package_index []u2
}
func (self ModulePackagesAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.package_index = r.ReadU2s()
}


func (self *ModulePackagesAttr) PackageIndex() []u2{
	return self.package_index
}
