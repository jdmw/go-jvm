package jd.jvmexample.instruction.controls;

public class EInstControls {

    public boolean tableswitch(int param){
        switch (param) {
            case 1 :
            case 2 :
            case 3 : return true ;
            default: return false ;
        }
    }

    public boolean lookupswitch(int param){
        switch (param) {
            case 1 :
            case 3 : return true ;
            default: return false ;
        }
    }
    
    public static void main(String[] args) {
        System.out.println("");
    }
}
