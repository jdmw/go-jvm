package classfile

import "../../util"


import "fmt"

/*
   *ref:https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html

   StackMapTable_attribute {
       util.U2              attribute_name_index;
       util.U4              attribute_length;
       util.U2              number_of_entries;
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
       util.U1 tag = ITEM_Top; /* 0
}

Integer_variable_info {
util.U1 tag = ITEM_Integer; /* 1
}

Float_variable_info {
util.U1 tag = ITEM_Float; /* 2
}

Null_variable_info {
util.U1 tag = ITEM_Null; /* 5
}

UninitializedThis_variable_info {
util.U1 tag = ITEM_UninitializedThis; /* 6
}

Object_variable_info {
util.U1 tag = ITEM_Object; /* 7
util.U2 cpool_index;
}

Uninitialized_variable_info {
util.U1 tag = ITEM_Uninitialized; /* 8
util.U2 offset;
}

Long_variable_info {
util.U1 tag = ITEM_Long; /* 4
}

Double_variable_info {
util.U1 tag = ITEM_Double; /* 3
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
util.U1 frame_type = SAME; /* 0-63
}

same_locals_1_stack_item_frame {
util.U1 frame_type = SAME_LOCALS_1_STACK_ITEM; /* 64-127
verification_type_info stack[1];
}

same_locals_1_stack_item_frame_extended {
util.U1 frame_type = SAME_LOCALS_1_STACK_ITEM_EXTENDED; /* 247
util.U2 offset_delta;
verification_type_info stack[1];
}

chop_frame {
util.U1 frame_type = CHOP; /* 248-250
util.U2 offset_delta;
}

same_frame_extended {
util.U1 frame_type = SAME_FRAME_EXTENDED; /* 251
util.U2 offset_delta;
}

append_frame {
util.U1 frame_type = APPEND; /* 252-254
util.U2 offset_delta;
verification_type_info locals[frame_type - 251];
}

full_frame {
util.U1 frame_type = FULL_FRAME; /* 255
util.U2 offset_delta;
util.U2 number_of_locals;
verification_type_info locals[number_of_locals];
util.U2 number_of_stack_items;
verification_type_info stack[number_of_stack_items];
}
*/

const ITEM_Top =  0
const ITEM_Integer =  1
const ITEM_Float =  2
const ITEM_Double =  3
const ITEM_Long =  4
const ITEM_Null =  5
const ITEM_UninitializedThis =  6
const ITEM_Object =  7
const ITEM_Uninitialized =  8


type StackMapTableAttr []StackMapFrame
type StackMapFrame struct{
	cp ConstantPool
	frame_type	util.U1
	offset_delta util.U2
	locals	[]VerificationTypeInfo
	stack []VerificationTypeInfo
}

func (self *StackMapTableAttr) parse(cf ClassFile,length util.U4,r *util.BigEndianReader) {
	number_of_entries := r.ReadU2()
	entries := make([]StackMapFrame,number_of_entries)
	for i := 0;i<int(number_of_entries);i++ {
		entry := StackMapFrame{}
		entry.cp = cf.constant_pool
		entry.frame_type = r.ReadU1()
		entry.parse(cf,r)
		entries[i] = entry
		//fmt.Printf("frame_type = %v ,offset_delta = %v ,locals %v ,stack %v\n",entry.frame_type,entry.offset_delta,entry.locals,entry.stack)
	}
	*self = entries
}

/**
util.U1 frame_type = SAME; /* 0-63
util.U1 frame_type = SAME_LOCALS_1_STACK_ITEM; /* 64-127
util.U1 frame_type = SAME_LOCALS_1_STACK_ITEM_EXTENDED; /* 247
util.U1 frame_type = CHOP; /* 248-250
util.U1 frame_type = SAME_FRAME_EXTENDED; /* 251
util.U1 frame_type = APPEND; /* 252-254
util.U1 frame_type = FULL_FRAME; /* 255
 */
func (self *StackMapFrame) parse(cf ClassFile,r *util.BigEndianReader){
	frame_type := self.frame_type
	if  frame_type < 248 {
		if frame_type < 64 { // frame_type = SAME; /* 0-63
			self.parseSameFrame(cf, r)
		} else if frame_type <= 127 { // frame_type = SAME_LOCALS_1_STACK_ITEM; /* 64-127
			self.parseSameLocals1StackItemFrame(cf, r)
		} else if frame_type == 247 { //frame_type = SAME_LOCALS_1_STACK_ITEM_EXTENDED; /* 247
			self.parseSameLocals1StackItemFrameExtended(cf, r)
		} else {
			// not frame_type
		}
	}else {
		if self.frame_type < 251{ // frame_type = CHOP; /* 248-250
			self.parseChopFrame(cf, r)
		} else if self.frame_type == 251 { //frame_type = SAME_FRAME_EXTENDED; /* 251
			self.parseSameFrameExtended(cf, r)
		} else if frame_type < 255 {
			self.parseAppendFrame(cf, r)
		} else {
			self.parseFullFrame(cf, r)
		}
	}
}

func (self *StackMapFrame) FrameType() util.U1{
	return self.frame_type
}

func (self *StackMapFrame) OffsetDelta() util.U1{
	return self.frame_type
}

func (self *StackMapFrame) Locals() []VerificationTypeInfo {
	return self.locals
}

func (self *StackMapFrame) Stack() []VerificationTypeInfo{
	return self.stack
}

func (self *StackMapFrame) parseSameFrame(cf ClassFile,r *util.BigEndianReader) {
	self.offset_delta = util.U2(self.frame_type)
}

func (self *StackMapFrame) parseSameLocals1StackItemFrame(cf ClassFile,r *util.BigEndianReader) {
	//self.stack = make([]VerificationTypeInfo,1)
	self.stack= []VerificationTypeInfo{ parseVerificationTypeInfo(cf,r)}
	self.offset_delta = util.U2( 64 - self.frame_type)
}

func (self *StackMapFrame) parseSameLocals1StackItemFrameExtended(cf ClassFile,r *util.BigEndianReader) {
	self.offset_delta = r.ReadU2()
	self.stack = []VerificationTypeInfo{parseVerificationTypeInfo(cf,r)}
}

func (self *StackMapFrame) parseChopFrame(cf ClassFile,r *util.BigEndianReader) {
	self.offset_delta = r.ReadU2()
}


func (self *StackMapFrame) parseSameFrameExtended(cf ClassFile,r *util.BigEndianReader) {
	self.offset_delta = r.ReadU2()
}

func (self *StackMapFrame) parseAppendFrame(cf ClassFile,r *util.BigEndianReader) {
	self.offset_delta = r.ReadU2()
	self.locals = make([]VerificationTypeInfo,self.frame_type - 251)
	for i := range self.locals {
		self.locals[i] = parseVerificationTypeInfo(cf,r)
	}
	fmt.Printf("self",self)
}


func (self *StackMapFrame) parseFullFrame(cf ClassFile,r *util.BigEndianReader) {
	self.offset_delta = r.ReadU2()
	self.locals = make([]VerificationTypeInfo,self.frame_type - 251)
	for i := range self.locals {
		self.locals[i] = parseVerificationTypeInfo(cf,r)
	}
	self.stack = make([]VerificationTypeInfo,self.frame_type - 251)
	for i := range self.stack {
		self.stack[i] = parseVerificationTypeInfo(cf,r)
	}
}

/**********************************************************
 **   StackMapTable.stack_map_frame[i].stack[i]          **
 **     -> VerificationTypeInfo                          **
 **********************************************************/
/**
Top_variable_info {
    util.U1 tag = ITEM_Top; /* 0
}

Integer_variable_info {
util.U1 tag = ITEM_Integer; /* 1
}

Float_variable_info {
util.U1 tag = ITEM_Float; /* 2
}

Null_variable_info {
util.U1 tag = ITEM_Null; /* 5
}

UninitializedThis_variable_info {
util.U1 tag = ITEM_UninitializedThis; /* 6
}

Object_variable_info {
util.U1 tag = ITEM_Object; /* 7
util.U2 cpool_index;
}

Uninitialized_variable_info {
util.U1 tag = ITEM_Uninitialized; /* 8
util.U2 offset;
}

Long_variable_info {
util.U1 tag = ITEM_Long; /* 4
}

Double_variable_info {
util.U1 tag = ITEM_Double; /* 3
}

*/

type VerificationTypeInfo struct {
	//cp ConstantPool
	tag util.U1
	cpool_index_or_offset util.U2
	//offset util.U2
}

func  parseVerificationTypeInfo(cf ClassFile,r *util.BigEndianReader) VerificationTypeInfo{
	//self.cp = cf.constant_pool
	info := VerificationTypeInfo{}
	info.tag = r.ReadU1()
	if (info.tag == ITEM_Object  || info.tag == ITEM_Uninitialized ){
		info.cpool_index_or_offset = r.ReadU2()
	}
	return info
}

func (self VerificationTypeInfo) Tag() util.U1{
	return self.tag
}

func (self VerificationTypeInfo) getCPoolIndex() util.U2{
	if (self.tag == ITEM_Object) {
		return self.cpool_index_or_offset
	}
	return 0
}

func (self *VerificationTypeInfo) getOffset() util.U2{
	if (self.tag == ITEM_Uninitialized) {
		return self.cpool_index_or_offset
	}
	return 0
}

/*type StackMapFrame interface{
	parse(cf ClassFile,r *util.BigEndianReader)
	FrameType() util.U1
	OffsetDelta() util.U1
}*/
