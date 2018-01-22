package main

import "fmt"
import "strings"
import "jvmgo/ch02/classpath"

func main(){
	cmd:=parseCmd()
	if cmd.versionFlag{
		fmt.Println("version 0.0.1")
	}else if cmd.helpFlag || cmd.class ==""{
		printUsage()
	}else{
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd){
	cp:=classpath.Parse(cmd.XjreOption,cmd.cpOption)
	fmt.Printf("classpath:%s    class:%s     args:%v\n",cmd.cpOption,cmd.class,cmd.args)
	className:=strings.Replace(cmd.class,".","/",-1)
	fmt.Printf("classname:"+className+"\n")
	classData, _,err:=cp.ReadClass(className)
	if err!=nil{
		fmt.Printf("could not find or load main class %s\n",cmd.class)
		// return
	}
	fmt.Printf("class data:%v\n",classData)
}