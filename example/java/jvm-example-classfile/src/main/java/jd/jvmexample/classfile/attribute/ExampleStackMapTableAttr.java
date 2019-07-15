package jd.jvmexample.classfile.attribute;

class ExampleStackMapTableAttr{
  
  public void f1(){
    int a = 1;
    if(a > 2){
      Object o = new Object() ;
      java.util.Date d = new java.util.Date();
      a = 2 ;
    }else if(a> 3){
      long c = 3 ;
      a = (int)c ;
    }else {
      String h = "h";
      a = h.length();
    }
  }
  
  public int parallel(){
    short a = 1;
    if(a > 1){
      boolean flag = false ;
      if(flag ) {
        int b1 = 11 ;
        return b1 ;
      }else{
        throw new RuntimeException(flag+"");
      }
    }else if(a> 3){
      int c = 3 ;
      return c ;
    }else {
      String h = "h";
      String y = h + "c";
      return y.length() ;
    }
  }
}