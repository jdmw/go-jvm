package classfile

/**
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
 */

const CLASSFILE_MAGICNUM = 0xCAFEBABE

type ClassFile struct {
//u4            magic
minor_version   u2             
major_version   u2
constant_pool 	ConstantPool
access_flags    u2             
this_class      u2             
super_class     u2
interfaces      []u2
fields          []FieldInfo
methods         []MethodInfo
attributes      []AttributeInfo
}

func ParseClassFile(data []byte) ClassFile {
	cf := ClassFile{}
	r := &BigEndianReader{data}
	if( CLASSFILE_MAGICNUM != r.ReadU4()){
		panic("java.lang.ClassFormatError: magic!")
	}

	cf.minor_version = r.ReadU2()
	cf.major_version = r.ReadU2()

	accFlag := 0
	cf.constant_pool,accFlag = parseConstPool(int(r.ReadU2()),r)

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

	attrs := make([]AttributeInfo,r.ReadU2())
	for i := range attrs {
		attrs[i] = parseAttributeInfo(r)
	}
	cf.attributes = attrs
	return cf
}