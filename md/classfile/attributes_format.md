
# 4.7. [Attributes](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7)

#### Table 4.7-A. Predefined class file attributes (by section)
Attribute|Section|class file|Java SE
---|---|---|---
ConstantValue|§4.7.2|45.3|1.0.2
Code|§4.7.3|45.3|1.0.2
StackMapTable|§4.7.4|50.0|6
Exceptions|§4.7.5|45.3|1.0.2
InnerClasses|§4.7.6|45.3|1.1
EnclosingMethod|§4.7.7|49.0|5.0
Synthetic|§4.7.8|45.3|1.1
Signature|§4.7.9|49.0|5.0
SourceFile|§4.7.10|45.3|1.0.2
SourceDebugExtension|§4.7.11|49.0|5.0
LineNumberTable|§4.7.12|45.3|1.0.2
LocalVariableTable|§4.7.13|45.3|1.0.2
LocalVariableTypeTable|§4.7.14|49.0|5.0
Deprecated|§4.7.15|45.3|1.1
RuntimeVisibleAnnotations|§4.7.16|49.0|5.0
RuntimeInvisibleAnnotations|§4.7.17|49.0|5.0
RuntimeVisibleParameterAnnotations|§4.7.18|49.0|5.0
RuntimeInvisibleParameterAnnotations|§4.7.19|49.0|5.0
RuntimeVisibleTypeAnnotations|§4.7.20|52.0|8
RuntimeInvisibleTypeAnnotations|§4.7.21|52.0|8
AnnotationDefault|§4.7.22|49.0|5.0
BootstrapMethods|§4.7.23|51.0|7
MethodParameters|§4.7.24|52.0|8
Module|§4.7.25|53.0|9
ModulePackages|§4.7.26|53.0|9
ModuleMainClass|§4.7.27|53.0|9
NestHost|§4.7.28|55.0|11
NestMembers|§4.7.29|55.0|11

## 4.7.2 [ConstantValue](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.2) Attribute
```text


ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
```
## 4.7.3 [Code](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.3) Attribute
```text


Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
```
## 4.7.4 [StackMapTable](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.4) Attribute
```text


StackMapTable_attribute {
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
    u1 tag = ITEM_Top; /* 0 */
}

Integer_variable_info {
    u1 tag = ITEM_Integer; /* 1 */
}

Float_variable_info {
    u1 tag = ITEM_Float; /* 2 */
}

Null_variable_info {
    u1 tag = ITEM_Null; /* 5 */
}

UninitializedThis_variable_info {
    u1 tag = ITEM_UninitializedThis; /* 6 */
}

Object_variable_info {
    u1 tag = ITEM_Object; /* 7 */
    u2 cpool_index;
}

Uninitialized_variable_info {
    u1 tag = ITEM_Uninitialized; /* 8 */
    u2 offset;
}

Long_variable_info {
    u1 tag = ITEM_Long; /* 4 */
}

Double_variable_info {
    u1 tag = ITEM_Double; /* 3 */
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
    u1 frame_type = SAME; /* 0-63 */
}

same_locals_1_stack_item_frame {
    u1 frame_type = SAME_LOCALS_1_STACK_ITEM; /* 64-127 */
    verification_type_info stack[1];
}

same_locals_1_stack_item_frame_extended {
    u1 frame_type = SAME_LOCALS_1_STACK_ITEM_EXTENDED; /* 247 */
    u2 offset_delta;
    verification_type_info stack[1];
}

chop_frame {
    u1 frame_type = CHOP; /* 248-250 */
    u2 offset_delta;
}

same_frame_extended {
    u1 frame_type = SAME_FRAME_EXTENDED; /* 251 */
    u2 offset_delta;
}

append_frame {
    u1 frame_type = APPEND; /* 252-254 */
    u2 offset_delta;
    verification_type_info locals[frame_type - 251];
}

full_frame {
    u1 frame_type = FULL_FRAME; /* 255 */
    u2 offset_delta;
    u2 number_of_locals;
    verification_type_info locals[number_of_locals];
    u2 number_of_stack_items;
    verification_type_info stack[number_of_stack_items];
}
```
## 4.7.5 [Exceptions](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.5) Attribute
```text


Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
```
## 4.7.6 [InnerClasses](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.6) Attribute
example: OuterClass.InnerClass
```java
class OuterClass {

    public class InnerClass{}
    static class StaticNestedClass { }
    interface InnerInterface{}
    enum InnerEnum{}
    @interface InnerAnnotation{}
}

```

```text

InnerClasses_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_classes;
    {   u2 inner_class_info_index; // 指向CONSTANT_Class_info，匿名类为0 --> OuterClass$InnerClass  
        u2 outer_class_info_index; // 指向CONSTANT_Class_info，局部类、匿名类为0 --> OuterClass, 
        u2 inner_name_index; // --> InnerClass
        u2 inner_class_access_flags;  // ACC_PUBLIC
    } classes[number_of_classes];
}
```


## 4.7.7 [EnclosingMethod](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.7) Attribute
```text


EnclosingMethod_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 class_index;
    u2 method_index;
}
```
## 4.7.8 [Synthetic](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.8) Attribute
```text


Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
```
## 4.7.9 [Signature](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.9) Attribute
```text


Signature_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 signature_index;
}
```
## 4.7.10 [SourceFile](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.10) Attribute
```text


SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
```
## 4.7.11 [SourceDebugExtension](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.11) Attribute
```text


SourceDebugExtension_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 debug_extension[attribute_length];
}
```
## 4.7.12 [LineNumberTable](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.12) Attribute
```text


LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;	
    } line_number_table[line_number_table_length];
}
```
## 4.7.13 [LocalVariableTable](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.13) Attribute
```text


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
## 4.7.14 [LocalVariableTypeTable](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.14) Attribute
```text


LocalVariableTypeTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_type_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 signature_index;
        u2 index;
    } local_variable_type_table[local_variable_type_table_length];
}
```
## 4.7.15 [Deprecated](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.15) Attribute
```text


Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
```
## 4.7.16 [RuntimeVisibleAnnotations](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.16) Attribute
```text


RuntimeVisibleAnnotations_attribute {
    u2         attribute_name_index;
    u4         attribute_length;
    u2         num_annotations;
    annotation annotations[num_annotations];
}

annotation {
    u2 type_index;
    u2 num_element_value_pairs;
    {   u2            element_name_index;
        element_value value;
    } element_value_pairs[num_element_value_pairs];
}

element_value {
    u1 tag;
    union {
        u2 const_value_index;

        {   u2 type_name_index;
            u2 const_name_index;
        } enum_const_value;

        u2 class_info_index;

        annotation annotation_value;

        {   u2            num_values;
            element_value values[num_values];
        } array_value;
    } value;
}
```
## 4.7.17 [RuntimeInvisibleAnnotations](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.17) Attribute
```text


RuntimeInvisibleAnnotations_attribute {
    u2         attribute_name_index;
    u4         attribute_length;
    u2         num_annotations;
    annotation annotations[num_annotations];
}
```
example
```java
@interface Copyright {
    String value();
}

@Copyright("2002 Yoyodyne Propulsion Systems, Inc.")
class A{}
```

## 4.7.18 [RuntimeVisibleParameterAnnotations](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.18) Attribute

RuntimeVisibleParameterAnnotations与RuntimeInvisibleAnnotations的区别在于运行时可见，Annation使用@Retention(RetentionPolicy.RUNTIME)修饰

```text
RuntimeVisibleParameterAnnotations_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 num_parameters;
    {   u2         num_annotations;
        annotation annotations[num_annotations];
    } parameter_annotations[num_parameters];
}
```

example:
```java
    @Retention(RetentionPolicy.RUNTIME)
    @Target(ElementType.METHOD)
    public @interface MethodInfo {
        String name() default "long";
    }

    @MethodInfo
    public void fun(){}
```


## 4.7.19 [RuntimeInvisibleParameterAnnotations](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.19) Attribute
```text


RuntimeInvisibleParameterAnnotations_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 num_parameters;
    {   u2         num_annotations;
        annotation annotations[num_annotations];
    } parameter_annotations[num_parameters];
}
```
## 4.7.20 [RuntimeVisibleTypeAnnotations](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.20) Attribute
```text


RuntimeVisibleTypeAnnotations_attribute {
    u2              attribute_name_index;
    u4              attribute_length;
    u2              num_annotations;
    type_annotation annotations[num_annotations];
}

type_annotation {
    u1 target_type;
    union {
        type_parameter_target;
        supertype_target;
        type_parameter_bound_target;
        empty_target;
        formal_parameter_target;
        throws_target;
        localvar_target;
        catch_target;
        offset_target;
        type_argument_target;
    } target_info;
    type_path target_path;
    u2        type_index;
    u2        num_element_value_pairs;
    {   u2            element_name_index;
        element_value value;
    } element_value_pairs[num_element_value_pairs];
}

type_parameter_target {
    u1 type_parameter_index;
}

supertype_target {
    u2 supertype_index;
}

type_parameter_bound_target {
    u1 type_parameter_index;
    u1 bound_index;
}

empty_target {
}

formal_parameter_target {
    u1 formal_parameter_index;
}

throws_target {
    u2 throws_type_index;
}

localvar_target {
    u2 table_length;
    {   u2 start_pc;
        u2 length;
        u2 index;
    } table[table_length];
}

catch_target {
    u2 exception_table_index;
}

offset_target {
    u2 offset;
}

type_argument_target {
    u2 offset;
    u1 type_argument_index;
}

@Foo String[][]   // Annotates the class type String
String @Foo [][]  // Annotates the array type String[][]
String[] @Foo []  // Annotates the array type String[]

@Foo Outer.Middle.Inner
Outer.@Foo Middle.Inner
Outer.Middle.@Foo Inner

@Foo Map&lt;String,Object&gt;
Map&lt;@Foo String,Object&gt;
Map&lt;String,@Foo Object&gt;

List&lt;@Foo ? extends String&gt;
List&lt;? extends @Foo String&gt;

type_path {
    u1 path_length;
    {   u1 type_path_kind;
        u1 type_argument_index;
    } path[path_length];
}

class Outer {
  class Middle {
    class Inner {}
  }
}

class Outer {
  static class MiddleStatic {
    class Inner {}
  }
}

class Outer {
  static class MiddleStatic {
    static class InnerStatic {}
  }
}

class Outer {
  class Middle&lt;T&gt; {
    class Inner&lt;U&gt; {}
  }
}
```
## 4.7.21 [RuntimeInvisibleTypeAnnotations](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.21) Attribute
```text


RuntimeInvisibleTypeAnnotations_attribute {
    u2              attribute_name_index;
    u4              attribute_length;
    u2              num_annotations;
    type_annotation annotations[num_annotations];
}
```
## 4.7.22 [AnnotationDefault](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.22) Attribute
```text


AnnotationDefault_attribute {
    u2            attribute_name_index;
    u4            attribute_length;
    element_value default_value;
}
```
## 4.7.23 [BootstrapMethods](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.23) Attribute
```text


BootstrapMethods_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 num_bootstrap_methods;
    {   u2 bootstrap_method_ref;
        u2 num_bootstrap_arguments;
        u2 bootstrap_arguments[num_bootstrap_arguments];
    } bootstrap_methods[num_bootstrap_methods];
}
```
## 4.7.24 [MethodParameters](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.24) Attribute
```text


MethodParameters_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 parameters_count;
    {   u2 name_index;
        u2 access_flags;
    } parameters[parameters_count];
}
```
## 4.7.25 [Module](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.25) Attribute
```text


Module_attribute {
    u2 attribute_name_index;
    u4 attribute_length;

    u2 module_name_index;
    u2 module_flags;
    u2 module_version_index;

    u2 requires_count;
    {   u2 requires_index;
        u2 requires_flags;
        u2 requires_version_index;
    } requires[requires_count];

    u2 exports_count;
    {   u2 exports_index;
        u2 exports_flags;
        u2 exports_to_count;
        u2 exports_to_index[exports_to_count];
    } exports[exports_count];

    u2 opens_count;
    {   u2 opens_index;
        u2 opens_flags;
        u2 opens_to_count;
        u2 opens_to_index[opens_to_count];
    } opens[opens_count];

    u2 uses_count;
    u2 uses_index[uses_count];

    u2 provides_count;
    {   u2 provides_index;
        u2 provides_with_count;
        u2 provides_with_index[provides_with_count];
    } provides[provides_count];
}
```
## 4.7.26 [ModulePackages](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.26) Attribute
```text


ModulePackages_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 package_count;
    u2 package_index[package_count];
}
```
## 4.7.27 [ModuleMainClass](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.27) Attribute
```text


ModuleMainClass_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 main_class_index;
}
```
## 4.7.28 [NestHost](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.28) Attribute
```text


NestHost_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 host_class_index;
}
```
## 4.7.29 [NestMembers](https://docs.oracle.com/javase/specs/jvms/se12/html/jvms-4.html#jvms-4.7.29) Attribute
```text


NestMembers_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_classes;
    u2 classes[number_of_classes];
}
```

