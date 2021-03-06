package runtime

import (
	"../util"
)

// NoOperandsInstruction singletons
var (
	nop         = &NOP{}
	/* aconst_null ~  dconst_1 : push constant value into oprand stack*/
	aconst_null = &ACONST_NULL{}
	iconst_m1   = &ICONST_N{-1}
	iconst_0    = &ICONST_N{0}
	iconst_1    = &ICONST_N{1}
	iconst_2    = &ICONST_N{2}
	iconst_3    = &ICONST_N{3}
	iconst_4    = &ICONST_N{4}
	iconst_5    = &ICONST_N{5}
	lconst_0    = &LCONST_N{0}
	lconst_1    = &LCONST_N{1}
	fconst_0    = &FCONST_N{0}
	fconst_1    = &FCONST_N{1}
	fconst_2    = &FCONST_N{2}
	dconst_0    = &DCONST_N{0}
	dconst_1    = &DCONST_N{1}

	bipush = &BIPUSH{}
	sipush = &SIPUSH{}
	ldc = &LDC{} // load from constant pool
	ldc_w = &LDC_W{}
	ldc2_w = &LDC2_W{}

	/* load var from local variable */
	iload = &U4LOAD{}
	lload = &U8LOAD{}
	fload = &U4LOAD{}
	dload = &U8LOAD{}
	aload = &ALOAD{}
	iload_0     = &U4LOAD_N{0}
	iload_1     = &U4LOAD_N{1}
	iload_2     = &U4LOAD_N{2}
	iload_3     = &U4LOAD_N{3}
	lload_0     = &U8LOAD_N{0}
	lload_1     = &U8LOAD_N{1}
	lload_2     = &U8LOAD_N{2}
	lload_3     = &U8LOAD_N{3}
	fload_0     = &U4LOAD_N{0}
	fload_1     = &U4LOAD_N{1}
	fload_2     = &U4LOAD_N{2}
	fload_3     = &U4LOAD_N{3}
	dload_0     = &U8LOAD_N{0}
	dload_1     = &U8LOAD_N{1}
	dload_2     = &U8LOAD_N{2}
	dload_3     = &U8LOAD_N{3}
	aload_0     = &ALOAD_N{0}
	aload_1     = &ALOAD_N{1}
	aload_2     = &ALOAD_N{2}
	aload_3     = &ALOAD_N{3}
	iaload      = &IALOAD{}
	laload      = &LALOAD{}
	faload      = &FALOAD{}
	daload      = &DALOAD{}
	aaload      = &AALOAD{}
	baload      = &BALOAD{}
	caload      = &CALOAD{}
	saload      = &SALOAD{}

	/* store into the local variable */
	istore = &U4STORE{}
	lstore = &U8STORE{}
	fstore = &U4STORE{}
	dstore = &U8STORE{}
	astore = &U4STORE{}
	istore_0    = &U4STORE_N{0}
	istore_1    = &U4STORE_N{1}
	istore_2    = &U4STORE_N{2}
	istore_3    = &U4STORE_N{3}
	lstore_0    = &U8STORE_N{0}
	lstore_1    = &U8STORE_N{1}
	lstore_2    = &U8STORE_N{2}
	lstore_3    = &U8STORE_N{3}
	fstore_0    = &U4STORE_N{0}
	fstore_1    = &U4STORE_N{1}
	fstore_2    = &U4STORE_N{2}
	fstore_3    = &U4STORE_N{3}
	dstore_0    = &U8STORE_N{0}
	dstore_1    = &U8STORE_N{1}
	dstore_2    = &U8STORE_N{2}
	dstore_3    = &U8STORE_N{3}
	astore_0    = &U4STORE_N{0}
	astore_1    = &U4STORE_N{1}
	astore_2    = &U4STORE_N{2}
	astore_3    = &U4STORE_N{3}
	iastore     = &IASTORE{}
	lastore     = &LASTORE{}
	fastore     = &FASTORE{}
	dastore     = &DASTORE{}
	aastore     = &AASTORE{}
	bastore     = &BASTORE{}
	castore     = &CASTORE{}
	sastore     = &SASTORE{}

	/* stack instructions */
	pop         = &POP{}
	pop2        = &POP2{}
	dup         = &DUP{}
	dup_x1      = &DUP_X1{}
	dup_x2      = &DUP_X2{}
	dup2        = &DUP2{}
	dup2_x1     = &DUP2_X1{}
	dup2_x2     = &DUP2_X2{}
	swap        = &SWAP{}

	/* MATH */
	iadd        = &IADD{}
	ladd        = &LADD{}
	fadd        = &FADD{}
	dadd        = &DADD{}
	isub        = &ISUB{}
	lsub        = &LSUB{}
	fsub        = &FSUB{}
	dsub        = &DSUB{}
	imul        = &IMUL{}
	lmul        = &LMUL{}
	fmul        = &FMUL{}
	dmul        = &DMUL{}
	idiv        = &IDIV{}
	ldiv        = &LDIV{}
	fdiv        = &FDIV{}
	ddiv        = &DDIV{}
	irem        = &IREM{}
	lrem        = &LREM{}
	frem        = &FREM{}
	drem        = &DREM{}
	ineg        = &INEG{}
	lneg        = &LNEG{}
	fneg        = &FNEG{}
	dneg        = &DNEG{}
	ishl        = &ISHL{}
	lshl        = &LSHL{}
	ishr        = &ISHR{}
	lshr        = &LSHR{}
	iushr       = &IUSHR{}
	lushr       = &LUSHR{}
	iand        = &IAND{}
	land        = &LAND{}
	ior         = &IOR{}
	lor         = &LOR{}
	ixor        = &IXOR{}
	lxor        = &LXOR{}
	iinc 		= &IINC{}

	/* Type Conversion Instructions */
	i2l         = &I2L{}
	i2f         = &I2F{}
	i2d         = &I2D{}
	l2i         = &L2I{}
	l2f         = &L2F{}
	l2d         = &L2D{}
	f2i         = &F2I{}
	f2l         = &F2L{}
	f2d         = &F2D{}
	d2i         = &D2I{}
	d2l         = &D2L{}
	d2f         = &D2F{}
	i2b         = &I2B{}
	i2c         = &I2C{}
	i2s         = &I2S{}

  /* Comparisons */
	lcmp        = &LCMP{}
	fcmpl       = &FCMP_OP{-1}
	fcmpg       = &FCMP_OP{1}
	dcmpl       = &DCMP_OP{-1}
	dcmpg       = &DCMP_OP{1}
	ifeq		= &IF_COND{&eq{}}
	ifne		= &IF_COND{&ne{}}
	iflt		= &IF_COND{&lt{}}
	ifle		= &IF_COND{&le{}}
	ifgt		= &IF_COND{&gt{}}
	ifge		= &IF_COND{&ge{}}
	if_icmpeq		= &IF_ICMP_COMD{&eq{}}
	if_icmpne		= &IF_ICMP_COMD{&ne{}}
	if_icmplt		= &IF_ICMP_COMD{&lt{}}
	if_icmple		= &IF_ICMP_COMD{&le{}}
	if_icmpgt		= &IF_ICMP_COMD{&gt{}}
	if_icmpge		= &IF_ICMP_COMD{&ge{}}
	if_acmpeq		= &IF_ACMP_COMD{true}
	if_acmpne		= &IF_ACMP_COMD{false}
	// extend instructions
	ifnull 			= &IFNULL{true} // 198 (0xc6)
	ifnonnull 		= &IFNULL{false} // 199 (0xc7)

  /* References Control */
  _goto        = &GOTO{}
  jsr         = &JSR{}
  // extend 
  goto_w        = &GOTO_W{}
  jsr_w         = &JSR_W{}
  ret         = &RET{}
  tableswitch = &TABLE_SWITCH{}
  lookupswitch= &LOOKUP_SWITCH{}
  ireturn     = &IRETURN{}
  lreturn     = &LRETURN{}
  freturn     = &FRETURN{}
  dreturn     = &DRETURN{}
  areturn     = &ARETURN{}
  _return      = &RETURN{}

	//arraylength = &ARRAY_LENGTH{}
	//// athrow        = &ATHROW{}
	//// monitorenter  = &MONITOR_ENTER{}
	//// monitorexit   = &MONITOR_EXIT{}
	//invoke_native = &INVOKE_NATIVE{}
	INSTR_CDOE_WIDE = util.U1(0xc4)
)

type Instruction interface {
	execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)
}
func NewInstructionEngine(opcode byte) InstructionEngine {
	insts := make([]Instruction,256)
	engine := InstructionEngine{insts}

	/* Constants */
	insts[0x00] = nop
	insts[0x01] = aconst_null
	insts[0x02] = iconst_m1
	insts[0x03] = iconst_0
	insts[0x04] = iconst_1
	insts[0x05] = iconst_2
	insts[0x06] = iconst_3
	insts[0x07] = iconst_4
	insts[0x08] = iconst_5
	insts[0x09] = lconst_0
	insts[0x0a] = lconst_1
	insts[0x0b] = fconst_0
	insts[0x0c] = fconst_1
	insts[0x0d] = fconst_2
	insts[0x0e] = dconst_0
	insts[0x0f] = dconst_1
	insts[0x10] = bipush
	insts[0x11] = sipush
	insts[0x12] = ldc
	insts[0x13] = ldc_w
	insts[0x14] = ldc2_w

	/* Loads */
	insts[0x15] = iload
	insts[0x16] = lload
	insts[0x17] = fload
	insts[0x18] = dload
	insts[0x19] = aload
	insts[0x1a] = iload_0
	insts[0x1b] = iload_1
	insts[0x1c] = iload_2
	insts[0x1d] = iload_3
	insts[0x1e] = lload_0
	insts[0x1f] = lload_1
	insts[0x20] = lload_2
	insts[0x21] = lload_3
	insts[0x22] = fload_0
	insts[0x23] = fload_1
	insts[0x24] = fload_2
	insts[0x25] = fload_3
	insts[0x26] = dload_0
	insts[0x27] = dload_1
	insts[0x28] = dload_2
	insts[0x29] = dload_3
	insts[0x2a] = aload_0
	insts[0x2b] = aload_1
	insts[0x2c] = aload_2
	insts[0x2d] = aload_3
	insts[0x2e] = iaload
	insts[0x2f] = laload
	insts[0x30] = faload
	insts[0x31] = daload
	insts[0x32] = aaload
	insts[0x33] = baload
	insts[0x34] = caload
	insts[0x35] = saload

	/* Stores */
	insts[0x36] = istore
	insts[0x37] = lstore
	insts[0x38] = fstore
	insts[0x39] = dstore
	insts[0x3a] = astore
	insts[0x3b] = istore_0
	insts[0x3c] = istore_1
	insts[0x3d] = istore_2
	insts[0x3e] = istore_3
	insts[0x3f] = lstore_0
	insts[0x40] = lstore_1
	insts[0x41] = lstore_2
	insts[0x42] = lstore_3
	insts[0x43] = fstore_0
	insts[0x44] = fstore_1
	insts[0x45] = fstore_2
	insts[0x46] = fstore_3
	insts[0x47] = dstore_0
	insts[0x48] = dstore_1
	insts[0x49] = dstore_2
	insts[0x4a] = dstore_3
	insts[0x4b] = astore_0
	insts[0x4c] = astore_1
	insts[0x4d] = astore_2
	insts[0x4e] = astore_3
	insts[0x4f] = iastore
	insts[0x50] = lastore
	insts[0x51] = fastore
	insts[0x52] = dastore
	insts[0x53] = aastore
	insts[0x54] = bastore
	insts[0x55] = castore
	insts[0x56] = sastore

	/* Stack */
	insts[0x57] = pop
	insts[0x58] = pop2
	insts[0x59] = dup
	insts[0x5a] = dup_x1
	insts[0x5b] = dup_x2
	insts[0x5c] = dup2
	insts[0x5d] = dup2_x1
	insts[0x5e] = dup2_x2
	insts[0x5f] = swap

	/* Math */
	insts[0x60] = iadd
	insts[0x61] = ladd
	insts[0x62] = fadd
	insts[0x63] = dadd
	insts[0x64] = isub
	insts[0x65] = lsub
	insts[0x66] = fsub
	insts[0x67] = dsub
	insts[0x68] = imul
	insts[0x69] = lmul
	insts[0x6a] = fmul
	insts[0x6b] = dmul
	insts[0x6c] = idiv
	insts[0x6d] = ldiv
	insts[0x6e] = fdiv
	insts[0x6f] = ddiv
	insts[0x70] = irem
	insts[0x71] = lrem
	insts[0x72] = frem
	insts[0x73] = drem
	insts[0x74] = ineg
	insts[0x75] = lneg
	insts[0x76] = fneg
	insts[0x77] = dneg
	insts[0x78] = ishl
	insts[0x79] = lshl
	insts[0x7a] = ishr
	insts[0x7b] = lshr
	insts[0x7c] = iushr
	insts[0x7d] = lushr
	insts[0x7e] = iand
	insts[0x7f] = land
	insts[0x80] = ior
	insts[0x81] = lor
	insts[0x82] = ixor
	insts[0x83] = lxor
	insts[0x84] = iinc

	/* Type Conversion Instructions */
	insts[0x85] = i2l
	insts[0x86] = i2f
	insts[0x87] = i2d
	insts[0x88] = l2i
	insts[0x89] = l2f
	insts[0x8a] = l2d
	insts[0x8b] = f2i
	insts[0x8c] = f2l
	insts[0x8d] = f2d
	insts[0x8e] = d2i
	insts[0x8f] = d2l
	insts[0x90] = d2f
	insts[0x91] = i2b
	insts[0x92] = i2c
	insts[0x93] = i2s

  /* Comparisons */
	insts[0x94] = lcmp
	insts[0x95] = fcmpl
	insts[0x96] = fcmpg
	insts[0x97] = dcmpl
	insts[0x98] = dcmpg
	insts[0x99] = ifeq
	insts[0x9a] = ifne
	insts[0x9b] = iflt
	insts[0x9c] = ifge
	insts[0x9d] = ifgt
	insts[0x9e] = ifle
	insts[0x9f] = if_icmpeq
	insts[0xa0] = if_icmpne
	insts[0xa1] = if_icmplt
	insts[0xa2] = if_icmpge
	insts[0xa3] = if_icmpgt
	insts[0xa4] = if_icmple
	insts[0xa5] = if_acmpeq
	insts[0xa6] = if_acmpne

  /* References Control */
	insts[0xa7] = _goto
	insts[0xa8] = jsr 
	insts[0xa9] = ret 
	insts[0xaa] = tableswitch
	insts[0xab] = lookupswitch
	insts[0xac] = ireturn
	insts[0xad] = lreturn
	insts[0xae] = freturn
	insts[0xaf] = dreturn
	insts[0xb0] = areturn
	insts[0xb1] = _return
  
	//insts[0xb2] = &GET_STATIC{}
	//insts[0xb3] = &PUT_STATIC{}
	//insts[0xb4] = &GET_FIELD{}
	//insts[0xb5] = &PUT_FIELD{}
	//insts[0xb6] = &INVOKE_VIRTUAL{}
	//insts[0xb7] = &INVOKE_SPECIAL{}
	//insts[0xb8] = &INVOKE_STATIC{}
	//insts[0xb9] = &INVOKE_INTERFACE{}
	//// insts[0xba] = &INVOKE_DYNAMIC{}
	//insts[0xbb] = &NEW{}
	//insts[0xbc] = &NEW_ARRAY{}
	//insts[0xbd] = &ANEW_ARRAY{}
	//insts[0xbe] = arraylength
	//// insts[0xbf] = athrow
	//insts[0xc0] = &CHECK_CAST{}
	//insts[0xc1] = &INSTANCE_OF{}
	//// insts[0xc2]=  monitorenter
	//// insts[0xc3] = return monitorexit

	/* Extended */
	//insts[0xc4] = &WIDE{}
	//insts[0xc5] = &MULTI_ANEW_ARRAY{}
	insts[0xc6] = ifnull
	insts[0xc7] = ifnonnull
	insts[0xc8] = goto_w
	insts[0xc9] = jsr_w
	// insts[0xca: breakpoint
	// insts[0xfe] = invoke_native
	return engine
}


type InstructionEngine struct {
	instructions []Instruction
}

type NOP struct {
}

func (self *NOP) execute(reader *util.BigEndianReader,frame *StackFrame,wideMode bool)   {
}

func (self *InstructionEngine) execute(thread Thread){
	frame := thread.stack.Top()
	if( len(frame.method.Code) > 0){
		reader := util.NewBigEndianReader(frame.method.Code)
		// Extend local variable index by additional bytes
		// where <opcode> is one of iload, fload, aload, lload, dload, istore,
		//   fstore, astore, lstore, dstore,  ret  or iinc
		wideMode := false
		opcode := reader.ReadU1()
		for  ; reader.HasNext() ;   {
			opcode = reader.ReadU1()
			wideMode = INSTR_CDOE_WIDE == opcode
			if( wideMode){
				opcode = reader.ReadU1()
			}
			self.instructions[opcode].execute(reader,thread.stack.top,wideMode)
			if(thread.stack.top.errorCode > 0) {
				// exit current method
				break
			}
		}
	}

	thread.stack.Pop()
}


func index(reader *util.BigEndianReader,wideMode bool ) util.U2{
	var index util.U2
	if wideMode {
		index = reader.ReadU2()
	}else {
		index = util.U2(reader.ReadU1())
	}
	return index
}