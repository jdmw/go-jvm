package classfile

import "../../util"

/*
 * ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * RuntimeVisibleAnnotations_attribute {
 *     util.U2         attribute_name_index;
 *     util.U4         attribute_length;
 *     util.U2         num_annotations;
 *     annotation annotations[num_annotations];
 * }
 *
	RuntimeInvisibleAnnotations_attribute {
		util.U2         attribute_name_index;
		util.U4         attribute_length;
		util.U2         num_annotations;
		annotation annotations[num_annotations];
	}

 * annotation {
 *     util.U2 type_index;
 *     util.U2 num_element_value_pairs;
 *     {   util.U2            element_name_index;
 *         element_value value;
 *     } element_value_pairs[num_element_value_pairs];
 * }
 * 
 * element_value {
 *     util.U1 tag;
 *     union {
 *         util.U2 const_value_index;
 * 
 *         {   util.U2 type_name_index;
 *             util.U2 const_name_index;
 *         } enum_const_value;
 * 
 *         util.U2 class_info_index;
 * 
 *         annotation annotation_value;
 * 
 *         {   util.U2            num_values;
 *             element_value values[num_values];
 *         } array_value;
 *     } value;
 * }

	Table 4.7.16.1-A. Interpretation of tag values as types

	tag Item	Type	value Item	Constant Type
	B	byte	const_value_index	CONSTANT_Integer
	C	char	const_value_index	CONSTANT_Integer
	D	double	const_value_index	CONSTANT_Double
	F	float	const_value_index	CONSTANT_Float
	I	int	const_value_index	CONSTANT_Integer
	J	long	const_value_index	CONSTANT_Long
	S	short	const_value_index	CONSTANT_Integer
	Z	boolean	const_value_index	CONSTANT_Integer
	s	String	const_value_index	CONSTANT_Utf8
	e	Enum type	enum_const_value	Not applicable
	c	Class	class_info_index	Not applicable
	@	Annotation type	annotation_value	Not applicable
	[	Array type	array_value	Not applicable
 */
type RuntimeVisibleAnnotationsAttr []AnnotationInfo
type RuntimeInvisibleAnnotationsAttr []AnnotationInfo

type AnnotationInfo struct{
	cp ConstantPool
	typeIndex util.U2 // annotation classname
	elementValuePairs []AnElementValuePair
}

type AnElementValuePair struct{
	elementNameIndex util.U2
	//tag util.U1
	value AnElementValue
}

type AnElementValue interface {
}
type AnElementConstValue struct{
	cp ConstantPool
	const_value_index util.U2
}
type AnElementEnumConstValue struct{
	cp ConstantPool
	type_name_index util.U2
	const_name_index util.U2
}
type AnElementClassInfoValue struct{
	cp ConstantPool
	class_info_index util.U2
}
type AnElementAnnotationValue struct{
	cp ConstantPool
	annotation AnnotationInfo
}
type AnElementArrayValue struct{
	cp ConstantPool
	annotation AnElementValue
}

func (self *RuntimeVisibleAnnotationsAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader)   {
	*self = *parseAnnotationsAtt(cf,r)
}
func (self *RuntimeInvisibleAnnotationsAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader)   {
	*self = *parseAnnotationsAtt(cf,r)
}

func parseAnnotationsAtt(cf ClassFile,r *util.BigEndianReader) *([]AnnotationInfo)  {
	ans := []AnnotationInfo{}
	num_annotations := r.ReadU2()
	for i := 0; i< int(num_annotations); i++ {
		an := parseAnnotationInfo(cf,r)
		ans = append(ans,an)
	}
	return &ans
}

func parseAnnotationInfo(cf ClassFile,r *util.BigEndianReader) AnnotationInfo {
	an := AnnotationInfo{cf.constant_pool,r.ReadU2(),
		make([]AnElementValuePair,r.ReadU2())}
	for j := range an.elementValuePairs {
		pair := an.elementValuePairs[j]
		pair.elementNameIndex = r.ReadU2()
		pair.value = parseAnElementValue(cf,r);
	}
	return an ;
}
func parseAnElementValue(cf ClassFile,r *util.BigEndianReader) AnElementValue {
	tag := r.ReadU1()
	switch tag {
		case 'e' : return AnElementEnumConstValue{cf.constant_pool,r.ReadU2(),r.ReadU2()};
		case 'c' : return AnElementClassInfoValue{cf.constant_pool,r.ReadU2()};
		case '@' : return parseAnnotationInfo(cf,r);
		case '[' : return parseAnElementValue(cf,r);

		case 'B' :
		case 'C' :
		case 'D' :
		case 'F' :
		case 'I' :
		case 'J' :
		case 'S' :
		case 'Z' :
		case 's' :
		default:
	}
	return AnElementConstValue{cf.constant_pool,r.ReadU2()}
}

func (self AnnotationInfo) GetType() string{
	return self.cp.getUtf8String(self.typeIndex)
}





/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 *
 * RuntimeVisibleParameterAnnotations_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U1 num_parameters;
 *     {   util.U2         num_annotations;
 *         annotation annotations[num_annotations];
 *     } parameter_annotations[num_parameters];
 * }

 * RuntimeInvisibleParameterAnnotations_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U1 num_parameters;
 *     {   util.U2         num_annotations;
 *         annotation annotations[num_annotations];
 *     } parameter_annotations[num_parameters];
 * }
*/
type RuntimeInvisibleParameterAnnotationsAttr []ParameterAnnotationInfo
type RuntimeVisibleParameterAnnotationsAttr []ParameterAnnotationInfo
type ParameterAnnotationInfo []AnnotationInfo

func (self *RuntimeInvisibleParameterAnnotationsAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	*self = *parseParameterAnnotationInfo(cf,r)
}
func (self *RuntimeVisibleParameterAnnotationsAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	*self = *parseParameterAnnotationInfo(cf,r)
}


func  parseParameterAnnotationInfo(cf ClassFile,r *util.BigEndianReader) *([]ParameterAnnotationInfo) {
	parameter_annotations := make([]ParameterAnnotationInfo,r.ReadU2())
	for i := range parameter_annotations {
		parameter_annotations[i] = *parseAnnotationsAtt(cf,r)
	}
	return &parameter_annotations
}


/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 *
 * AnnotationDefault_attribute {
 *     util.U2            attribute_name_index;
 *     util.U4            attribute_length;
 *     element_value default_value;
 * }
*/
type AnnotationDefaultAttr struct{
	default_value AnElementValue
}
func (self *AnnotationDefaultAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	self.default_value = parseAnElementValue(cf,r)
}

