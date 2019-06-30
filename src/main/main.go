package main

import (
	"classloader/classpath"
	"fmt"
)

func printUsage(){
	fmt.Printf("Usage: java [options] class [args...]")
}

func startJvm(cmd *Cmd ){
	fmt.Printf("Start "+cmd.classname+" ");
	fmt.Println(cmd.args)

	classPath := classloader.NewClassPath(cmd.Xjre,cmd.classpath)
	data,_,_ := classPath.ReadClass(cmd.classname)
	if data == nil {
		panic("can not load main class " + cmd.classname)
	}
	fmt.Printf("class data: %v",len(data))
}
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