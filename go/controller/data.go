package controller

import (
	"fmt"
	"poker/model"
	"poker/model/httpModel"
	"poker/service"
)
func Begining(cal*httpModel.CalData)float32{
	var req float32
	
	win:=float32(0)
	beginner:=&model.Begin{}
	beginner.Hand=[]int{58,59}
	beginner.Id=0
	beginner.PublicCard=[]int{}
	beginner.Person=6
    beginner.Frequency=50000
	result:=func(s float32){
		win+=s
	}
	for i:=0;i<beginner.Frequency;i++{
		service.GameMain(beginner,result)
	}
	req=(float32(win)/float32(beginner.Frequency))
	fmt.Println(req)
    return req
}