package utils

import "poker/model"
func(T*TexasDealer) Deal (p *model.GameData,Cards[]int){
	all:=[][]int{}
	play:=p.GetAllPlayer()
    
  for _,v:=range play{
	process:=append([]int(nil),Cards...)
    process=append(process,v.Hand...)
	all=append(all,process)
  }
  
}