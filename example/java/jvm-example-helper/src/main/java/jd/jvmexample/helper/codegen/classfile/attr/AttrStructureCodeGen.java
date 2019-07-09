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

    public static Map<String,Element> list(){
        Map<String,Element> attrNodes = new HashMap<>(30);
        for(Element e : doc.select("a[name=\"jvms-4.7-300\"]").next().next().first().select(".table-contents table tbody tr")) {
            String name = e.select("code").html();
            String href = e.select("a").attr("href");
            href = href.substring(href.indexOf("#")+1);
            if(!new File(gegCodeBase,"attr_"+toUnderscopeName(name)+".go").exists()){
                System.out.printf("attr: %s - a[name='"+href+"']\n",name,href);
                attrNodes.putIfAbsent(name,doc.select("a[name='"+href+"']").parents().select(".section").first());
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

    public static AttrStructFile parseAttr(String name ,Element e){
        AttrStructFile struct =  new AttrStructFile(name+"Attr");
        StringBuilder pres = new StringBuilder();
        Elements preNodes = e.select("pre");

        for(Element pre : preNodes ){
            pres.append("\n\n").append(pre.html());
        }
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
                if(!"attribute_name_index".equals(arr[1]) && !"attribute_length".equals(arr[1])){
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
        list().forEach((name,e)->{
            if("RuntimeVisibleAnnotations".equals(name)) {
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
        //gegCodeBase.getAbsolutePath()
        //System.out.println(gegCodeBase);
        complete.forEach(System.out::println);



        //System.out.println("");
    }
}
