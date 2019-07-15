package jd.jvmexample.classfile.start;

import java.io.IOException;

/**
 this_calss: 1 // jd.jvmexample.main.classfile.startfromzero.EmptyInterface
 super_class: 2 // java/lang/Object
 Constant pool:
 flags: ACC_PUBLIC, ACC_INTERFACE, ACC_ABSTRACT
 Constant pool:
 #1 = Class              #5              // jd/jvmexample/main/classfile/startfromzero/EmptyInterface
 #2 = Class              #6              // java/lang/Object
 #3 = Utf8               SourceFile
 #4 = Utf8               EmptyInterface.java
 #5 = Utf8               jd/jvmexample/main/classfile/startfromzero/EmptyInterface
 #6 = Utf8               java/lang/Object
 Attributes:
 #0 SourceFile: "EmptyInterface.java"

 */
interface EmptyInterface {
}



/**
 flags: ACC_PUBLIC, ACC_SUPER
 Constant pool:
 #1 = Methodref          #3.#13         // java/lang/Object."<init>":()V
 #2 = Class              #14            // jd/jvmexample/main/classfile/startfromzero/EmptyClassFile
 #3 = Class              #15            // java/lang/Object
 #4 = Utf8               <init>
 #5 = Utf8               ()V
 #6 = Utf8               Code
 #7 = Utf8               LineNumberTable
 #8 = Utf8               LocalVariableTable
 #9 = Utf8               this
 #10 = Utf8               Ljd/jvmexample/main/classfile/startfromzero/EmptyClassFile;
 #11 = Utf8               SourceFile
 #12 = Utf8               EmptyClassFile.java
 #13 = NameAndType        #4:#5          // "<init>":()V
 #14 = Utf8               jd/jvmexample/main/classfile/startfromzero/EmptyClassFile
 #15 = Utf8               java/lang/Object

 Methods:
 #0: <init>
 descriptor: #5 -> ()V
 flags: ACC_PUBLIC
 stack=1, locals=1, args_size=1
 Code:
 0: aload_0
 1: invokespecial #1                  // Method java/lang/Object."<init>":()V
 4: return
 LineNumberTable:
 line 3: 0
 LocalVariableTable:
 Start  Length  Slot  Name   Signature
 0       5     0  this   Ljd/jvmexample/main/classfile/startfromzero/EmptyClassFile;
 }

 Attributes:
 #0 SourceFile: "EmptyClassFile.java"

 */
class EmptyClassFile {
}

enum EmptyEnum{}


interface SingleMethodInterface{
    void fun();
}


class SingleMethodClass {
    void fun(){}
}

class ThrowException{
    void fun() throws IOException,IllegalArgumentException,RuntimeException {}
}

