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
type RuntimeVisibleAnnotationsAttr []RuntimeVisibleAnnotation

type RuntimeVisibleAnnotation struct{
	cp ConstantPool
	typeIndex u2 // annotation classname
	elementValuePairs []RvanElementValuePair
}

type RvanElementValuePair struct{
	elementNameIndex u2
	//tag u1
	value RvanElementValue
}

type RvanElementValue interface {
}
type RvanElementConstValue struct{
	cp ConstantPool
	const_value_index u2
}
type RvanElementEnumConstValue struct{
	cp ConstantPool
	type_name_index u2
	const_name_index u2
}
type RvanElementClassInfoValue struct{
	cp ConstantPool
	class_info_index u2
}
type RvanElementAnnotationValue struct{
	cp ConstantPool
	annotation RuntimeVisibleAnnotation
}
type RvanElementArrayValue struct{
	cp ConstantPool
	annotation RuntimeVisibleAnnotation
}

func parse(cf ClassFile,length u4,r *BigEndianReader)  RuntimeVisibleAnnotationsAttr {
	ans := []RuntimeVisibleAnnotation{}
	num_annotations := r.ReadU2()
	for i := 0; i< int(num_annotations); i++ {
		an := parseRuntimeVisibleAnnotation(cf,r)
		ans = append(ans,an)
	}
	return ans
}
func parseRuntimeVisibleAnnotation(cf ClassFile,r *BigEndianReader) RuntimeVisibleAnnotation {
	an := RuntimeVisibleAnnotation{cf.constant_pool,r.ReadU2(),
		make([]RvanElementValuePair,r.ReadU2())}
	for j := range an.elementValuePairs {
		pair := an.elementValuePairs[j]
		pair.elementNameIndex = r.ReadU2()
		pair.value = parseRvanElementValue(cf,r);
	}
	return an ;
}
func parseRvanElementValue(cf ClassFile,r *BigEndianReader) RvanElementValue {
	tag := r.ReadU1()
	switch tag {
		case 'e' : return RvanElementEnumConstValue{cf.constant_pool,r.ReadU2(),r.ReadU2()};
		case 'c' : return RvanElementClassInfoValue{cf.constant_pool,r.ReadU2()};
		case '@' : return parseRuntimeVisibleAnnotation(cf,r);
		case '[' : return parseRvanElementValue(cf,r);

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
	return RvanElementConstValue{cf.constant_pool,r.ReadU2()}
}

func (self RuntimeVisibleAnnotation) GetType() string{
	return self.cp.getUtf8String(self.typeIndex)
}
