package jd.jvmexample.main.classfile.annotation;


import java.lang.annotation.ElementType;
import java.lang.annotation.Repeatable;
import java.lang.annotation.Target;


@Repeatable(FooContainer.class)
@interface Foo {}

//@Target(ElementType.ANNOTATION_TYPE)
@interface FooContainer {
    Foo[] value();
}

@Foo
@Foo
interface X {}


