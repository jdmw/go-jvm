package main

import (
	"../classloader"
	"fmt"
)

func printUsage(){
	fmt.Printf("Usage: java [options] class [args...]")
}

func startJvm(cmd *Cmd ){
	fmt.Printf("Start "+cmd.classname+" ");
	fmt.Println(cmd.args)

	loader := classloader.NewClassloader(cmd.Xjre,cmd.classpath)
	classfile := loader.LoadClass(cmd.classname)
	if classfile == nil {
		panic("can not load main class " + cmd.classname)
	}
}

// run:  main -cp example\java\jvm-example-main\target\classes jd.jvmexample.main.Main
func main(){
	cmd := parseCmd()
	if cmd.helpFlag {
		printUsage()
	} else if cmd.versionFlag {
		fmt.Println( "version " + VERSION)
	} else {
		startJvm(cmd)
	}
}