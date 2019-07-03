

# The ClassFile Structure

https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.4
```
ClassFile {
    u4             magic; // 0xCAFEBABE
    u2             minor_version; 
    u2             major_version;
    
    u2             constant_pool_count; // 常量池中常量个数，从1开始计
    cp_info        constant_pool[constant_pool_count-1]; // 常量池
    
    u2             access_flags; 
    u2             this_class; // 当前类名常量定义所在常量池位置，指向CONSTANT_Class_info结构
    u2             super_class;// 父类名常量定义,也是常量池索引。为0表示父类为Object
    u2             interfaces_count; // 
    u2             interfaces[interfaces_count];
    
    u2             fields_count;
    field_info     fields[fields_count];
    
    u2             methods_count;
    method_info    methods[methods_count];
    
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
```


 Class access and property modifiers

Flag Name	|Value|	Interpretation
----|---|---
ACC_PUBLIC			| 0x0001	| Declared public; may be accessed from outside its package.
ACC_FINAL			  | 0x0010	| Declared final; no subclasses allowed.
ACC_SUPER			  | 0x0020	| Treat superclass methods specially when invoked by the invokespecial instruction.
ACC_INTERFACE		| 0x0200	| Is an interface, not a class.
ACC_ABSTRACT		| 0x0400	| Declared abstract; must not be instantiated.
ACC_SYNTHETIC		| 0x1000	| Declared synthetic; not present in the source code.
ACC_ANNOTATION  | 0x2000	| Declared as an annotation type.
ACC_ENUM			  | 0x4000	| Declared as an enum type.
ACC_MODULE			| 0x8000	| Is a module, not a class or interface.


# The Constant Pool


    cp_info {
        u1 tag;
        u1 info[];
    }

Table 4.4-A. Constant pool tags (by section)

Constant Kind	| Tag	
---|---
CONSTANT_Class	            | 7	
CONSTANT_Fieldref	          | 9	
CONSTANT_Methodref	        | 10
CONSTANT_InterfaceMethodref	| 11
CONSTANT_String	        | 8	
CONSTANT_Integer	      | 3	
CONSTANT_Float	        | 4	
CONSTANT_Long	          | 5	
CONSTANT_Double	        | 6	
CONSTANT_NameAndType	 | 12	
CONSTANT_Utf8	         |   1	
CONSTANT_MethodHandle	 | 15	
CONSTANT_MethodType	   | 16	
CONSTANT_Dynamic	     | 17	
CONSTANT_InvokeDynamic | 	18
CONSTANT_Module	       | 19	
CONSTANT_Package	     | 20	

## The CONSTANT_Class_info Structure

表示类和接口，ClassFile结构的this_class、super_class、interface指向此接口。

       CONSTANT_Class_info {
           u1 tag;
           u2 name_index; // 指向CONSTANT_Utf8_info 常量
       }
       

## The CONSTANT_Fieldref_info, CONSTANT_Methodref_info, and CONSTANT_InterfaceMethodref_info Structures


    CONSTANT_Fieldref_info {
        u1 tag;
        u2 class_index;
        u2 name_and_type_index;
    }
    
    CONSTANT_Methodref_info {
        u1 tag;
        u2 class_index;
        u2 name_and_type_index;
    }
    
    CONSTANT_InterfaceMethodref_info {
        u1 tag;
        u2 class_index;
        u2 name_and_type_index;
    }
    
## The CONSTANT_String_info Structure
    
represent constant objects of the type String:
    
    CONSTANT_String_info {
        u1 tag;
        u2 string_index;
    }
    
## CONSTANT_Integer_info and CONSTANT_Float_info Structures
   
represent 4-byte numeric (int and float) constants: (big-endian) 
   
 
      CONSTANT_Integer_info {
        u1 tag;
        u4 bytes;
      }
      
      CONSTANT_Float_info {
        u1 tag;
        u4 bytes;
      }   

   
## CONSTANT_Long_info and CONSTANT_Double_info Structures

represent 8-byte numeric (long and double) constants:


    CONSTANT_Long_info {
        u1 tag;
        u4 high_bytes;
        u4 low_bytes;
    }
    
    CONSTANT_Double_info {
        u1 tag;
        u4 high_bytes;
        u4 low_bytes;
    }
    
## CONSTANT_NameAndType_info Structure
represent a field or method, without indicating which class or interface type it belongs to:
   
       
       CONSTANT_NameAndType_info {
           u1 tag;
           u2 name_index; // --> CONSTANT_Utf8_info 
           u2 descriptor_index; // --> CONSTANT_Utf8_info 
       }
       
## CONSTANT_MethodHandle_info 

 represent a method handle:
 
 
 
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind; //  range 1 to 9
    u2 reference_index;
}


##  CONSTANT_MethodType_info Structure
   
   The CONSTANT_MethodType_info structure is used to represent a method type:
   
       CONSTANT_MethodType_info {
           u1 tag;
           u2 descriptor_index;
       }
       
##  CONSTANT_Dynamic_info and CONSTANT_InvokeDynamic_info Structures

    
    CONSTANT_Dynamic_info {
        u1 tag;
        u2 bootstrap_method_attr_index; // index of the bootstrap_methods array of the bootstrap method table
        u2 name_and_type_index; // --> CONSTANT_NameAndType_info 
    }
    
    CONSTANT_InvokeDynamic_info {
        u1 tag;
        u2 bootstrap_method_attr_index;
        u2 name_and_type_index;
    }
    

## The CONSTANT_Module_info Structure

represent a module:

    CONSTANT_Module_info {
        u1 tag;
        u2 name_index; // CONSTANT_Utf8_info 
    }

## CONSTANT_Package_info        


    CONSTANT_Package_info {
        u1 tag;
        u2 name_index; // -->CONSTANT_Utf8_info
    }
    
# Fields
   
   
       field_info {
           u2             access_flags;
           u2             name_index; // 字段名称--> CONSTANT_Utf8_info
           u2             descriptor_index;// 数据类型 --> CONSTANT_Utf8_info
           u2             attributes_count;
           attribute_info attributes[attributes_count];
       }   
       
Table 4.5-A. Field access and property flags

Flag Name	| Value
---|---
ACC_PUBLIC	|0x0001	
ACC_PRIVATE	|0x0002	
ACC_PROTECTED|	0x0004
ACC_STATIC	|0x0008   
ACC_FINAL	|0x0010	    
ACC_VOLATILE	|0x0040	
ACC_TRANSIENT	|0x0080	
ACC_SYNTHETIC	|0x1000 
ACC_ENUM	|0x4000	       


```java
java:
  public int field = 2;
  
  
compile：
 access_flags : ACC_PUBLIC
 name_index; #11 -> field
 descriptor_index; #26 -> I
 attributes_account: 0
```

# Methods

    method_info {
        u2             access_flags;
        u2             name_index; // 方法名称，指向CONSTANT_Utf8_info类型常量索引
        u2             descriptor_index; //  -->CONSTANT_Utf8_info： (输入变量类型，“;”隔开)返回类型
        u2             attributes_count; 
        attribute_info attributes[attributes_count];
    }


Method access and property flags

Flag Name|	Value
---|---
ACC_PUBLIC|	0x0001	
ACC_PRIVATE	|0x0002	
ACC_PROTECTED|	0x0004
ACC_STATIC	|0x0008	
ACC_FINAL	|0x0010	
ACC_SYNCHRONIZED|0x0020
ACC_BRIDGE	|0x0040	
ACC_VARARGS	|0x0080	
ACC_NATIVE	|0x0100	
ACC_ABSTRACT|0x0400	
ACC_STRICT	|0x0800
ACC_SYNTHETIC	|0x1000	

例子：

```java
java：
public void fun(String a,int b){}

 
compile：
 access_flags : ACC_PUBLIC
 name_index; #25 -> func
 descriptor_index; #26 -> (Ljava/lang/String;I)V
 attributes_account:1
 attributes: [#0(Code)]
```

# Attributes
  
Attributes are used in the ClassFile, field_info, method_info, and Code_attribute structures of the class file format (§4.1, §4.5, §4.6, §4.7.3).
  
      attribute_info {
          u2 attribute_name_index;
          u4 attribute_length;
          u1 info[attribute_length];
      }

```

LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;	
    } line_number_table[line_number_table_length];
}


LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}

```
