package classfile

import "../../util"

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * ModulePackages_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U2 package_count;
 *     util.U2 package_index[package_count];
 * }
 */
type ModulePackagesAttr struct{
	cp ConstantPool
	package_index []util.U2
}
func (self ModulePackagesAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	self.cp = cf.constant_pool
	self.package_index = r.ReadU2s()
}


func (self *ModulePackagesAttr) PackageIndex() []util.U2{
	return self.package_index
}
