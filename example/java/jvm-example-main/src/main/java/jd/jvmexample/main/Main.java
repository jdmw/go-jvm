package jd.jvmexample.main;

import jd.jvmexample.common.JvmExampleCommon;
import jd.jvmexample.util.JvmExampleUtil;

public class Main {

    public static void main(String[] args){
        System.out.println("This is the jvm example of version " + JvmExampleUtil.toString(JvmExampleCommon.VERSION));
    }
}
