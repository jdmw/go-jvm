package jd.jvmexample.classfile.attribute.annotation;


import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;



public class RuntimeAnnotation {
    /**
     * 适用类、接口(包括注解类型)或枚举
     */
    @Retention(RetentionPolicy.RUNTIME)
    @Target(ElementType.TYPE)
    public @interface ClassInfo {
        String value();
    }

    /**
     * 适用变量，也包括枚举常量
     */
    @Retention(RetentionPolicy.RUNTIME)
    @Target(ElementType.FIELD)
    public @interface FieldInfo {
        int[] value();
    }

    /**
     * 适用方法
     */
    @Retention(RetentionPolicy.RUNTIME)
    @Target(ElementType.METHOD)
    public @interface MethodInfo {
        String name() default "long";
    }

    @MethodInfo
    public void fun(){}
}