package jd.jvmexample.classfile.instrution.loadstore;

/**
 * Created by huangxia on 2008/12/4.
 */
public class DemoLoadStoreInstructions {

    /*
    Code:
      stack=1, locals=1, args_size=1
         0: aload_0 // load from local variable table
         1: areturn
      LocalVariableTable:
        Start  Length  Slot  Name   Signature
            0       2     0  this   Ljd/jvmexample/classfile/instrution/loadstore/DemoLoadStoreInstructions;
    */
    public DemoLoadStoreInstructions getThis(){
        return this ;
    }

    /******************************************************
     *                  load constant value
     ******************************************************/

    /*
        Code:
          stack=1, locals=0, args_size=0
             0: iconst_1 // load constant 1
             1: ireturn
    */
    public static int getStaticConstantInt1(){
        return 1 ;
    }

    /**
     Code:
         stack=1, locals=1, args_size=1
             0: iconst_m1
             1: ireturn
         LocalVariableTable:
             Start  Length  Slot  Name   Signature
             0       2     0  this   Ljd/jvmexample/classfile/instrution/loadstore/DemoLoadStoreInstructions;
     */
    public int getConstantIntN1(){
        return -1 ;
    }

    /**
     Constant pool:
         #1 = Methodref          #6.#57         // java/lang/Object."<init>":()V
         #2 = Long               -1l
         #4 = Integer            2147483647
         #5 = Class              #58            // jd/jvmexample/classfile/instrution/loadstore/DemoLoadStoreInstructions
         ...
     Code:
         stack=2, locals=1, args_size=1
         0: ldc2_w        #2                  // long -1l
         3: lreturn
         LocalVariableTable:
         Start  Length  Slot  Name   Signature
         0       4     0  this   Ljd/jvmexample/classfile/instrution/loadstore/DemoLoadStoreInstructions;
     */
    public long getConstantLongN1(){
        return -1 ;
    }


    public int getConstantInt100(){
        /**
         * Code:
             stack=1, locals=1, args_size=1
                 0: bipush        100
                 2: ireturn
             LocalVariableTable:
                 Start  Length  Slot  Name   Signature
                 0       3     0  this   Ljd/jvmexample/classfile/instrution/loadstore/DemoLoadStoreInstructions;
         */
        return 100 ;
    }
    public int getConstantIntMax(){
        /**
         * Code:
             stack=1, locals=1, args_size=1
                 0: ldc           #4                  // int 2147483647
                 2: ireturn
             LocalVariableTable:
                 Start  Length  Slot  Name   Signature
                 0       3     0  this   Ljd/jvmexample/classfile/instrution/loadstore/DemoLoadStoreInstructions;
         */
        return 0x7fffffff ;
    }


    /******************************************************
     *                  load from local variable table
     ******************************************************/

    public static int getStaticInt(int a){
        /**
         *  Code:
             stack=1, locals=1, args_size=1
                 0: iload_0
                 1: ireturn
             LocalVariableTable:
                 Start  Length  Slot  Name   Signature
                 0       2     0     a   I
         */
        return a ;
    }

    public int getInt(int a){
        /**
         *     Code:
             stack=1, locals=2, args_size=2
                 0: iload_1
                 1: ireturn
             LocalVariableTable:
                 Start  Length  Slot  Name   Signature
                 0       2     0  this   Ljd/jvmexample/classfile/instrution/loadstore/DemoLoadStoreInstructions;
                 0       2     1     a   I
         */
        return a ;
    }

    public int getInt2(int a,int b){
        /**
         * Code:
             stack=1, locals=3, args_size=3
                 0: iload_2
                 1: ireturn
             LocalVariableTable:
                 Start  Length  Slot  Name   Signature
                 0       2     0  this   Ljd/jvmexample/classfile/instrution/loadstore/DemoLoadStoreInstructions;
                 0       2     1     a   I
                 0       2     2     b   I
         */
        return b ;
    }

    public int getIntAnother(int a,int b){
        /**
         * Code:
             stack=1, locals=4, args_size=3
                 0: bipush        100
                 2: istore_3
                 3: iload_3
                 4: ireturn
             LocalVariableTable:
                 Start  Length  Slot  Name   Signature
                 0       5     0  this   Ljd/jvmexample/classfile/instrution/loadstore/DemoLoadStoreInstructions;
                 0       5     1     a   I
                 0       5     2     b   I
                 3       2     3     c   I
         */
        int c = 100 ;
        return c ;
    }

    /******************************************************
     *                  store onto local variable table
     ******************************************************/


    public int storeInt2(int a,int b){
        /**
         * Code:
             stack=1, locals=4, args_size=3
                 0: iload_1
                 1: istore_3
                 2: iload_3
                 3: ireturn
             LocalVariableTable:
                 Start  Length  Slot  Name   Signature
                 0       4     0  this   Ljd/jvmexample/classfile/instrution/loadstore/DemoLoadStoreInstructions;
                 0       4     1     a   I
                 0       4     2     b   I
                 2       2     3     c   I
         */
        int c = a ;
        return c ;
    }

    public Object store1(int pi,long pl,float pf,byte pb,boolean pbl,double pd,Object pref,int[] parr){
        /**
         * Code:
             stack=3, locals=21, args_size=9
                 0: iload_1
                 1: istore        11
                 3: lload_2
                 4: lstore        12
                 6: fload         4
                 8: fstore        14
                 10: iload         5
                 12: istore        15
                 14: iload         6
                 16: istore        16
                 18: dload         7
                 20: dstore        17
                 22: aload         9
                 24: astore        19
                 26: aload         10
                 28: astore        20
                 30: aload         10
                 32: iconst_0
                 33: iload_1
                 34: iastore
                 35: aload         9
                 37: areturn
             LocalVariableTable:
                 Start  Length  Slot  Name   Signature
                 0      38     0  this   Ljd/jvmexample/classfile/instrution/loadstore/DemoLoadStoreInstructions;
                 0      38     1    pi   I
                 0      38     2    pl   J
                 0      38     4    pf   F
                 0      38     5    pb   B
                 0      38     6   pbl   Z
                 0      38     7    pd   D
                 0      38     9  pref   Ljava/lang/Object;
                 0      38    10  parr   [I
                 3      35    11     a   I
                 6      32    12     b   J
                 10      28    14     c   F
                 14      24    15     d   B
                 18      20    16     e   Z
                 22      16    17     f   D
                 26      12    19     g   Ljava/lang/Object;
                 30       8    20     h   [I
         */
        int a = pi ;
        long b = pl ;
        float c = pf;
        byte d = pb ;
        boolean e = pbl ;
        double f = pd ;
        Object g = pref;
        int[] h = parr ;
        parr[0] = pi ;

        return pref;
    }


}
