package jd.jvmexample.classfile.instrution;

/**
 * Created by huangxia on 2008/12/1.
 */
public class DemoPojo {
    private int a ;
    private static int b ;

    public static int getB() {
        return b;
    }

    public static void setB(int b) {
        DemoPojo.b = b;
    }

    public DemoPojo() {}
    public DemoPojo(int a) {
        this.a = a;
    }

    public int getA() {
        return a;
    }

    public void setA(int a) {
        this.a = a;
    }

    public static void main(String[] args){
        DemoPojo d = new DemoPojo(1);
        d.setA(1);
        int a = d.getA();
        d.setB(2);
        int b = d.getB();
    }
}
