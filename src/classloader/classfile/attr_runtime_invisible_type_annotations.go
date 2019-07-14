package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * RuntimeInvisibleTypeAnnotations_attribute {
 *     u2              attribute_name_index;
 *     u4              attribute_length;
 *     u2              num_annotations;
 *     type_annotation annotations[num_annotations];
 * }
 */
type RuntimeInvisibleTypeAnnotationsAttr  []TypeAnnotation
/**
type_annotation {
	u1 target_type;
	union {
		type_parameter_target;
		supertype_target;
		type_parameter_bound_target;
		empty_target;
		method_formal_parameter_target;
		throws_target;
		localvar_target;
		catch_target;
		offset_target;
		type_argument_target;
	} target_info;

	type_path target_path;
	u2 type_index;

	u2 num_element_value_pairs;
	{ u2 element_name_index;
	element_value value;
	} element_value_pairs[num_element_value_pairs];
}
 */
type TypeAnnotation struct {

}
/*func (self RuntimeInvisibleTypeAnnotationsAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	//self.cp = cf.constant_pool
	annotations := make([]TypeAnnotation,r.ReadU2())
	for i,an := range annotations {

	}
}
*/

