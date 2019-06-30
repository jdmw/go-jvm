package main

import "flag"

const VERSION = "1.0.0"

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	classpath string
	enableassertions bool
	javaagent string
	classname string
	Xjre string
	args        []string
}


func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag,"?",false,"print help messag")
	flag.BoolVar(&cmd.versionFlag,"version",false,"print version and exit")
	flag.StringVar(&cmd.classpath,"cp","/","class path")
	flag.StringVar(&cmd.classpath,"classpath","/","class path")
	flag.BoolVar(&cmd.enableassertions,"enableassertions",false,"enable assertions")
	flag.BoolVar(&cmd.enableassertions,"ea",false,"enable assertions")
	flag.StringVar(&cmd.javaagent,"javaagent","","java agent");
	flag.StringVar(&cmd.Xjre,"Xjre","","jre path");
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.classname = args[0]
		cmd.args = args[1:]
	}else{
		cmd.helpFlag = true
	}
	return cmd ;
}
