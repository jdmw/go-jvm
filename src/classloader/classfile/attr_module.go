package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * Module_attribute {
 *     u2 attribute_name_index;
 *     u4 attribute_length;
 * 
 *     u2 module_name_index;
 *     u2 module_flags;
 *     u2 module_version_index;
 * 
 *     u2 requires_count;
 *     {   u2 requires_index;
 *         u2 requires_flags;
 *         u2 requires_version_index;
 *     } requires[requires_count];
 * 
 *     u2 exports_count;
 *     {   u2 exports_index;
 *         u2 exports_flags;
 *         u2 exports_to_count;
 *         u2 exports_to_index[exports_to_count];
 *     } exports[exports_count];
 * 
 *     u2 opens_count;
 *     {   u2 opens_index;
 *         u2 opens_flags;
 *         u2 opens_to_count;
 *         u2 opens_to_index[opens_to_count];
 *     } opens[opens_count];
 * 
 *     u2 uses_count;
 *     u2 uses_index[uses_count];
 * 
 *     u2 provides_count;
 *     {   u2 provides_index;
 *         u2 provides_with_count;
 *         u2 provides_with_index[provides_with_count];
 *     } provides[provides_count];
 * }
 */

type ModuleRequires struct{
	requires_index           u2
	requires_flags           u2
	requires_version_index   u2
}

type ModuleExports struct{
	exports_index u2
	exports_flags u2
	exports_to_index []u2
}

type ModuleOpens struct{
	opens_index     u2
	opens_flags     u2
	opens_to_index  []u2
}

type ModuleProvides struct{
	provides_index  u2
	provides_with_index []u2
}

type ModuleAttr struct{
	cp ConstantPool
	attribute_name_index	u2
	attribute_length	u4
	module_name_index	u2
	module_flags	u2
	module_version_index	u2

	requires []ModuleRequires
	exports []ModuleExports
	opens []ModuleOpens

	uses_index []u2

	provides []ModuleProvides
}

func (self *ModuleAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
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

func (self ModuleAttribute) AttributeNameIndex() u2{
	return self.attribute_name_index
}

func (self ModuleAttribute) AttributeLength() u4{
	return self.attribute_length
}

func (self ModuleAttribute) ModuleNameIndex() u2{
	return self.module_name_index
}


func (self ModuleAttribute) ModuleFlags() u2{
	return self.module_flags
}


func (self ModuleAttribute) ModuleVersionIndex() u2{
	return self.module_version_index
}

/*
func (self ModuleAttribute) Requires() u2{
	return self.requires
}


func (self ModuleAttribute) RequiresFlags() u2{
	return self.requires_flags
}


func (self ModuleAttribute) RequiresVersionIndex() u2{
	return self.requires_version_index
}


func (self ModuleAttribute) ExportsCount() u2{
	return self.exports_count
}


func (self ModuleAttribute) ExportsFlags() u2{
	return self.exports_flags
}


func (self ModuleAttribute) ExportsToCount() u2{
	return self.exports_to_count
}


func (self ModuleAttribute) ExportsToIndex[exportsToCount]() u2{
	return self.exports_to_index[exports_to_count]
}


func (self ModuleAttribute) OpensCount() u2{
	return self.opens_count
}


func (self ModuleAttribute) OpensFlags() u2{
	return self.opens_flags
}


func (self ModuleAttribute) OpensToCount() u2{
	return self.opens_to_count
}


func (self ModuleAttribute) OpensToIndex[opensToCount]() u2{
	return self.opens_to_index[opens_to_count]
}


func (self ModuleAttribute) UsesCount() u2{
	return self.uses_count
}


func (self ModuleAttribute) UsesIndex[usesCount]() u2{
	return self.uses_index[uses_count]
}


func (self ModuleAttribute) ProvidesCount() u2{
	return self.provides_count
}


func (self ModuleAttribute) ProvidesWithCount() u2{
	return self.provides_with_count
}


func (self ModuleAttribute) ProvidesWithIndex[providesWithCount]() u2{
	return self.provides_with_index[provides_with_count]
}
*/