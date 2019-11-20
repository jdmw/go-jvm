package jd.jvmexample.helper.codegen.classfile.attr;

import jd.jvmexample.helper.codegen.URLS;

import java.util.*;

public class AttrStructure {

    String name ;
    String originCodes ;
    boolean finish ;
    LinkedHashMap<String,LinkedHashMap<String,String>> attrsMap = new LinkedHashMap<>();

    public AttrStructure(String name) {
        this.name = name;
    }

    public void putStruct(String code){

    }
    public String toString(){
        StringBuilder sb = new StringBuilder("package classfile\n");
        sb.append("\n/*\n*ref:").append(URLS.SPCS_CLASSFILE_FORMAT_URL).append(originCodes.replaceAll("/\\*","//").replaceAll("\\*/","").replaceAll("\n","\n * ")).append("\n */\n");

        attrsMap.forEach((originName,attrs)-> {
            String name = toUpperCaseName(originName);
            sb.append("\ttype ").append(name).append(" struct{\n\t\tcp ConstantPool\n");
            attrs.forEach((n, type) -> {
                sb.append("\t\t").append(n).append("\t").append(type).append("\n");
            });
            finish = true;
            // constructor
            sb.append("\t}\n\tfunc (self ").append(name).append(") parse(cf ClassFile,"+(originName.equals(this.name.replace("Attr","_attribute"))?"length util.U4,":"")+"r *util.BigEndianReader) {\n")
                    .append("\t\tself.cp = cf.constant_pool\n");

            attrs.forEach((n, type) -> {
                if (type.startsWith("u")) {
                    sb.append("\t\tself.").append(n).append(" = r.ReadU").append(type.substring(1)).append("()\n");
                }else {
                    sb.append("\t\t// TODO: self.").append(n).append(" = \n");
                    finish = false;
                }
            });
            sb.append("\t}\n\n");

            /*attrs.forEach((n, type) -> {
                if (type.startsWith("u")) {
                    sb.append("\t\tself.").append(name).append(" = r.ReadU").append(type.substring(1)).append("()\n");
                } else {
                    sb.append("\t\t// TODO: self.").append(name).append(" = ");
                    finish = false;
                }
            });*/

            attrs.forEach((n, type) -> {
                sb.append("\n\n\tfunc (self *").append(name).append(") ").append(toUpperCaseName(n))
                        .append("() ").append(type).append("{\n\t\t return self.").append(n).append("\n\t}\n");
            });
        });
        return sb.toString();
    }

    public static String toUpperCaseName(String structname){
        char[] arr = structname.toCharArray();
        StringBuffer filename = new StringBuffer(structname.substring(0,1).toUpperCase()) ;
        for(int i=1;i<arr.length;i++){
            if(arr[i] == '_'){
                i++;
                filename.append((char)(arr[i]-0x20));
            }else {
                filename.append(arr[i]);
            }
        }
        return filename.toString();
    }

    public static void main(String[] args) {
        System.out.println(toUpperCaseName("u_ab_file"));
        System.out.println(AttrStructureCodeGen.toUnderscopeName("AbCdeFile"));
    }
}
