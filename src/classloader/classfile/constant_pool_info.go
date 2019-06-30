package classfile

type ConstantPoolInfo interface{
  readInfo(r *BigEndianReader)
}

/* Constant pool tags */
//    Constant Type					        Value  class file	Java SE
const CONSTANT_Class				      = 7	  // 45.3			1.0.2
const CONSTANT_Fieldref			      = 9	  // 45.3			1.0.2
const CONSTANT_Methodref			    = 10  // 45.3			1.0.2
const CONSTANT_InterfaceMethodref	= 11  // 45.3			1.0.2
const CONSTANT_String				      = 8	  // 45.3			1.0.2
const CONSTANT_Integer			      = 3	  // 45.3			1.0.2
const CONSTANT_Float				      = 4	  // 45.3			1.0.2
const CONSTANT_Long				        = 5	  // 45.3			1.0.2
const CONSTANT_Double				      = 6	  // 45.3			1.0.2
const CONSTANT_NameAndType	      = 12  // 45.3			1.0.2
const CONSTANT_Utf8				        = 1	  // 45.3			1.0.2
const CONSTANT_MethodHandle		    = 15  // 51.0			7
const CONSTANT_MethodType			    = 16  // 51.0			7
const CONSTANT_Dynamic	          = 17
const CONSTANT_InvokeDynamic	    = 18  // 51.0			7
const CONSTANT_Module 			      = 19  // 53.0			9
const CONSTANT_Package 			      = 20  // 53.0			9

/*var funs = []func(r *BigEndianReader) ConstantPoolInfo {
	parseConstClassInfo
} */


func parseConstantPoolInfo(cp ConstantPool,r *BigEndianReader) (ConstantPoolInfo,int) {
	var info ConstantPoolInfo
	tag := r.ReadU1()

	checkAccFlag := 0
	switch tag {
		case CONSTANT_Class: info = &ConstClassInfo{cp,-1};break;
		case CONSTANT_Fieldref :
		case CONSTANT_Methodref :
		case CONSTANT_InterfaceMethodref : info = &ConstMemberRefInfo{cp,-1,-1};break;
		case CONSTANT_String : info = &ConstStringInfo{cp,-1};break;
		case CONSTANT_Integer: info = &ConstIntegerInfo{};break;
		case CONSTANT_Float: info = &ConstFloatInfo{};break;
		case CONSTANT_Long : info = &ConstLongInfo{};break;
		case CONSTANT_Double				      : info = &ConstDoubleInfo{};break;
		case CONSTANT_NameAndType	      : info = &ConstNameAndTypeInfo{};break;
		case CONSTANT_Utf8				        : info = &ConstUtf8Info{};break;
		case CONSTANT_MethodHandle		    : info = &ConstMethodHandleInfo{};break;
		case CONSTANT_MethodType			    : info = &ConstMethodTypeInfo{};break;
		case CONSTANT_Dynamic	          : info = &ConstDynamicInfo{};break;
		case CONSTANT_InvokeDynamic	    : info = &ConstInvokeDynamicInfo{};break;
		case CONSTANT_Module 			      : info = &ConstModuleInfo{};checkAccFlag = ACC_MODULE;break;
		case CONSTANT_Package 			      : info = &ConstPackageInfo{};checkAccFlag = ACC_MODULE;break;
		default:panic("Unsupported tag " + string(tag))
	}

	info.readInfo(r)
	return info,checkAccFlag
}


