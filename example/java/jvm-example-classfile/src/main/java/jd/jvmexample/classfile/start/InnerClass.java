package jd.jvmexample.classfile.start;

import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;

class OuterClass {

    class InnerClass{}
    static class StaticNestedClass { }
    interface InnerInterface{}
    enum InnerEnum{}
    @interface InnerAnnotation{}

    @Retention(RetentionPolicy.RUNTIME)
    @interface InnerRuntimeAnnotation{}

    @InnerAnnotation
    class ComplexInnerClass extends InnerClass implements  InnerInterface {
        InnerEnum e ;

        @InnerRuntimeAnnotation
        public void fun(){
            class NestedClassInMethod {} // 局部类，无外部类
            Runnable r = new Runnable() { // 匿名类，无外部类、无类名，内部类名有编译器自动命名
                @Override
                public void run() {

                }
            };
        }
    }

}
