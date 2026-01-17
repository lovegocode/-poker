package main

import (
	"poker/service"
	"fmt"
)

func main(){
	//service.GameStartText()
	win:=0
	fre:=10000
	score:=func(s int){
		win+=s
	}
	for i:=0;i<fre;i++{
    service.GameMain(10,score)
	} 
	 w:=float32(win)
	 f:=float32(fre)
	s:=w / f
	fmt.Println(s) 
	 
} 