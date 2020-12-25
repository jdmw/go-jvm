package jd.jvmexample.classfile.attribute.annotation;


@interface Preliminary {}
@interface Copyright {
    String value();
}
@interface Endorsers {
    String[] value();
}

@Copyright("2002 Yoyodyne Propulsion Systems, Inc.")
class OscillationOverthruster {  }

@Endorsers({"Children", "Unscrupulous dentists"})
class Lollipop { }
@Endorsers("Epicurus") // the curly braces are omitted)
class Pleasure { }
class GorgeousFormatter implements Formatter {  }

@PrettyPrinter(GorgeousFormatter.class)
class Petunia { }

// Illegal; String is not a subtype of Formatter
/*@PrettyPrinter(String.class)
class Begonia {}*/
@interface Name{
    String first();
    String last();
}
@interface Author{
    Name value();
}
@Author(@Name(first = "Joe", last = "Hacker"))
class BitTwiddle {  }

//  permits other element declarations besides method declarations
@interface Quality {
    enum Level { BAD, INDIFFERENT, GOOD }
    Level value();
}
