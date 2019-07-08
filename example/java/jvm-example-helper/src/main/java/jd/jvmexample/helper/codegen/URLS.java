package jd.jvmexample.helper.codegen;

import org.apache.commons.io.IOUtils;
import org.jsoup.Connection;
import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;

import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.OutputStream;

public class URLS {

    public final static String JVM8_SPCS = "https://docs.oracle.com/javase/specs/jvms/se12/html" ;
    public final static String SPCS_CLASSFILE_FORMAT_URL = JVM8_SPCS + "/jvms-4.html" ;


    public static Document get(String url)  {
        try {
            File file = getCacheFile(url) ;
            if(file.exists()){
                return Jsoup.parse(file,"UTF-8");
            }else {
                Connection.Response response = Jsoup.connect(url).execute();
                try (OutputStream os = new FileOutputStream(file)){
                    os.write(response.bodyAsBytes());
                }
                return response.parse();
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    private static File getCacheFile(String url){
        File dir = new File(System.getProperty("java.io.tmpdir"),"jvm-example-helper");
        if(url.contains("://")){
            url = url.substring(url.indexOf("//")+2);
        }
        File file = new File(dir,url.replaceAll("/","_"));
        if(!file.getParentFile().exists()){
            file.getParentFile().mkdirs();
        }
        return file ;
    }
}