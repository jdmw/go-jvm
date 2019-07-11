package classfile

/*
*ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html

StackMapTableAttr {
    u2              attribute_name_index;
    u4              attribute_length;
    u2              number_of_entries;
    stack_map_frame entries[number_of_entries];
}

union verification_type_info {
    Top_variable_info;
    Integer_variable_info;
    Float_variable_info;
    Long_variable_info;
    Double_variable_info;
    Null_variable_info;
    UninitializedThis_variable_info;
    Object_variable_info;
    Uninitialized_variable_info;
}

Top_variable_info {
    u1 tag = ITEM_Top; // 0
}

Integer_variable_info {
    u1 tag = ITEM_Integer; // 1
}

Float_variable_info {
    u1 tag = ITEM_Float; // 2
}

Null_variable_info {
    u1 tag = ITEM_Null; // 5
}

UninitializedThis_variable_info {
    u1 tag = ITEM_UninitializedThis; // 6
}

Object_variable_info {
    u1 tag = ITEM_Object; // 7
    u2 cpool_index;
}

Uninitialized_variable_info {
    u1 tag = ITEM_Uninitialized; // 8
    u2 offset;
}

Long_variable_info {
    u1 tag = ITEM_Long; // 4
}

Double_variable_info {
    u1 tag = ITEM_Double; // 3
}

union stack_map_frame {
    same_frame;
    same_locals_1_stack_item_frame;
    same_locals_1_stack_item_frame_extended;
    chop_frame;
    same_frame_extended;
    append_frame;
    full_frame;
}

same_frame {
    u1 frame_type = SAME; // 0-63
}

same_locals_1_stack_item_frame {
    u1 frame_type = SAME_LOCALS_1_STACK_ITEM; // 64-127
    verification_type_info stack[1];
}

same_locals_1_stack_item_frame_extended {
    u1 frame_type = SAME_LOCALS_1_STACK_ITEM_EXTENDED; // 247
    u2 offset_delta;
    verification_type_info stack[1];
}

chop_frame {
    u1 frame_type = CHOP; // 248-250
    u2 offset_delta;
}

same_frame_extended {
    u1 frame_type = SAME_FRAME_EXTENDED; // 251
    u2 offset_delta;
}

append_frame {
    u1 frame_type = APPEND; // 252-254
    u2 offset_delta;
    verification_type_info locals[frame_type - 251];
}

full_frame {
    u1 frame_type = FULL_FRAME; // 255
    u2 offset_delta;
    u2 number_of_locals;
    verification_type_info locals[number_of_locals];
    u2 number_of_stack_items;
    verification_type_info stack[number_of_stack_items];
}
*/
type StackMapTableAttr struct{
	cp ConstantPool
	attribute_name_index	u2
	attribute_length	u4
	number_of_entries	u2
	entries[number_of_entries]	stack_map_frame
}
func (self StackMapTableAttr) parse(cf ClassFile,length u4,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.attribute_name_index = r.ReadU2()
	self.attribute_length = r.ReadU4()
	self.number_of_entries = r.ReadU2()
	// TODO: self.entries[number_of_entries] = 
}


type verification_type_info struct{
	cp ConstantPool
}
func (self verification_type_info) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
}

type Top_variable_info struct{
	cp ConstantPool
	tag	u1
}
func (self Top_variable_info) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.tag = r.ReadU1()
}



func (self *Top_variable_info) Tag() u1{
	return self.tag
}
type Integer_variable_info struct{
	cp ConstantPool
	tag	u1
}
func (self Integer_variable_info) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.tag = r.ReadU1()
}



func (self *Integer_variable_info) Tag() u1{
	return self.tag
}
type Float_variable_info struct{
	cp ConstantPool
	tag	u1
}
func (self Float_variable_info) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.tag = r.ReadU1()
}



func (self *Float_variable_info) Tag() u1{
	return self.tag
}
type Null_variable_info struct{
	cp ConstantPool
	tag	u1
}
func (self Null_variable_info) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.tag = r.ReadU1()
}



func (self *Null_variable_info) Tag() u1{
	return self.tag
}
type UninitializedThis_variable_info struct{
	cp ConstantPool
	tag	u1
}
func (self UninitializedThis_variable_info) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.tag = r.ReadU1()
}



func (self *UninitializedThis_variable_info) Tag() u1{
	return self.tag
}
type Object_variable_info struct{
	cp ConstantPool
	tag	u1
	cpool_index	u2
}
func (self Object_variable_info) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.tag = r.ReadU1()
	self.cpool_index = r.ReadU2()
}



func (self *Object_variable_info) Tag() u1{
	return self.tag
}


func (self *Object_variable_info) CpoolIndex() u2{
	return self.cpool_index
}
type Uninitialized_variable_info struct{
	cp ConstantPool
	tag	u1
	offset	u2
}
func (self Uninitialized_variable_info) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.tag = r.ReadU1()
	self.offset = r.ReadU2()
}



func (self *Uninitialized_variable_info) Tag() u1{
	return self.tag
}


func (self *Uninitialized_variable_info) Offset() u2{
	return self.offset
}
type Long_variable_info struct{
	cp ConstantPool
	tag	u1
}
func (self Long_variable_info) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.tag = r.ReadU1()
}



func (self *Long_variable_info) Tag() u1{
	return self.tag
}
type Double_variable_info struct{
	cp ConstantPool
	tag	u1
}
func (self Double_variable_info) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.tag = r.ReadU1()
}



func (self *Double_variable_info) Tag() u1{
	return self.tag
}
type stack_map_frame struct{
	cp ConstantPool
}
func (self stack_map_frame) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
}

type same_frame struct{
	cp ConstantPool
	frame_type	u1
}
func (self same_frame) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.frame_type = r.ReadU1()
}



func (self *same_frame) FrameType() u1{
	return self.frame_type
}
type same_locals_1_stack_item_frame struct{
	cp ConstantPool
	frame_type	u1
	stack[1]	verification_type_info
}
func (self same_locals_1_stack_item_frame) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.frame_type = r.ReadU1()
	// TODO: self.stack[1] = 
}



func (self *same_locals_1_stack_item_frame) FrameType() u1{
	return self.frame_type
}


func (self *same_locals_1_stack_item_frame) Stack[1]() verification_type_info{
	return self.stack[1]
}
type same_locals_1_stack_item_frame_extended struct{
	cp ConstantPool
	frame_type	u1
	offset_delta	u2
	stack[1]	verification_type_info
}
func (self same_locals_1_stack_item_frame_extended) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.frame_type = r.ReadU1()
	self.offset_delta = r.ReadU2()
	// TODO: self.stack[1] = 
}



func (self *same_locals_1_stack_item_frame_extended) FrameType() u1{
	return self.frame_type
}


func (self *same_locals_1_stack_item_frame_extended) OffsetDelta() u2{
	return self.offset_delta
}


func (self *same_locals_1_stack_item_frame_extended) Stack[1]() verification_type_info{
	return self.stack[1]
}
type chop_frame struct{
	cp ConstantPool
	frame_type	u1
	offset_delta	u2
}
func (self chop_frame) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.frame_type = r.ReadU1()
	self.offset_delta = r.ReadU2()
}



func (self *chop_frame) FrameType() u1{
	return self.frame_type
}


func (self *chop_frame) OffsetDelta() u2{
	return self.offset_delta
}
type same_frame_extended struct{
	cp ConstantPool
	frame_type	u1
	offset_delta	u2
}
func (self same_frame_extended) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.frame_type = r.ReadU1()
	self.offset_delta = r.ReadU2()
}



func (self *same_frame_extended) FrameType() u1{
	return self.frame_type
}


func (self *same_frame_extended) OffsetDelta() u2{
	return self.offset_delta
}
type append_frame struct{
	cp ConstantPool
	frame_type	u1
	offset_delta	u2
	locals[frame_type	verification_type_info
}
func (self append_frame) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.frame_type = r.ReadU1()
	self.offset_delta = r.ReadU2()
	// TODO: self.locals[frame_type = 
}



func (self *append_frame) FrameType() u1{
	return self.frame_type
}


func (self *append_frame) OffsetDelta() u2{
	return self.offset_delta
}


func (self *append_frame) Locals[frameType() verification_type_info{
	return self.locals[frame_type
}
type full_frame struct{
	cp ConstantPool
	frame_type	u1
	offset_delta	u2
	number_of_locals	u2
	locals[number_of_locals]	verification_type_info
	number_of_stack_items	u2
	stack[number_of_stack_items]	verification_type_info
}
func (self full_frame) parse(cf ClassFile,r *BigEndianReader) {
	self.cp = cf.constant_pool
	self.frame_type = r.ReadU1()
	self.offset_delta = r.ReadU2()
	self.number_of_locals = r.ReadU2()
	// TODO: self.locals[number_of_locals] = 
	self.number_of_stack_items = r.ReadU2()
	// TODO: self.stack[number_of_stack_items] = 
}



func (self *full_frame) FrameType() u1{
	return self.frame_type
}


func (self *full_frame) OffsetDelta() u2{
	return self.offset_delta
}


func (self *full_frame) NumberOfLocals() u2{
	return self.number_of_locals
}


func (self *full_frame) Locals[numberOfLocals]() verification_type_info{
	return self.locals[number_of_locals]
}


func (self *full_frame) NumberOfStackItems() u2{
	return self.number_of_stack_items
}


func (self *full_frame) Stack[numberOfStackItems]() verification_type_info{
	return self.stack[number_of_stack_items]
}


Process finished with exit code 0
