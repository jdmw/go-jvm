package jd.jvmexample.helper.codegen.classfile.attr;

import jd.jvmexample.helper.codegen.URLS;
import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;
import org.jsoup.nodes.Element;
import org.jsoup.select.Elements;

import java.io.*;
import java.util.*;


public class AttrStructureCodeGen {

    static Document doc = URLS.get(URLS.SPCS_CLASSFILE_FORMAT_URL);
    static File gegCodeBase = new File(new File("..").getAbsolutePath().substring(0,new File("..")
            .getAbsolutePath().lastIndexOf("example"))+"src"+File.separator+"classloader"+File.separator+"classfile");

    static class ElementInfo{
        String name ;
        String chapter ;
        String url ;
        File goSrcFile ;
        Element element;
        public ElementInfo(String name) {
            this.name = name;
        }
    }
    public static LinkedHashMap<String,ElementInfo> list(boolean skipExistGoSrc){
        LinkedHashMap<String,ElementInfo> attrNodes = new LinkedHashMap<>(30);
        for(Element e : doc.select("a[name=\"jvms-4.7-300\"]").next().next().first().select(".table-contents table tbody tr")) {
            String name = e.select("code").html();
            ElementInfo info = new ElementInfo(name);
            String href = e.select("a").attr("href");
            href = href.substring(href.indexOf("#")+1);
            info.url = URLS.SPCS_CLASSFILE_FORMAT_URL + "#" + href;
            info.chapter = href.split("-")[1];
            info.goSrcFile = new File(gegCodeBase,"attr_"+toUnderscopeName(name)+".go");
            if(!skipExistGoSrc || !info.goSrcFile.exists()){
                //System.out.printf("attr: %s - a[name='"+href+"']\n",name,href);
                info.element = doc.select("a[name='"+href+"']").parents().select(".section").first();
                attrNodes.putIfAbsent(name,info);
            }
        }
        return attrNodes;
    }

    public static String toUnderscopeName(String structname){
        StringBuffer filename = new StringBuffer(structname.substring(0,1)) ;
        for(char ch : structname.substring(1).toCharArray()){
            if(ch >= 'A' && ch <= 'Z'){
                filename.append("_");
            }
            filename.append(ch);
        }
        return filename.toString().toLowerCase();
    }

    static  StringBuilder mdGenSb = new StringBuilder("# 4.7. [Attributes]("+URLS.SPCS_CLASSFILE_FORMAT_URL+"#jvms-4.7)\n");
    public static AttrStructFile parseAttr(String name ,ElementInfo info){
        AttrStructFile struct =  new AttrStructFile(name+"Attr");
        StringBuilder pres = new StringBuilder();
        Elements preNodes = info.element.select("pre");

        mdGenSb.append("## ").append(info.chapter).append(" [").append(info.name).append("]("+info.url+") Attribute\n")
                .append("```text\n");
        for(Element pre : preNodes ){
            pres.append("\n\n").append(pre.html());
            mdGenSb.append("\n\n").append(pre.html());
        }
        mdGenSb.append("\n```\n");
        struct.originCodes = pres.toString();
        if(preNodes.size() > 0){
            Scanner scanner = new Scanner(preNodes.get(0).text());
            scanner.useDelimiter("\n");
            scanner.next();
            //scanner.next();
            while (scanner.hasNext()){
                String line = scanner.next().trim();
                if(line.startsWith("{") || line.startsWith("}")) {
                    if(line.length() > 1){
                        struct.finish = false ;
                    }
                    continue;
                }

                String[] arr = line.replace(";","").split("\\s+");
                //System.out.append("")
                if(arr.length > 2 && !"attribute_name_index".equals(arr[1]) && !"attribute_length".equals(arr[1])){
                    struct.attrs.putIfAbsent(arr[1],arr[0]);
                }
            }
            System.out.append(pres);
        }else {
            System.err.println(name + " no codes");
        }
        return struct;
    }
    public static void main(String[] args) throws IOException {

        List<String> complete = new ArrayList<>();
        list(true).forEach((name,e)->{
            if("Exceptions".equals(name)) {
                AttrStructFile s = parseAttr(name,e);
                complete.add(String.format("\t\t%scase \"%s\" : return &%s{}\n",s.finish?"":"\\ TODO: ",name,s.name));
                try(OutputStream f = new FileOutputStream(new File(gegCodeBase,"attr_"+toUnderscopeName(name)+".go"))){
                    f.write(s.toString().getBytes());
                } catch (FileNotFoundException e1) {
                    e1.printStackTrace();
                } catch (IOException e1) {
                    e1.printStackTrace();
                }
            }
        });
        //System.out.println(mdGenSb);
        //gegCodeBase.getAbsolutePath()
        //System.out.println(gegCodeBase);
        //complete.forEach(System.out::println);



        //System.out.println("");
    }
}
