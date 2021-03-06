package classfile

import "../../util"


/**
attribute_info {
    util.U2 attribute_name_index;
    util.U4 attribute_length;
    util.U1 info[attribute_length];
}
 */

type Attributes []AttributeInfo

func parseAttributeInfo(cf ClassFile,r *util.BigEndianReader) AttributeInfo{
	attributeName := cf.constant_pool.getUtf8String(r.ReadU2())
	length := r.ReadU4()
	//bytes := r.ReadByteArray(length)
	info := newAttributeInfo(attributeName,r)
	info.parse(cf,length,r)
	return info
}

func newAttributeInfo(attributeName string ,r *util.BigEndianReader) AttributeInfo{
	switch attributeName {
		case "ConstantValue" : return &ConstantValueAttr{}
		case "Code" : return &CodeAttr{}
		case "StackMapTable" : return &StackMapTableAttr{}
		case "Exceptions" : return &ExceptionsAttr{}
		case "InnerClasses" : return &InnerClassesAttr{}
		case "EnclosingMethod" : return &EnclosingMethodAttr{}
		case "Synthetic" : return &SyntheticAttr{}
		case "Signature" : return &SignatureAttr{}
		case "SourceFile" : return &SourceFileAttr{}
		case "SourceDebugExtension" : return &SourceDebugExtensionAttr{}
		case "LineNumberTable" : return &LineNumberTableAttr{}
		case "LocalVariableTable" : return &LocalVariableTableAttr{}
		case "LocalVariableTypeTable" : return &LocalVariableTypeTableAttr{}
		case "Deprecated" : return &DeprecatedAttr{}
		case "RuntimeVisibleAnnotations" : return &RuntimeVisibleAnnotationsAttr{}
		case "RuntimeInvisibleAnnotations" : return &RuntimeInvisibleAnnotationsAttr{}
		case "RuntimeVisibleParameterAnnotations" : return &RuntimeVisibleParameterAnnotationsAttr{}
		case "RuntimeInvisibleParameterAnnotations" : return &RuntimeInvisibleParameterAnnotationsAttr{}
		//  case "RuntimeVisibleTypeAnnotations" : return &RuntimeVisibleTypeAnnotationsAttr{}
		//  case "RuntimeInvisibleTypeAnnotations" : return &RuntimeInvisibleTypeAnnotationsAttr{}
		case "AnnotationDefault" : return &AnnotationDefaultAttr{}
		case "BootstrapMethods" : return &BootstrapMethodsAttr{}
		case "MethodParameters" : return &MethodParametersAttr{}
		case "Module" : return &ModuleAttr{}
		case "ModulePackages" : return &ModulePackagesAttr{}
		case "ModuleMainClass" : return &ModuleMainClassAttr{}
		case "NestHost" : return &NestHostAttr{}
		case "NestMembers" : return &NestMembersAttr{}
		default: return &UnknownAttr{}
	}
}


func parseAttributes(cf ClassFile,length util.U2,r *util.BigEndianReader) Attributes{
	attrs := make([]AttributeInfo,length)
	for i:=0;i < int(length);i++ {
		attrs[i] = parseAttributeInfo(cf,r)
	}
	return attrs
}

type AttributeInfo interface {
	parse(cf ClassFile,length util.U4,r *util.BigEndianReader)
}

//var attrParser = map[string]AttributeInfoParser{}


/**
Table 4.7-A. Predefined class file attributes (by section)
Attribute	Section	class file	Java SE
ConstantValue	§4.7.2	45.3	1.0.2
Code	§4.7.3	45.3	1.0.2
StackMapTable	§4.7.4	50.0	6
Exceptions	§4.7.5	45.3	1.0.2
InnerClasses	§4.7.6	45.3	1.1
EnclosingMethod	§4.7.7	49.0	5.0
Synthetic	§4.7.8	45.3	1.1
Signature	§4.7.9	49.0	5.0
SourceFile	§4.7.10	45.3	1.0.2
SourceDebugExtension	§4.7.11	49.0	5.0
LineNumberTable	§4.7.12	45.3	1.0.2
LocalVariableTable	§4.7.13	45.3	1.0.2
LocalVariableTypeTable	§4.7.14	49.0	5.0
Deprecated	§4.7.15	45.3	1.1
RuntimeVisibleAnnotations	§4.7.16	49.0	5.0
RuntimeInvisibleAnnotations	§4.7.17	49.0	5.0
RuntimeVisibleParameterAnnotations	§4.7.18	49.0	5.0
RuntimeInvisibleParameterAnnotations	§4.7.19	49.0	5.0
RuntimeVisibleTypeAnnotations	§4.7.20	52.0	8
RuntimeInvisibleTypeAnnotations	§4.7.21	52.0	8
AnnotationDefault	§4.7.22	49.0	5.0
BootstrapMethods	§4.7.23	51.0	7
MethodParameters	§4.7.24	52.0	8
Module			§4.7.25	53.0	9
ModulePackages   §4.7.26	53.0	9
ModuleMainClass   §4.7.27	53.0	9
NestHost         §4.7.28	55.0	11
NestMembers        §4.7.29	55.0	11
 */

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 *
 * Synthetic_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length; // must be 0
 * }
*/
type SyntheticAttr struct{
}
func (self SyntheticAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
}


/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 *
 * Deprecated_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 * }
*/
type DeprecatedAttr struct{
}
func (self DeprecatedAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
}

type UnknownAttr struct {
	data []byte
}

func (self *UnknownAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	self.data = r.ReadByteArray(length)
}

func (self Attributes) GetCodeAttr() *CodeAttr{
	for _,e := range self {
		switch e.(type) {
			case *CodeAttr : return e.(*CodeAttr) ;
		}
	}
	return nil
}