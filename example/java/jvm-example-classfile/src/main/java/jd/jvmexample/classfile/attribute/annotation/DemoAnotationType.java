package jd.jvmexample.classfile.attribute.annotation;

/**
 * Created by huangxia on 2019/7/9.
 * java8 语言规范 9.6. Annotation Types
 * ref: https://docs.oracle.com/javase/specs/jls/se8/html/jls-9.html#jls-9.6
 *
 */
//  Annotation Type Declaration

@interface DefinedAnnotation {
    enum Sex{ M,W}

    int    id() default  1;        // primary type
    String name() ;     // String type
    Sex sex();         // enum type
    Class cla();
    String[] params() ;
    String synopsis();  // Synopsis of RFE
    String engineer();  // Name of engineer who implemented RFE
    String date();      // Date RFE was implemented
}
interface Formatter {}
// Designates a formatter to pretty-print the annotated class
@interface PrettyPrinter {
    Class<? extends Formatter> value();
}


@interface OtherNestedDeclarations {
    // also declare interfaces and classes
    interface Action {}
    class Clazz{
        int i ;
    }
}

class DemoAnotation{

    // RuntimeVisibleAnnotations Attribute
    //@DefinedAnnotation(name="abc")

    public String field ;
    @Quality(Quality.Level.BAD) OtherNestedDeclarations.Clazz e ;
}