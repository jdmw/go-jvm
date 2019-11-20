package classfile

import "../../util"

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html
 * 
 * BootstrapMethods_attribute {
 *     util.U2 attribute_name_index;
 *     util.U4 attribute_length;
 *     util.U2 num_bootstrap_methods;
 *     {   util.U2 bootstrap_method_ref;
 *         util.U2 num_bootstrap_arguments;
 *         util.U2 bootstrap_arguments[num_bootstrap_arguments];
 *     } bootstrap_methods[num_bootstrap_methods];
 * }
 */

type BootstrapMethods struct {
	cp ConstantPool
	bootstrap_method_ref util.U2
	bootstrap_arguments []util.U2
}
type BootstrapMethodsAttr []BootstrapMethods

func (self *BootstrapMethodsAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	*self = make([]BootstrapMethods, r.ReadU2())
	for i := range *self {
		(*self)[i] = BootstrapMethods{cf.constant_pool,r.ReadU2(),r.ReadU2s()}
	}
}
