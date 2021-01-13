package classfile

import "../../util"

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * Module_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 * 
 *     util.U2 module_name_index;
 *     util.U2 module_flags;
 *     util.U2 module_version_index;
 * 
 *     util.U2 requires_count;
 *     {   util.U2 requires_index;
 *         util.U2 requires_flags;
 *         util.U2 requires_version_index;
 *     } requires[requires_count];
 * 
 *     util.U2 exports_count;
 *     {   util.U2 exports_index;
 *         util.U2 exports_flags;
 *         util.U2 exports_to_count;
 *         util.U2 exports_to_index[exports_to_count];
 *     } exports[exports_count];
 * 
 *     util.U2 opens_count;
 *     {   util.U2 opens_index;
 *         util.U2 opens_flags;
 *         util.U2 opens_to_count;
 *         util.U2 opens_to_index[opens_to_count];
 *     } opens[opens_count];
 * 
 *     util.U2 uses_count;
 *     util.U2 uses_index[uses_count];
 * 
 *     util.U2 provides_count;
 *     {   util.U2 provides_index;
 *         util.U2 provides_with_count;
 *         util.U2 provides_with_index[provides_with_count];
 *     } provides[provides_count];
 * }
 */

type ModuleRequires struct{
	requires_index           util.U2
	requires_flags           util.U2
	requires_version_index   util.U2
}

type ModuleExports struct{
	exports_index util.U2
	exports_flags util.U2
	exports_to_index []util.U2
}

type ModuleOpens struct{
	opens_index     util.U2
	opens_flags     util.U2
	opens_to_index  []util.U2
}

type ModuleProvides struct{
	provides_index  util.U2
	provides_with_index []util.U2
}

type ModuleAttr struct{
	cp ConstantPool
	attribute_name_index	util.U2
	attribute_length	util.U4
	module_name_index	util.U2
	module_flags	util.U2
	module_version_index	util.U2

	requires []ModuleRequires
	exports []ModuleExports
	opens []ModuleOpens

	uses_index []util.U2

	provides []ModuleProvides
}

func (self *ModuleAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	self.cp = cf.constant_pool
	self.attribute_name_index = r.ReadU2()
	self.attribute_length = r.ReadU4()
	self.module_name_index = r.ReadU2()
	self.module_flags = r.ReadU2()
	self.module_version_index = r.ReadU2()
	self.requires = make([]ModuleRequires,r.ReadU2())
	for i := range self.requires {
		self.requires[i] = ModuleRequires{r.ReadU2(),r.ReadU2(),r.ReadU2()}
	}
	self.exports = make([]ModuleExports,r.ReadU2())
	for i := range self.exports {
		self.exports[i] = ModuleExports{r.ReadU2(),r.ReadU2(),r.ReadU2s()}
	}
	self.opens = make([]ModuleOpens,r.ReadU2())
	for i := range self.opens {
		self.opens[i] = ModuleOpens{r.ReadU2(),r.ReadU2(),r.ReadU2s()}
	}
	self.uses_index = r.ReadU2s()
	self.provides = make([]ModuleProvides,r.ReadU2())
	for i := range self.provides {
		self.provides[i] = ModuleProvides{r.ReadU2(),r.ReadU2s()}
	}
}

func (self ModuleAttr) AttributeNameIndex() util.U2{
	return self.attribute_name_index
}

func (self ModuleAttr) AttributeLength() util.U4{
	return self.attribute_length
}

func (self ModuleAttr) ModuleNameIndex() util.U2{
	return self.module_name_index
}


func (self ModuleAttr) ModuleFlags() util.U2{
	return self.module_flags
}


func (self ModuleAttr) ModuleVersionIndex() util.U2{
	return self.module_version_index
}

/*
func (self ModuleAttribute) Requires() util.U2{
	return self.requires
}


func (self ModuleAttribute) RequiresFlags() util.U2{
	return self.requires_flags
}


func (self ModuleAttribute) RequiresVersionIndex() util.U2{
	return self.requires_version_index
}


func (self ModuleAttribute) ExportsCount() util.U2{
	return self.exports_count
}


func (self ModuleAttribute) ExportsFlags() util.U2{
	return self.exports_flags
}


func (self ModuleAttribute) ExportsToCount() util.U2{
	return self.exports_to_count
}


func (self ModuleAttribute) ExportsToIndex[exportsToCount]() util.U2{
	return self.exports_to_index[exports_to_count]
}


func (self ModuleAttribute) OpensCount() util.U2{
	return self.opens_count
}


func (self ModuleAttribute) OpensFlags() util.U2{
	return self.opens_flags
}


func (self ModuleAttribute) OpensToCount() util.U2{
	return self.opens_to_count
}


func (self ModuleAttribute) OpensToIndex[opensToCount]() util.U2{
	return self.opens_to_index[opens_to_count]
}


func (self ModuleAttribute) UsesCount() util.U2{
	return self.uses_count
}


func (self ModuleAttribute) UsesIndex[usesCount]() util.U2{
	return self.uses_index[uses_count]
}


func (self ModuleAttribute) ProvidesCount() util.U2{
	return self.provides_count
}


func (self ModuleAttribute) ProvidesWithCount() util.U2{
	return self.provides_with_count
}


func (self ModuleAttribute) ProvidesWithIndex[providesWithCount]() util.U2{
	return self.provides_with_index[provides_with_count]
}
*/