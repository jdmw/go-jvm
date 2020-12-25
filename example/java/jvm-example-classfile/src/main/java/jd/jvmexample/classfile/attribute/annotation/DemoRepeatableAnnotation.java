package jd.jvmexample.classfile.attribute.annotation;


import java.lang.annotation.Repeatable;


@Repeatable(FooContainer.class)
@interface Foo {}

//@Target(ElementType.ANNOTATION_TYPE)
@interface FooContainer {
    Foo[] value();
}

@Foo
@Foo
interface X {}


