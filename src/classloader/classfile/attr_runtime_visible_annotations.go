package classfile

/*
 * ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * RuntimeVisibleAnnotations_attribute {
 *     u2         attribute_name_index;
 *     u4         attribute_length;
 *     u2         num_annotations;
 *     annotation annotations[num_annotations];
 * }
 *
	RuntimeInvisibleAnnotations_attribute {
		u2         attribute_name_index;
		u4         attribute_length;
		u2         num_annotations;
		annotation annotations[num_annotations];
	}

 * annotation {
 *     u2 type_index;
 *     u2 num_element_value_pairs;
 *     {   u2            element_name_index;
 *         element_value value;
 *     } element_value_pairs[num_element_value_pairs];
 * }
 * 
 * element_value {
 *     u1 tag;
 *     union {
 *         u2 const_value_index;
 * 
 *         {   u2 type_name_index;
 *             u2 const_name_index;
 *         } enum_const_value;
 * 
 *         u2 class_info_index;
 * 
 *         annotation annotation_value;
 * 
 *         {   u2            num_values;
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
	typeIndex u2 // annotation classname
	elementValuePairs []AnElementValuePair
}

type AnElementValuePair struct{
	elementNameIndex u2
	//tag u1
	value AnElementValue
}

type AnElementValue interface {
}
type AnElementConstValue struct{
	cp ConstantPool
	const_value_index u2
}
type AnElementEnumConstValue struct{
	cp ConstantPool
	type_name_index u2
	const_name_index u2
}
type AnElementClassInfoValue struct{
	cp ConstantPool
	class_info_index u2
}
type AnElementAnnotationValue struct{
	cp ConstantPool
	annotation AnnotationInfo
}
type AnElementArrayValue struct{
	cp ConstantPool
	annotation AnElementValue
}

func (self *RuntimeVisibleAnnotationsAttr) parse(cf ClassFile,length u4,r *BigEndianReader)   {
	*self = *parseAnnotationsAtt(cf,r)
}
func (self *RuntimeInvisibleAnnotationsAttr) parse(cf ClassFile,length u4,r *BigEndianReader)   {
	*self = *parseAnnotationsAtt(cf,r)
}

func parseAnnotationsAtt(cf ClassFile,r *BigEndianReader) *([]AnnotationInfo)  {
	ans := []AnnotationInfo{}
	num_annotations := r.ReadU2()
	for i := 0; i< int(num_annotations); i++ {
		an := parseAnnotationInfo(cf,r)
		ans = append(ans,an)
	}
	return &ans
}


func parseAnnotationInfo(cf ClassFile,r *BigEndianReader) AnnotationInfo {
	an := AnnotationInfo{cf.constant_pool,r.ReadU2(),
		make([]AnElementValuePair,r.ReadU2())}
	for j := range an.elementValuePairs {
		pair := an.elementValuePairs[j]
		pair.elementNameIndex = r.ReadU2()
		pair.value = parseAnElementValue(cf,r);
	}
	return an ;
}
func parseAnElementValue(cf ClassFile,r *BigEndianReader) AnElementValue {
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
