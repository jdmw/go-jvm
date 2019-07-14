package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * BootstrapMethods_attribute {
 *     u2 attribute_name_index;
 *     u4 attribute_length;
 *     u2 num_bootstrap_methods;
 *     {   u2 bootstrap_method_ref;
 *         u2 num_bootstrap_arguments;
 *         u2 bootstrap_arguments[num_bootstrap_arguments];
 *     } bootstrap_methods[num_bootstrap_methods];
 * }
 */

type BootstrapMethods struct {
	cp ConstantPool
	bootstrap_method_ref u2
	bootstrap_arguments []u2
}
type BootstrapMethodsAttr []BootstrapMethods

func (self *BootstrapMethodsAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	*self = make([]BootstrapMethods, r.ReadU2())
	for i := range *self {
		(*self)[i] = BootstrapMethods{cf.constant_pool,r.ReadU2(),r.ReadU2s()}
	}
}
