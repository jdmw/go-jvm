package classfile

/**
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}

CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}

CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */

type ConstClassInfo struct {
	cp ConstantPool
	name_index u2
}

func (self *ConstClassInfo) readInfo(r *BigEndianReader) {
	self.name_index = r.ReadU2()
}

func (self *ConstClassInfo) Name() string {
	return self.cp.getUtf8String(self.name_index)
}

type ConstMemberRefInfo struct {
	cp ConstantPool
	class_index			u2
	name_and_type_index u2
}

func (self *ConstMemberRefInfo) readInfo(r *BigEndianReader) {
	self.class_index = r.ReadU2()
	self.name_and_type_index = r.ReadU2()
}

func (self *ConstMemberRefInfo) ClassName() string {
	return self.cp.getUtf8String(self.class_index)
}

func (self *ConstMemberRefInfo) NameAndType() (string,string) {
	return self.cp.getNameAndType(self.name_and_type_index)
}

/**
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
 */

type ConstNameAndTypeInfo struct {
	cp ConstantPool
	name_index u2
	descriptor_index u2
}

func (self *ConstNameAndTypeInfo) readInfo(r *BigEndianReader) {
	self.name_index = r.ReadU2()
	self.descriptor_index = r.ReadU2()
}

func (self *ConstNameAndTypeInfo) Name() string {
	return self.cp.getUtf8String(self.name_index)
}

func (self *ConstNameAndTypeInfo) Descriptor() string {
	return self.cp.getUtf8String(self.descriptor_index)
}


/**
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/

const REF_getField          = 1   // CONSTANT_Fieldref_info
const REF_getStatic         = 2
const REF_putField          = 3
const REF_putStatic         = 4
const REF_invokeVirtual     = 5   // CONSTANT_Methodref_info
const REF_newInvokeSpecial  = 8
const REF_invokeStatic      = 6   // CONSTANT_Methodref_info  or  CONSTANT_InterfaceMethodref_info
const REF_invokeSpecial     = 7
const REF_invokeInterface   = 9

type ConstMethodHandleInfo struct {
	cp ConstantPool
	reference_kind u1
	reference_index u2
}

func (self *ConstMethodHandleInfo) readInfo(r *BigEndianReader) {
	self.reference_kind = r.ReadU1()
	self.reference_index = r.ReadU2()
}

func (self *ConstMethodHandleInfo) ClassMember()  (string,u1,string,string) {
	info := self.cp[self.reference_index].(*ConstMemberRefInfo)
	classname := self.cp.getClassName(info.class_index)
	name,desp := self.cp.getNameAndType(info.name_and_type_index)
	return classname,self.reference_kind,name,desp
}

/**
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/

type ConstMethodTypeInfo struct {
	cp ConstantPool
	descriptor_index u2
}

func (self *ConstMethodTypeInfo) readInfo(r *BigEndianReader) {
	self.descriptor_index = r.ReadU2()
}

func (self *ConstMethodTypeInfo) Descriptor()  string {
	return self.cp.getUtf8String(self.descriptor_index)
}


/**
CONSTANT_Dynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}

CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/

type ConstDynamicInfo struct {
	cp ConstantPool
	bootstrap_method_attr_index u2
	name_and_type_index u2
}

func (self *ConstDynamicInfo) readInfo(r *BigEndianReader) {
	self.bootstrap_method_attr_index = r.ReadU2()
	self.name_and_type_index = r.ReadU2()
}

func (self *ConstDynamicInfo) NameAndType()  (string,string) {
	return self.cp.getNameAndType(self.name_and_type_index)
}

type ConstInvokeDynamicInfo struct {
	cp ConstantPool
	bootstrap_method_attr_index u2
	name_and_type_index u2
}

func (self *ConstInvokeDynamicInfo) readInfo(r *BigEndianReader) {
	self.bootstrap_method_attr_index = r.ReadU2()
	self.name_and_type_index = r.ReadU2()
}

func (self *ConstInvokeDynamicInfo) NameAndType()  (string,string) {
	return self.cp.getNameAndType(self.name_and_type_index)
}


/**
CONSTANT_Module_info {
    u1 tag;
    u2 name_index;
}
check :the ACC_MODULE flag is set
 */


type ConstModuleInfo struct {
	cp ConstantPool
	name_index u2
}

func (self *ConstModuleInfo) readInfo(r *BigEndianReader) {
	self.name_index = r.ReadU2()
}

func (self *ConstModuleInfo) Name()  string {
	return self.cp.getUtf8String(self.name_index)
}


/**
CONSTANT_Package_info {
    u1 tag;
    u2 name_index;
}
check: the ACC_MODULE flag is set.
 */

type ConstPackageInfo struct {
	cp ConstantPool
	name_index u2
}

func (self *ConstPackageInfo) readInfo(r *BigEndianReader) {
	self.name_index = r.ReadU2()
}

func (self *ConstPackageInfo) Name()  string {
	return self.cp.getUtf8String(self.name_index)
}
