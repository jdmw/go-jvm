package classfile

import "../../util"



/**
参考： https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.1

ClassFile {
    util.U4             magic;
    util.U2             minor_version;
    util.U2             major_version;
    util.U2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    util.U2             access_flags;
    util.U2             this_class;
    util.U2             super_class;
    util.U2             interfaces_count;
    util.U2             interfaces[interfaces_count];
    util.U2             fields_count;
    field_info     fields[fields_count];
    util.U2             methods_count;
    method_info    methods[methods_count];
    util.U2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

const CLASSFILE_MAGICNUM = 0xCAFEBABE

type ClassFile struct {
	//util.U4            magic
	minor_version   util.U2
	major_version   util.U2
	constant_pool 	ConstantPool
	access_flags    util.U2
	this_class      util.U2
	super_class     util.U2
	interfaces      []util.U2
	fields          []FieldInfo
	methods         []MethodInfo
	attributes      []AttributeInfo
}

func ParseClassFile(data []byte) *ClassFile {
	cf := ClassFile{}
	r := util.NewBigEndianReader(data)
	if( CLASSFILE_MAGICNUM != r.ReadU4()){
		panic("java.lang.ClassFormatError: magic!")
	}

	cf.minor_version = r.ReadU2()
	cf.major_version = r.ReadU2()

	accFlag := 0
	cf.constant_pool,accFlag = parseConstPool(int(r.ReadU2()-1),r)

	cf.access_flags = r.ReadU2()
	if( accFlag & ACC_MODULE > 0 && cf.access_flags & ACC_MODULE == 0){
		panic("Acc_Module is not set")
	}

	cf.this_class = r.ReadU2()
	cf.super_class = r.ReadU2()

	cf.interfaces = r.ReadU2s()

	fields := make([]FieldInfo,r.ReadU2())
	for i:= range fields {
		fields[i] = parseFieldInfo(cf,r)
	}
	cf.fields = fields

	methods := make([]MethodInfo,r.ReadU2())
	for i:= range methods {
		methods[i] = parseMethodInfo(cf,r)
	}
	cf.methods = methods

	attrs := make([]AttributeInfo,r.ReadU2())
	for i := range attrs {
		attrs[i] = parseAttributeInfo(cf,r)
	}
	cf.attributes = attrs
	return &cf
}

func (self ClassFile) ClassInfo() (uint16,string,string,[]string,ConstantPool,[]FieldInfo,[]MethodInfo,[]AttributeInfo,){
	cp := self.constant_pool
	interfaceNames := make([]string,len(self.interfaces))
	for i,index := range self.interfaces{
		interfaceNames[i] = cp.getClassName(index)
	}
	return uint16(self.access_flags),cp.getClassName(self.this_class), cp.getClassName(self.super_class),interfaceNames,self.constant_pool,self.fields,self.methods,self.attributes
}