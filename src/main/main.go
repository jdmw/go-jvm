package main

import (
	"../classloader"
	"../runtime"
	"fmt"
)

func printUsage(){
	fmt.Printf("Usage: java [options] class [args...]")
}

func startJvm(cmd *Cmd ){
	fmt.Printf("Start "+cmd.classname+" ");
	fmt.Println(cmd.args)

	loader := classloader.NewClassloader(cmd.Xjre,cmd.classpath)
	mainClass := loader.LoadClass(cmd.classname)
	if mainClass == nil {
		panic("can not load main class " + cmd.classname)
	}

	mainMethod := mainClass.FindMethod("main","Ljava/lang/String;")
	if mainMethod == nil {
		panic("can't find main method in class " + cmd.classname)
	}

	mainThread := runtime.NewThread(STACK_DEPTH)
	mainFrame := mainThread.LoadStaticMethod(mainMethod)

}

// run:  main -cp example\java\jvm-example-main\target\classes jd.jvmexample.main.Main
func main(){
	//runtime.SlotTest()
	cmd := parseCmd()
	if cmd.helpFlag {
		printUsage()
	} else if cmd.versionFlag {
		fmt.Println( "version " + VERSION)
	} else {
		startJvm(cmd)
	}
}

