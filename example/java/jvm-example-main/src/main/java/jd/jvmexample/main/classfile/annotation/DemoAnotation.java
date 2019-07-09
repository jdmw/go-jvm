package jd.jvmexample.main.classfile.annotation;

/**
 * Created by huangxia on 2019/7/9.
 */
@interface DefinedAnnotation {

    public String name() ;
}


public class DemoAnotation{

    // RuntimeVisibleAnnotations Attribute
    @DefinedAnnotation(name="abc")
    public String field ;
}