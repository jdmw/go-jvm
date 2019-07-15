package jd.jvmexample.helper.codegen.classfile.attr;

import java.util.ArrayList;
import java.util.List;

public class AttrType {


    public enum DataType {
        Simple,
        Struct,
        Union,
        Array
    }

    private String name;
    private List<String> originDef ;// jvm 规范定义的结构
    private int originStart,originEnd ;

    private List<AttrType> membors ;

    public void putCode(List<String> originDef,int originStart,int originEnd){
        this.originDef = originDef;
        this.originStart = originStart;
        this.originEnd = originEnd;
    }

    public String printCode(String lineStart)
    {
        StringBuilder sb = new StringBuilder();
        for(int i=originStart;i<originEnd;i++){
            String line = this.originDef.get(i);
            if(line.contains("/*") && line.contains("*/")){
                line.replace("/*","//").replace("*/","");
            }
            sb.append(lineStart).append(line).append("\n");
        }
        return sb.toString() ;
    }

    public static void main(String[] args) {
        System.out.println("");
    }
}
